package main

import (
  "log"
  "net/http"
  "os"
  "time"

  "webrtc-video-chat/models"
  "webrtc-video-chat/routes"
  "webrtc-video-chat/ws"

  "github.com/rs/cors"
  "github.com/urfave/negroni"
)

//func determineListenAddress() (string, error) {
//	port := os.Getenv("PORT")
//	if port == "" {
//		return ":3000", fmt.Errorf("$PORT not set")
//	}
//	return ":" + port, nil
//}

const (
  //sslCert = "/etc/ssl/certs/rtc-selfsigned.crt"
  //sslKey  = "/etc/ssl/private/rtc-selfsigned.key"

  /*5.61.45.181*/
  sslCert = "rtc-selfsigned.crt"
  sslKey = "rtc-selfsigned.key"
)

func connectDatabase() {
  url := os.Getenv("DATABASE_URL")

  if url == "" {
    url = "user=postgres password=postgres dbname=postgres sslmode=disable"
  }

  db, err := models.Connect(url)

  if err != nil {
    log.Fatalf("Connection error: %s", err.Error())
  }

  models.SetDatabase(db)
}

func main() {
  go connectDatabase()

  go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
   http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
  }))

  hub := ws.H

  go hub.Run()

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
  })

  //addr, _ := determineListenAddress()
  routes := routes.NewRoutes()
  n := negroni.Classic()
  n.Use(c)
  n.UseHandler(routes)

  s := &http.Server{
    Addr:           ":443",
    Handler:        n,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
    //TLSConfig: &tls.Config{InsecureSkipVerify: true},
  }

  log.Fatal(s.ListenAndServeTLS(sslCert, sslKey))
}

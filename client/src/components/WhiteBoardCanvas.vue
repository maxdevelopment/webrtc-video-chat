<template>
  <div v-show="visibility" class="white-board-canvas">
    <p>
      <select v-model="currentColor">
        <option value="#000000">Black</option>
        <option value="#FFFFFF">White</option>
        <option value="#FF0000">Red</option>
        <option value="#00FF00">Green</option>
        <option value="#0000FF">Blue</option>
      </select>

      <select v-model="currentSize">
        <option value="3">Small</option>
        <option value="9">Medium</option>
        <option value="27">Large</option>
      </select>
    </p>

    <div id="canvascontainer">
      <canvas
        id="canvas"
        width="800"
        height="600"
        @mousedown="startDraw($event)"
        @mouseup="stopDraw($event)"
        @mouseout="stopDraw($event)"
        @mousemove="updatelastPosition($event)"></canvas>
      <i id="cursor" :style="cursorPosition"></i>
    </div>

    <p v-html="positionQueue.length"></p>
    <p v-html="positionQueue"></p>

  </div>
</template>

<script>
  const PARTIAL_SEND_LIMIT = 15;

  export default {
    name: "WhiteBoardCanvas",
    data: () => ({
      currentColor: '#000000',
      currentSize: 3,
      lastPosition: {
        x: 0,
        y: 0,
      },
      positionQueue: [],
      isDrawing: false,
      context: {},
      ws: false
    }),
    watch: {
      currentColor(value) {
        if (!this.context) {
          return;
        }
        this.context.strokeStyle = value;
      },
    },
    computed: {
      visibility() {
        return this.$store.getters.whiteBoard.visibility
      },
      cursorPosition() {
        return {
          top: `${this.lastPosition.y}px`,
          left: `${this.lastPosition.x}px`,
          'background-color': this.currentColor,
          width: `${this.currentSize}px`,
          height: `${this.currentSize}px`,
        };
      },
      currentConnection() {
        return this.$store.getters.whiteBoard
      }
    },
    mounted() {
      console.log('mounted');
      this.context = document.getElementById('canvas').getContext('2d');
      this.context.clearRect(0, 0, 800, 600);
      this.context.lineJoin = 'round';
      this.context.lineCap = 'round';
      // this.ws = new WebSocket(window.location.href.replace(/http(s?)/, 'ws$1'));
      const ws = this.$store.getters.whiteBoard.visibility;
      console.log(ws);
      // ws.onmessage = (event) => {
      //   console.log('event');
      //   console.log(event);
      //   //   const lineDetails = JSON.parse(event.data);
      //   //   this.drawLine(lineDetails.data, lineDetails.color, lineDetails.width);
      // };
      setInterval(() => {
        if (this.isDrawing) {
          const positionQueueLength = this.positionQueue.length;
          if (positionQueueLength > 1) {
            const previousPosition = this.positionQueue[positionQueueLength - 1];
            const xDifference = Math.abs(previousPosition.x - this.lastPosition.x);
            const yDifference = Math.abs(previousPosition.y - this.lastPosition.y);
            if (xDifference === 0 && yDifference === 0) {
              return;
            }
          }
          this.positionQueue.push({
            x: this.lastPosition.x,
            y: this.lastPosition.y,
          });
          if (this.positionQueue.length > PARTIAL_SEND_LIMIT) {
            this.sendLine(this.positionQueue.slice(0, PARTIAL_SEND_LIMIT));
            this.positionQueue = this.positionQueue.slice(PARTIAL_SEND_LIMIT - 1, this.positionQueue.length - 1);
          }
          if (this.positionQueue.length > 4) {
            this.drawLine(this.positionQueue.slice(-5), this.currentColor, this.currentSize);
          }
        }
      }, 15);
    },
    methods: {
      startDraw(e) {
        console.log('start');
        this.isDrawing = true;
      },
      stopDraw(e) {
        console.log('stop');
        if (this.isDrawing) {
          this.sendLine(this.positionQueue);
          this.positionQueue = [];
        }
        this.isDrawing = false;
      },
      updatelastPosition(e) {
        const rect = e.target.getBoundingClientRect();
        this.lastPosition.x = e.pageX - rect.left;
        this.lastPosition.y = e.pageY - rect.top - window.pageYOffset;
      },
      drawLine(points, color, size) {
        if (points.length === 0 || !this.context) {
          return;
        }
        this.context.strokeStyle = color;
        this.context.lineWidth = size;
        this.context.beginPath();
        this.context.moveTo(points[0].x, points[0].y);

        for (let index = 1; index < points.length; index += 1) {
          this.context.lineTo(points[index].x, points[index].y);
        }
        this.context.stroke();
      },
      sendLine(points) {
        const roomId = this.$store.getters.whiteBoard.roomId;
        const ws = this.$store.getters.whiteBoard.ws;

        ws.send(JSON.stringify({
          command: 'DRAW',
          color: this.currentColor,
          width: this.currentSize,
          data: points,
        }));
      }
    }
  }
</script>

<style lang="stylus" scoped>
  .white-board-canvas
    position absolute
    background-color #c6eaff
    top 50%
    left 50%
    transform translate(-50%, -50%)
    border-radius 5px
    -webkit-box-shadow: -4px 7px 8px 0px rgba(0, 0, 0, 0.25);
    box-shadow: -4px 7px 8px 0px rgba(0, 0, 0, 0.25);
    z-index 4

  #canvascontainer
    display block
    cursor none

  #cursor
    position absolute
    top -10px
    left -10px
    border-radius 50%
    pointer-events none
    transform translate(-50%, -50%)
</style>

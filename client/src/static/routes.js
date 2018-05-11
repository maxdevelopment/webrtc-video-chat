const basePath = 'http://127.0.0.1:3000'

const wsSchema = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
const baseSocket = `${wsSchema}//127.0.0.1:3000`

export default {
  createRoomPath: () => (
    `${basePath}/api/create_room`
  ),
  getRoomPath: slug => (
    `${basePath}/api/get_room/${slug}`
  ),
  connectToRoomPath: slug => (
    `${basePath}/api/room/${slug}`
  ),
  socketPath: slug => (
    `${baseSocket}/channel/${slug}`
  )
}

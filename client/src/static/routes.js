const basePath = window.location.hostname === 'rtc.loc' ? 'https://127.0.0.1' : 'https://5.61.45.181'
const wsSchema = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
const baseSocket = window.location.hostname === 'rtc.loc' ? `${wsSchema}//127.0.0.1` : `${wsSchema}//5.61.45.181`

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

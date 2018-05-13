export default {
  state: {
    visibility: true,
    roomId: null,
    ws: null
  },

  getters: {
    whiteBoard: state => state
  },

  mutations: {
    SHOW_WHITE_BOARD(state, data) {
      state.visibility = !state.visibility
    },
    SET_ROOM_ID(state, data) {
      state.roomId = data
    },
    SET_WS(state, data) {
      state.ws = data
    }
  },

  actions: {
    showWhiteBoard({ commit }) {
      commit('SHOW_WHITE_BOARD')
    },
    setRoomId({ commit }, data) {
      commit('SET_ROOM_ID', data)
    },
    setWs({ commit }, data) {
      commit('SET_WS', data)
    }
  }
}

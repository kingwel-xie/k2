const state = {
  logs: [],
  messages: []
}

const mutations = {
  ADD_ERROR_LOG: (state, log) => {
    state.logs.push(log)
  },
  CLEAR_ERROR_LOG: (state) => {
    state.logs.splice(0)
  },
  ADD_MESSAGE_LOG: (state, log) => {
    state.messages.push(log)
  },
  CLEAR_MESSAGE_LOG: (state) => {
    state.messages.splice(0)
  }
}

const actions = {
  addErrorLog({ commit }, log) {
    commit('ADD_ERROR_LOG', log)
  },
  clearErrorLog({ commit }) {
    commit('CLEAR_ERROR_LOG')
  },
  addMessageLog({ commit }, log) {
    commit('ADD_MESSAGE_LOG', log)
  },
  clearMessageLog({ commit }) {
    commit('CLEAR_MESSAGE_LOG')
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

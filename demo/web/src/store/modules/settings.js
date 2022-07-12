import variables from '@/styles/element-variables.scss'
import defaultSettings from '@/settings'
import storage from '@/utils/storage'

const LOCAL_STORAGE_KEY = '__app_setting__'

const state = Object.assign({ theme: variables.theme }, defaultSettings, storage.get(LOCAL_STORAGE_KEY))

const mutations = {
  CHANGE_SETTING: (state, { key, value }) => {
    // eslint-disable-next-line no-prototype-builtins
    if (state.hasOwnProperty(key)) {
      state[key] = value
      storage.set(LOCAL_STORAGE_KEY, state)
    }
  }
}

const actions = {
  changeSetting({ commit }, data) {
    commit('CHANGE_SETTING', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}


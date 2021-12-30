import { getAllDict } from '@/api/dict'

const state = {
  sys: [],
  countryList: []
}

const mutations = {
  SET_ALL_DICT: (state, dict) => {
    const { systemList, channelList, customerList, partitionList, countryList, labelList, routeList,
      portList, warehouseList, brokerList, shippingLineList, trailerList } = dict

    const key = 'dictType'
    state.sys = (systemList || []).reduce((ret, val) => {
      (ret[val[key]] = ret[val[key]] || []).push({ label: val.dictLabel, value: val.dictValue })
      return ret
    }, {})
    state.countryList = countryList || []
  }
}

const actions = {
  getAllDict({ commit }) {
    return new Promise((resolve, reject) => {
      getAllDict().then(response => {
        if (!response || !response.data) {
          resolve()
        }
        commit('SET_ALL_DICT', response.data)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

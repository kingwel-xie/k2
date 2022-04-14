import { getAllDict } from '@/api/dict'

export const LABEL_KEY = '__kobh_label'
export const VALUE_KEY = '__kobh_value'

/* eslint-disable no-multi-spaces */
const extDicts = [
  { kind: 'TbxCountry',        field: 'countryList', labelKey: 'nameCN' }
]
/* eslint-enable no-multi-spaces */

const alias = [
  {
    kind: 'TbxUS',
    factory: (dictRegistry, listRegistry) => {
      return [
        dictRegistry['TbxCountry'],
        listRegistry['TbxCountry'].filter((ch) => ch.code === 'US')
      ]
    }
  }
]

const uselessFields = ['createBy', 'createdAt', 'remark', 'status', 'updateBy', 'updatedAt']

function deleteUselessFields(x) {
  uselessFields.forEach((field) => {
    delete x[field]
  })
}

const state = {
  dictRegistry: {},
  listRegistry: {}
}

const mutations = {
  RESET_ALL_DICT: (state, dict) => {
    const dictRegistry = {}
    const listRegistry = {}

    extDicts.forEach(({ kind, field, codeKey = 'code', labelKey = 'name' }) => {
      const list = dict[field] || []
      const dicts = {}
      list.forEach((i) => {
        deleteUselessFields(i)
        // LABEL_KEY VALUE_KEY for @/components/DictSelect
        i[LABEL_KEY] = i[labelKey]
        i[VALUE_KEY] = i[codeKey]
        dicts[i[codeKey]] = i
      })
      dictRegistry[kind] = dicts
      listRegistry[kind] = list
    })

    const systemDicts = dict.systemList || []

    const groupByDictType = {}

    systemDicts.forEach(({ dictType, dictValue, dictLabel, remark }) => {
      const group = groupByDictType[dictType] ||= { list: [], dicts: {}}
      const item = { [LABEL_KEY]: dictLabel, [VALUE_KEY]: dictValue, label: dictLabel, value: dictValue, remark }
      group.list.push(item)
      group.dicts[dictValue] = item
    })

    Object.entries(groupByDictType).forEach(([kind, { dicts, list }]) => {
      if (dictRegistry[kind]) {
        console.warn(`dictionary index name '${kind}' is duplicated`)
        return
      }
      dictRegistry[kind] = dicts
      listRegistry[kind] = list
    })

    alias.forEach(({ kind, factory }) => {
      const [dicts, list] = factory(dictRegistry, listRegistry)
      dictRegistry[kind] = dicts
      listRegistry[kind] = list
    })

    state.dictRegistry = dictRegistry
    state.listRegistry = listRegistry
  },
  REGISTRY_MISSING_DICTS: (state, { kind, dicts }) => {
    if (!state.dictRegistry[kind]) {
      state.dictRegistry[kind] = dicts
    }
  },
  REGISTRY_MISSING_LIST: (state, { kind, list }) => {
    if (!state.listRegistry[kind]) {
      state.dictRegistry[kind] = list
    }
  }
}

const actions = {
  async getAllDict({ commit }) {
    const res = await getAllDict()
    commit('RESET_ALL_DICT', res.data)
    return res
  },
  registryMissingDicts({ commit }, payload) {
    commit('REGISTRY_MISSING_DICTS', payload)
  },
  registryMissingList({ commit }, payload) {
    commit('REGISTRY_MISSING_LIST', payload)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

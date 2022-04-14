import store from '@/store'

const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  username: state => state.user.userName,
  introduction: state => state.user.introduction,
  roles: state => state.user.roles,
  permisaction: state => state.user.permisaction,
  permission_routes: state => state.permission.routes,
  topbarRouters: state => state.permission.topbarRouters,
  defaultRoutes: state => state.permission.defaultRoutes,
  sidebarRouters: state => state.permission.sidebarRouters,
  errorLogs: state => state.errorLog.logs,
  messageLogs: state => state.errorLog.messages,
  appInfo: state => state.system.info,
  dictRegistry: (state) => {
    return (kind) => {
      let dicts = state.dictionary.dictRegistry[kind]
      if (!dicts) {
        console.trace(`missing dicts of '${kind}'`)
        dicts = {}
        store.dispatch('dictionary/registryMissingDicts', { kind, dicts }).catch((err) => {
          console.log(`failure to dispatch 'dictionary/registryMissingDicts' of ${kind}`, err)
        })
      }
      return dicts
    }
  },
  listRegistry: (state) => {
    return (kind) => {
      let list = state.dictionary.listRegistry[kind]
      if (!list) {
        console.trace(`missing dict list of '${kind}'`)
        list = []
        store.dispatch('dictionary/registryMissingList', { kind, list }).catch((err) => {
          console.log(`failure to dispatch 'dictionary/registryMissingList' of ${kind}`, err)
        })
      }
      return list
    }
  }
}
export default getters

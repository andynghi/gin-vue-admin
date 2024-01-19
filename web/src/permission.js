import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import Nprogress from 'nprogress'

const whiteList = ['Login', 'Init']

const getRouter = async(userStore) => {
  const routerStore = useRouterStore()
  await routerStore.SetAsyncRouter()
  await userStore.GetUserInfo()
  const asyncRouters = routerStore.asyncRouters
  asyncRouters.forEach(asyncRouter => {
    router.addRoute(asyncRouter)
  })
}

async function handleKeepAlive(to) {
  if (to.matched.some(item => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1]
        if (element.name === 'layout') {
          to.matched.splice(i, 1)
          await handleKeepAlive(to)
        }
        //If the on-demand loading is not completed, wait for loading
        if (typeof element.components.default === 'function') {
          await element.components.default()
          await handleKeepAlive(to)
        }
      }
    }
  }
}

router.beforeEach(async(to, from) => {
  const routerStore = useRouterStore()
  Nprogress.start()
  const userStore = useUserStore()
  to.meta.matched = [...to.matched]
  handleKeepAlive(to)
  const token = userStore.token
  // Judgment in the whitelist
  document.title = getPageTitle(to.meta.title, to)
  if(to.meta.client) {
    return true
  }
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
      }
      // The token can be parsed but does not exist as a user id or role id, which will result in infinite calls.
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
          return { name: userStore.userInfo.authority.defaultRouter }
        } else {
          return { path: '/layout/404' }
        }
      } else {
        //Force log out of account
        userStore.ClearStorage()
        return {
          name: 'Login',
          query: {
            redirect: document.location.hash
          }
        }
      }
    } else {
      return true
    }
  } else {
    // When not in the whitelist and already logged in
    if (token) {
      //Add flag to prevent multiple acquisitions of dynamic routing and stack overflow
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
        if (userStore.token) {
          if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
            return { ...to, replace: true }
          } else {
            return { path: '/layout/404' }
          }
        } else {
          return {
            name: 'Login',
            query: { redirect: to.href }
          }
        }
      } else {
        if (to.matched.length) {
          return true
        } else {
          return { path: '/layout/404' }
        }
      }
    }
    // When not in the whitelist and not logged in
    if (!token) {
      return {
        name: 'Login',
        query: {
          redirect: document.location.hash
        }
      }
    }
  }
})

router.afterEach(() => {
  //Close the progress bar after the route is loaded
  document.getElementsByClassName('main-cont main-right')[0]?.scrollTo(0, 0)
  Nprogress.done()
})

router.onError(() => {
  // Destroy the progress bar after an error occurs in the routing
  Nprogress.remove()
})

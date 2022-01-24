import { Context, Middleware } from '@nuxt/types'
import { Auth } from 'aws-amplify'

const authMiddleware: Middleware = async (context: Context) => {
  const { route, redirect } = context
  if (route.path === '/' || route.path === '/login') {
    return
  }
  try {
    await Auth.currentAuthenticatedUser()
  } catch {
    return redirect('/login')
  }
}

export default authMiddleware

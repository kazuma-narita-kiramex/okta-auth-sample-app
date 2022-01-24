import { Context } from '@nuxt/types'

export default function ({ $axios }: Context) {
  $axios.onRequest((config) => {
    // DO
  })
}

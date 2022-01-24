<template>
  <div class="flex flex-row h-screen items-top justify-center">
    <div class="m-auto text-center">
      <h1 class="text-4xl font-bold py-4">You Are Authenticated</h1>
      <div>
        <span>Hello! </span><span class="font-bold">{{ email }}</span>
      </div>
      <div class="mt-4">
        <NuxtLink
          to="/"
          class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
        >
          Back To Top Page
        </NuxtLink>
      </div>
      <div>
        <button
          class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white mt-5 mb-3 py-1 px-4 border border-blue-500 hover:border-transparent rounded"
          @click="signOut()"
        >
          signout
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Auth } from 'aws-amplify'

export default Vue.extend({
  name: 'AuthenticatedPage',
  data() {
    return { email: '' }
  },
  async mounted() {
    try {
      const payload = await (await Auth.currentSession()).getIdToken().payload
      this.email = payload.email
    } catch (e) {
      console.error(e)
    }
  },
  methods: {
    async signOut() {
      await Auth.signOut()
    },
  },
})
</script>

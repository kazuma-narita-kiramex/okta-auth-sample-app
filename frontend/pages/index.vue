<template>
  <div class="flex flex-row h-screen items-top justify-center">
    <div class="m-auto text-center">
      <h1 class="text-4xl font-bold py-4">Okta Auth Sample App</h1>
      <NuxtLink
        to="/authenticated"
        class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
      >
        Go To Auth Page
      </NuxtLink>
      <div class="mt-20 min-w-full min-h-full border border-blue-500 rounded">
        <h2 class="text-xl font-bold m-3">
          Call Backend API With Okta Access Token
        </h2>
        <p :class="{ 'text-red-500': invalid, 'text-green-500': valid }">
          {{ backendResult }}
        </p>
        <button
          class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white mt-1 mb-3 py-1 px-4 border border-blue-500 hover:border-transparent rounded"
          @click="callBackendAPI()"
        >
          Go
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  data() {
    return {
      valid: false,
      invalid: false,
      backendResult: 'ready',
    }
  },
  methods: {
    async callBackendAPI() {
      try {
        this.valid = false
        this.invalid = false
        const result = await this.$axios.$get(this.$config.BACKEND_SERVER)
        if (result.status === 'ok') {
          this.backendResult = 'You Have Valid Token'
          this.valid = true
        } else {
          this.backendResult = 'You Have Invalid Token'
          this.invalid = true
        }
      } catch (e) {
        console.error(e)
        this.backendResult = 'API Call Error'
        this.invalid = true
      }
    },
  },
})
</script>

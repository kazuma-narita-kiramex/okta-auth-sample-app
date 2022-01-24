<template>
  <div class="flex flex-row h-screen items-top justify-center">
    <div class="m-auto text-center">
      <p class="text-2xl py-4">Redirecting you to Okta to login</p>
      <div class="flex justify-center space-x-3">
        <div class="animate-ping h-2 w-2 bg-blue-600 rounded-full"></div>
        <div
          class="animate-ping h-2 w-2 bg-blue-600 rounded-full animation-delay-200"
        ></div>
        <div
          class="animate-ping h-2 w-2 bg-blue-600 rounded-full animation-delay-400"
        ></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Auth } from 'aws-amplify'

export default Vue.extend({
  name: 'LoginPage',
  async mounted() {
    await Auth.federatedSignIn()
  },
})
</script>

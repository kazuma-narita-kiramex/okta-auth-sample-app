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
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'AuthenticatedPage',
  middleware: 'auth',
  data() {
    return { email: '' }
  },
  mounted() {
    if (this.$auth.user) {
      this.email = this.$auth.user.email as string
    }
  },
})
</script>

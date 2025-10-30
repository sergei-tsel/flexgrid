<template>
<div class="flex flex-col">
  <span class="flex-1 content-center">{{ me.data.email }}</span>
  <Digital
      class="flex-2"
      @click="apiRequestLogout"
  >
    Выйти
  </Digital>
</div>
</template>

<script setup lang="ts">
import Digital from "~/components/Digital.vue";

const me = reactive({
  data: {
    email: "",
  }
})

onMounted(apiRequestMe())

/** --- Плейсхолдеры API --- */
async function apiRequestMe() {
  const res = fetch('/api/auth/me', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  })

  let body
  try { body = await (await res).json() } catch { body = {} }
  me.data = body.data

  return res
}

async function apiRequestLogout() {
  return fetch('/api/auth/logout', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  })
}
</script>

<style scoped>
</style>

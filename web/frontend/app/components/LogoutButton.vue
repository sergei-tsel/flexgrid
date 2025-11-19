<template>
  <Digital
      :is-active="logoutButtonActive"
      @click="logout"
  >
    Выйти
  </Digital>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useRuntimeConfig} from "#app";

const logoutButtonActive = ref(false);

async function logout() {
  logoutButtonActive.value = true;

  const res = await apiRequestLogout()

  if (res.ok) {
    navigateTo('/')
  }
}

/** --- Плейсхолдеры API --- */
async function apiRequestLogout() {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  return await fetch(`${apiBaseUrl}/auth/logout`, {
    method: 'POST',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: 'include'
  })
}
</script>

<style scoped>
</style>

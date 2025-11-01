<template>
<div class="flex flex-col">
  <h2 class="flex-1 md:text-2xl max-md:text-shadow-md   content-center text-center overflow-x-auto overflow-y-hidden">
    Профиль
  </h2>
  <span class="flex-2 content-center text-center overflow-x-auto overflow-y-hidden">{{ me?.Email }}</span>
  <Digital
      class="flex-1 md:m-16 max-md:m-2"
      :is-active="logoutButtonActive"
      @click="logout"
  >
    Выйти
  </Digital>
</div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useRuntimeConfig} from "#app";

const logoutButtonActive = ref(false);

interface User {
  Email: string;
}

const me = ref<User | null>(null);

onMounted(async () => {
  await apiRequestMe()
})

async function logout() {
  logoutButtonActive.value = true;

  const res = await apiRequestLogout()

  if (res.ok) {
    navigateTo('/')
  }
}

/** --- Плейсхолдеры API --- */
async function apiRequestMe() {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  const res = await fetch(`${apiBaseUrl}/auth/me`, {
    method: 'GET',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
  })

  if (res.ok) {
    me.value = await res.json()
  } else {
    me.value = null
  }

  return res
}

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

<template>
  <h1 class="grid-element-i content-center text-center md:text-3xl max-md:text-xl f-full">Кабинет автора</h1>
  <h3 class="grid-element-c content-center text-center md:text-xl max-md:text-shadow-md max-md:mb-3">{{ me?.Email }}</h3>
  <LogoutButton
      class="grid-element-e"
  />
  <Digital
      class="grid-element-d"
      :is-active="modalWindowVisible"
      @click="toggleModelWindow(true)"
  >
    Поиск
  </Digital>
  <ArticleRedactor
      class="grid-element-a max-md:mt-3"
      :action="'create'"
      :article-id="''"
  />
  <ModalWindow
      v-if="modalWindowVisible"
      class="overflow-x-hidden overflow-y-auto"
      :header="'Поиск'"
      @clickClose="toggleModelWindow(false)"
  >
    <Search />
  </ModalWindow>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useRuntimeConfig} from "#app";

definePageMeta({
  layout: 'page',
})

const modalWindowVisible = ref(false);

const toggleModelWindow = (isOpen: boolean) => {
  modalWindowVisible.value = isOpen;
};

interface User {
  Email: string;
}

const me = ref<User | null>(null);

onMounted(async () => {
  await apiRequestMe()
})

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
</script>

<style scoped>
.grid-element-a {
  grid-area: A;
}

.grid-element-c {
  grid-area: C;
}

.grid-element-d {
  grid-area: D;
}

.grid-element-e {
  grid-area: E;
}

.grid-element-i {
  grid-area: I;
}
</style>

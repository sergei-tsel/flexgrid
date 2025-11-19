<template>
  <section class="mx-auto max-w-sm max-h-sm overflow-x-hidden overflow-y-auto">
    <h1 class="text-center md:text-2xl max-md:text-shadow-md" v-if="action === 'create'">
      Создать статью
    </h1>
    <h1 class="text-center md:text-2xl max-md:text-shadow-md" v-else-if="action === 'edit'">
      Редактировать статью
    </h1>
    <div class="flex flex-col gap-4">
      <label v-if="action === 'edit'" class="font-bold" for="id">ID:</label>
      <input
          v-if="action === 'edit'"
          id="id"
          type="text"
          v-model="articleEntity.value.Id"
          disabled
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <label class="font-bold" for="title">Заголовок:</label>
      <input
          id="title"
          type="text"
          v-model="articleEntity.value.Title"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <label class="font-bold" for="content">Контент:</label>
      <input
          id="content"
          type="text"
          v-model="articleEntity.value.Content"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <p v-if="error" class="text-sm text-red-600 mt-2">{{ error }}</p>

      <button
          class="mt-4 w-full rounded-md cursor-pointer bg-blue-600 text-white py-2 disabled:opacity-50"
          @click="saveChanges"
      >
        Сохранить
      </button>
    </div>
  </section>

</template>

<script setup lang="ts">
import {reactive, ref} from "vue";
import {useRuntimeConfig} from "#app";

const props = defineProps({
  action: {
    type: String,
    required: true,
  },
  articleId: {
    type: String,
  },
});

interface User {
  Id: number;
}

const me = ref<User | null>(null);

const articleEntity = reactive({
  value: {
    Id: 0,
    Title: "",
    Content: "",
  }
})

const error: Ref<string|null> = ref(null)

onMounted(async () => {
  await fetchItemData();
});

async function fetchItemData() {
  try {
    if (!props.articleId) {
      return;
    }

    await apiRequestGetOneArticle(props.articleId)
  } catch (err) {
    error.value = "Ошибка загрузки данных";
  }
}

async function saveChanges() {
  try {
    if (!props.articleId) {
      articleEntity.value = await (await apiRequestCreateArticle(
          articleEntity.value!.Title,
          articleEntity.value!.Content
      )).json();
    } else {
      articleEntity.value = await (await apiRequestUpdateArticle(
          articleEntity.value!.Id.toString(),
          articleEntity.value!.Title,
          articleEntity.value!.Content
      )).json();
    }

    navigateTo("/article/" + articleEntity.value.Id);
  } catch (err) {
    error.value = "Ошибка сохранения данных";
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

  me.value = await res.json();

  return res;
}

async function apiRequestGetOneArticle(id: string) {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  const res = await fetch(`${apiBaseUrl}/article/${id}`, {
    method: 'GET',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
  })

  articleEntity.value = await res.json();

  return res;
}

async function apiRequestCreateArticle(title: string, content: string) {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;
  return await fetch(`${apiBaseUrl}/article`, {
    method: 'POST',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
    body: JSON.stringify({ title, content })
  })
}

async function apiRequestUpdateArticle(id: string, title: string, content: string) {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  return fetch(`${apiBaseUrl}/article/${id}`, {
    method: 'POST',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
    body: JSON.stringify({ title, content })
  })
}
</script>

<style scoped>
</style>

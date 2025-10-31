<template>
  <section class="mx-auto max-w-sm max-h-sm overflow-x-hidden overflow-y-auto">
    <h1 class="text-center md:text-2xl" v-if="action === 'create'">
      Создать статью
    </h1>
    <h1 class="text-center md:text-2xl" v-else-if="action === 'edit'">
      Редактировать статью
    </h1>
    <div class="flex flex-col gap-4">
      <label
          v-if="action === 'edit'"
          class="font-bold">ID:
      </label>
      <input
          v-if="action === 'edit'"
          type="text"
          v-model="article.data.id"
          disabled
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <label class="font-bold">Заголовок:</label>
      <input
          type="text"
          v-model="article.data.title"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <label class="font-bold">Контент:</label>
      <input
          type="text"
          v-model="article.data.content"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />

      <button
          class="mt-4 w-full rounded-md bg-blue-600 text-white py-2 disabled:opacity-50"
          @click="saveChanges"
      >
        Сохранить
      </button>
    </div>
  </section>

</template>

<script setup lang="ts">
import {reactive} from "vue";

const props = defineProps({
  action: {
    type: String,
    required: true,
  },
  articleId: {
    type: String,
  }
});

const article = reactive({
  data: {
    id: 0,
    userId: 0,
    title: "",
    content: "",
  }
})

onMounted(async () => {
  await fetchItemData();
});

async function fetchItemData() {
  try {
    if (!props.articleId) return;

    await apiRequestGetOneArticle(props.articleId)
  } catch (error) {
    console.error('Ошибка загрузки данных:', error);
  }
}

async function saveChanges() {
  try {
    if (!props.articleId) {
      await apiRequestCreateArticle(article.data.userId, article.data.title, article.data.content);
    } else {
      await apiRequestUpdateArticle(article.data.id, article.data.title, article.data.content);
    }

    alert('Данные успешно сохранены!');
  } catch (error) {
    console.error('Ошибка сохранения данных:', error);
  }
}

/** --- Плейсхолдеры API --- */
async function apiRequestGetOneArticle(id: string) {
  const res = fetch(`/api/article/${id}`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  })

  let body
  try { body = await (await res).json() } catch { body = {} }
  article.data = body.data

  return res
}

async function apiRequestCreateArticle(userId: number, title: string, content: string) {
  return fetch('/api/article', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ userId, title, content })
  })
}

async function apiRequestUpdateArticle(id: number, title: string, content: string) {
  return fetch(`/api/article/${id}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, content })
  })
}
</script>

<style scoped>
</style>

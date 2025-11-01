<template>
  <article class="overflow-x-auto overflow-y-auto m-3">
    <header class="content-center text-center">
      <h1 class="text-2xl font-semibold text-gray-800">{{ articleEntity?.Title }}</h1>
    </header>
    <main class="mt-3 mb-3 overflow-x-auto overflow-y-auto">
      {{ articleEntity?.Content }}
    </main>
    <footer class="flex md:flex-row max-md:flex-col">
        <span>Создана: {{ new Date(articleEntity?.CreatedAt as string).toLocaleString() }}</span>
        <span class="md:ml-9">Изменена: {{ new Date(articleEntity?.UpdatedAt as string).toLocaleString() }}</span>
    </footer>
  </article>
</template>

<script setup lang="ts">
import {useRuntimeConfig} from "#app";
import {onMounted, ref} from "vue";

const props = defineProps({
  articleId: {
    type: String,
    required: true,
  },
})

interface ArticleEntity {
  Id: number;
  Title: string;
  Content: string;
  CreatedAt: string;
  UpdatedAt: string;
}

const articleEntity = ref<ArticleEntity | null>(null);

onMounted(async () => {
  await apiRequestGetOneArticle(props.articleId)
});

/** --- Плейсхолдеры API --- */
async function apiRequestGetOneArticle(id: string): Promise<Response> {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  const res = await fetch(`${apiBaseUrl}/article/${id}`, {
    method: 'GET',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
  });

  articleEntity.value = await res.json();

  return res;
}
</script>

<style scoped>

</style>

<template>
  <aside class="overflow-x-hidden overflow-y-auto">
    <h1 class="text-center md:text-2xl">
      Поиск
    </h1>
    <div class="mx-auto max-w-sm max-h-sm">
      <label class="font-bold mb-2">Заголовок статьи или эмейл автора:</label>
      <input
          type="text"
          v-model="search"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
          class="mt-4 w-full rounded-md cursor-pointer bg-blue-600 text-white py-2 disabled:opacity-50"
          @click="apiRequestSearchArticles(search)"
      >
        Найти
      </button>
    </div>
    <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 w-full mt-4"
    >
      <ArticleCard
          v-for="article in articles" :key="article.Id"
          :article="article"
          @click="getArticle(article)"
      />
    </div>
  </aside>
</template>

<script setup lang="ts">
const search = ref("");

interface ArticleEntity {
  Id: string;
  UserId: string;
  Title: string;
  Content: string;
  IsPublic: boolean;
}

const articles = ref<Array<ArticleEntity>>([]);

async function getArticle(article: ArticleEntity) {
  navigateTo("/article/" + article.Id);
}

/** --- Плейсхолдеры API --- */
async function apiRequestSearchArticles(search: string): Promise<Response> {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  const res = await fetch(`${apiBaseUrl}/article/search`, {
    method: 'POST',
    headers: {
      cookie: useRequestEvent()?.req.headers.cookie || '',
    },
    credentials: "include",
    body: JSON.stringify({ search }),
  });

  articles.value = await res.json();

  return res;
}
</script>

<style scoped>

</style>
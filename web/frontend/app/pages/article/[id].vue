<template>
  <article>
    <header>
      <a href="/cabinet">
        <Digital class="md:m-3 max-md:m-1">
          На главную
        </Digital>
      </a>
      <h1>{{ article?.title }}</h1>
    </header>
    <main>
      {{ article?.content }}
    </main>
    <footer>
      <p>
        Создана: {{ article?.created_at }}
        Изменена: {{ article?.updated_at }}
      </p>
    </footer>
  </article>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import Digital from "~/components/Digital.vue";

interface Article {
  id: number;
  title: string;
  content: string;
  created_at: string;
  updated_at: string;
}

const article = ref<Article | null>(null);

const route = useRoute();

onMounted(async () => {
  await apiRequestGetOneArticle(route.params.id as string);
});

/** --- Плейсхолдеры API --- */
async function apiRequestGetOneArticle(id: string): Promise<Response> {
  const res = await fetch(`/api/article/${id}`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  });

  let body: Article;
  try {
    body = await res.json();
  } catch (error) {
    console.error('Ошибка при разборе JSON:', error);
    body = {} as Article;
  }

  article.value = body;

  return res;
}
</script>

<style scoped>
</style>

<template>
  <h1 class="grid-element-intro content-center text-center md:text-3xl max-md:text-xl f-full">Кабинет автора</h1>
  <Profile class="grid-element-easy"/>
  <Digital
      class="grid-element-button md:m-3 max-md:m-1"
      :is-active="newArticle"
      @click="toggleForm(true)"
  >
    Создать статью
  </Digital>
  <Digital
      class="grid-element-control md:m-3 max-md:m-1"
      :is-active="!newArticle"
      @click="toggleForm(false)"
  >
    Найти статью
  </Digital>
  <Article
      class="grid-element-auth"
      v-if="!showSearchForm"
      :action="newArticle ? 'create' : 'edit'"
      :article-id="selectedArticle?.id ?? ''"
  />
  <aside v-if="showSearchForm" class="grid-element-auth">
    <h1 class="text-center md:text-2xl">
      Найти статью
    </h1>
    <div class="mx-auto max-w-sm max-h-sm overflow-x-hidden overflow-y-auto">
      <label class="block text-sm mb-2">Заголовок статьи или эмейл автора:</label>
      <input
          type="text"
          v-model="search"
          class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
          class="mt-4 w-full rounded-md bg-blue-600 text-white py-2 disabled:opacity-50"
      @click="apiRequestSearchArticles(search)">Сохранить</button>
    </div>
  </aside>
  <section v-if="showSelectedArticle" class="grid-element-daily">
      <span>Выбранная статья:</span>
      <h3>{{ selectedArticle?.title }}</h3>
      <button
          class="mt-4 py-2 cursor-pointer"
          v-if="!selectedArticle?.isPublic"
          @click="apiRequestPublishArticle(selectedArticle!.id)"
      >
        Опубликовать
      </button>
  </section>
  <section
      class="grid-element-ground grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 w-full mb-4 overflow-x-hidden overflow-y-auto"
  >
    <Card
      v-for="article in articles" :key="article.id"
      :article="article"
      @click="setArticle(article)"
    />
  </section>
</template>

<script setup lang="ts">
import Digital from "~/components/Digital.vue";

interface ArticleModel {
  id: string;
  title: string;
  content: string;
  isPublic: boolean;
}

definePageMeta({
  layout: 'base',
})

const showSearchForm = ref(false);
const newArticle = ref(true);

const toggleForm = (isNewArticle: boolean) => {
  newArticle.value = isNewArticle;
  showSearchForm.value = !isNewArticle;
  showSelectedArticle.value = false;
};

const search = ref("");
const articles = ref<Array<ArticleModel>>([]);

const showSelectedArticle = ref(false);
const selectedArticle = ref<ArticleModel | null>(null);

async function setArticle(article: ArticleModel) {
  selectedArticle.value = article;
  showSelectedArticle.value = true;
}

interface SearchResponse {
  data: Array<ArticleModel>;
}

async function apiRequestSearchArticles(search: string): Promise<SearchResponse> {
  const res = await fetch('/api/article/search', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ search })
  });

  const body = await res.json();
  articles.value = body.data;
  return body;
}

async function apiRequestPublishArticle(id: string): Promise<Response> {
  return fetch(`/api/article/${id}/publish`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });
}
</script>

<style scoped>
.grid-element-auth {
  grid-area: A;
}

.grid-element-button {
  grid-area: B;
}

.grid-element-control {
  grid-area: C;
}

.grid-element-daily {
  grid-area: D;
}

.grid-element-easy {
  grid-area: E;
}

.grid-element-ground {
  grid-area: G;
}

.grid-element-intro {
  grid-area: I;
}
</style>

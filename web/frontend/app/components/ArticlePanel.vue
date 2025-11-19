<template>
  <div v-if="isMy" class="flex flex-col justify-between gap-1">
    <Digital
        v-if="!articleEntity?.IsPublic"
        class="flex-1"
        :is-active="publishButtonActive"
        @click="publishArticle()"
    >
      Открыть статью публике
    </Digital>
    <Digital
        class="flex-1"
        :is-active="editButtonActive"
        @click="toggleModelWindow(true)"
    >
      Открыть редактор
    </Digital>
    <ModalWindow
        class="overflow-x-hidden overflow-y-auto"
        v-if="modalWindowVisible"
        :header="'Редактировать статью'"
        @clickClose="toggleModelWindow(false)"
    >
      <ArticleRedactor
          class="flex-11"
          :action="'edit'"
          :article-id="props.articleId"
      />
    </ModalWindow>
  </div>
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

const publishButtonActive = ref(false);
const editButtonActive = ref(false);

const modalWindowVisible = ref(false);

const toggleModelWindow = (isOpen: boolean) => {
  editButtonActive.value = isOpen
  modalWindowVisible.value = isOpen;
};

const isMy = ref(false);

interface User {
  Id: number;
}

const me = ref<User | null>(null);

interface ArticleEntity {
  Id: number;
  UserId: number;
  Title: string;
  Content: string;
  IsPublic: boolean;
}

const articleEntity = ref<ArticleEntity | null>(null);

function publishArticle(): void {
  publishButtonActive.value = true;
  apiRequestPublishArticle(articleEntity.value!.Id.toString())
}

onMounted(async () => {
  await apiRequestMe()
  await apiRequestGetOneArticle(props.articleId)
  isMy.value = articleEntity.value?.UserId == me.value?.Id;
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

  me.value = await res.json();

  return res;
}

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

async function apiRequestPublishArticle(id: string): Promise<Response> {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  const res = await fetch(`${apiBaseUrl}/article/${id}/publish`, {
    method: 'POST',
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

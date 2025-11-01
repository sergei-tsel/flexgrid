<template>
  <section class="grid-element-i flex flex-row content-start">
    <Digital
        class="md:m-9 max-md:m-1"
        :is-active="backButtonActive"
        @click="backButtonActive = true"
    >
      <a href="/cabinet" class="m-1">
        На главную
      </a>
    </Digital>
  </section>
  <ArticlePanel
      class="grid-element-e md:m-3 max-md:m-1 overflow-x-hidden overflow-y-auto"
      :article-id="route.params.id as string"
  />
  <Publication
      class="grid-element-a"
      :article-id="route.params.id as string"
  />
  <aside class="grid-element-d flex flex-col content-center">
    <Digital
        class="max-md:m-3 md:mt-6 md:mr-9 md:ml-9"
        :is-active="modalWindowVisible"
        @click="toggleModelWindow(true)"
    >
      Поиск
    </Digital>
  </aside>
  <ModalWindow
      class="overflow-x-hidden overflow-y-auto"
      v-if="modalWindowVisible"
      :header="'Поиск'"
      @clickClose="toggleModelWindow(false)"
  >
    <Search />
  </ModalWindow>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'

definePageMeta({
  layout: 'base',
})

const backButtonActive = ref(false)

const route = useRoute();

const modalWindowVisible = ref(false);

const toggleModelWindow = (isOpen: boolean) => {
  modalWindowVisible.value = isOpen;
};
</script>

<style scoped>
.grid-element-a {
  grid-area: A;
}

.grid-element-b {
  grid-area: B;
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

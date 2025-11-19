<template>
  <section class="grid-element-i flex flex-row content-start">
    <Digital
        class="md:m-1 max-md:ml-1 max-md:mt-1"
        :is-active="backButtonActive"
        @click="clickBuckButton"
    >
      На главную
    </Digital>
  </section>
  <ArticlePanel
      class="grid-element-e"
      :article-id="route.params.id as string"
  />
  <Publication
      class="grid-element-a"
      :article-id="route.params.id as string"
  />
  <Digital
      class="grid-element-d"
      :is-active="modalWindowVisible"
      @click="toggleModelWindow(true)"
  >
    Поиск
  </Digital>
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
  layout: 'page',
})

const backButtonActive = ref(false)

const route = useRoute();

const modalWindowVisible = ref(false);

const toggleModelWindow = (isOpen: boolean) => {
  modalWindowVisible.value = isOpen;
};

function clickBuckButton() {
  backButtonActive.value = true
  navigateTo('/cabinet')
}
</script>

<style scoped>
.grid-element-a {
  grid-area: A;
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

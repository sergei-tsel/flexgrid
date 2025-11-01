<template>
  <h1 class="grid-element-i content-center text-center md:text-3xl max-md:text-xl f-full">Кабинет автора</h1>
  <Profile class="grid-element-e md:m-3 max-md:m-1"/>
  <ArticleRedactor
      class="grid-element-a"
      :action="'create'"
      :article-id="''"
  />
  <aside class="grid-element-d flex flex-col content-center">
    <Digital
        class="max-md:m-2 md:mt-7 md:mr-9 md:ml-9"
        :is-active="modalWindowVisible"
        @click="toggleModelWindow(true)"
    >
      Поиск
    </Digital>
  </aside>
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
definePageMeta({
  layout: 'base',
})

const modalWindowVisible = ref(false);

const toggleModelWindow = (isOpen: boolean) => {
  modalWindowVisible.value = isOpen;
};
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

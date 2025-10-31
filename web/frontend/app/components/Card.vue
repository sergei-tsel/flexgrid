<template>
  <a
      class="min-w-0 bg-white border border-gray-100 rounded-2xl shadow-md p-5 flex flex-col items-start"
      href="'/article/' + article.data.id"
  >
    <h3 class="text-xl font-semibold text-gray-800">{{ article.data.title }}</h3>

    <div class="flex-grow-1"></div>

    <div class="flex flex-row items-center justify-between w-full mt-5">
      <div v-if="!article.data.isPublic" class="px-2 py-1 text-xs text-center hidden sm:block self-end min-w-20">
         Не опубликована
      </div>
    </div>
  </a>
</template>

<script setup lang="ts">
defineProps({
  article: {
    type: Object,
    required: true
  },
});
</script>

<style scoped>
</style>

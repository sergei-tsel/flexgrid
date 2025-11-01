<template>
  <a
      class="min-w-0 bg-white border border-gray-500 rounded-2xl shadow-md p-5 flex flex-col items-start"
      :href="'/article/' + article.Id"
  >
    <h3 class="text-xl font-semibold text-gray-800">{{ article.Title }}</h3>

    <div class="flex flex-row items-center justify-between w-full mt-5">
      <p v-if="!article.IsPublic" class="px-2 py-1 text-sm text-center sm:block self-end min-w-20">
         Не опубликована
      </p>
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

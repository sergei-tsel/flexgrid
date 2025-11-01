<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
    <div class="w-full max-w-lg mx-auto my-auto bg-white rounded-lg shadow-lg">
      <div class="flex items-center justify-between px-4 py-2 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-center grow">{{ props.header }}</h2>
        <button @click="onCloseButtonClick" class="ml-4 text-gray-500 hover:text-gray-700">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="p-4 overflow-x-hidden overflow-y-auto">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  header: {
    type: String,
    default: '',
  }
})

const emit = defineEmits(['clickClose'])

const onCloseButtonClick = () => {
  emit('clickClose')
}
</script>

<style scoped>
</style>

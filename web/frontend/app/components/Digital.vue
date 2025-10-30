<template>
  <section class="md:m-3">
    <button
      :class="[
        'transition-all duration-300 ease-in-out w-full h-full max-md:text-sm overflow-clip',
        props.isActive ? 'bg-yellow-300 text-black' : 'bg-black text-yellow-300'
      ]">
      <slot></slot>
    </button>
  </section>
</template>

<script setup lang="ts">
const props = defineProps({
  isActive: {
    type: Boolean,
    default: false,
  }
})
</script>

<style scoped>
</style>

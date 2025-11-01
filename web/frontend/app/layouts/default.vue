<template>
  <div class="min-h-screen grid place-content-stretch grid-rows-[1fr_10fr_1fr]">
    <SiteHeader class="row-start-1 row-end-2"/>
    <div class="row-start-2 row-end-3">
      <slot></slot>
    </div>
    <SiteFooter class="row-start-3 row-end-4"/>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
</style>

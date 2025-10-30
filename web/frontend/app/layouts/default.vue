<template>
  <div class="min-h-screen grid place-content-stretch grid-rows-[1fr_10fr_1fr]">
    <Header class="row-start-1 row-end-2"/>
    <div class="row-start-2 row-end-3">
      <slot></slot>
    </div>
    <Footer class="row-start-3 row-end-4"/>
  </div>
</template>

<script setup lang="ts">
import Header from "~/components/Header.vue";
</script>

<style scoped>
</style>

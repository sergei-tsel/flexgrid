<template>
  <Profile class="grid-element-easy"/>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'base',
})
</script>

<style scoped>
.grid-element-auth {
  grid-area: A;
}

.grid-element-button {
  grid-area: B;
}

.grid-element-control {
  grid-area: C;
}

.grid-element-daily {
  grid-area: D;
}

.grid-element-easy {
  grid-area: E;
}

.grid-element-foot {
  grid-area: F;
}

.grid-element-ground {
  grid-area: G;
}

.grid-element-head {
  grid-area: H;
}

.grid-element-intro {
  grid-area: I;
}
</style>

<template>
  <Digital
      class="grid-element-b md:m-3 max-md:m-1"
      :is-active="!isNewUser"
      @click="toggleForm(false)"
  >
    Войти
  </Digital>
  <Digital
      class="grid-element-c md:m-3 max-md:m-1"
      :is-active="isNewUser"
      @click="toggleForm(true)"
  >
    Создать аккаунт
  </Digital>
  <Auth v-if="showForm" :is-new-user="isNewUser" class="grid-element-a"></Auth>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'base',
})

import { ref } from "vue";

const showForm = ref(false);
const isNewUser = ref(false);

const toggleForm = (isRegistration: boolean) => {
  isNewUser.value = isRegistration;
  showForm.value = true;
};
</script>

<style scoped>
.grid-element-a {
  grid-area: A;
}

.grid-element-b {
  grid-area: B;
}

.grid-element-c {
  grid-area: C;
}
</style>

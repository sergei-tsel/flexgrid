<template>
  <Digital
      class="grid-element-button md:m-3 max-md:m-1"
      :is-active="!newUser"
      @click="toggleForm(false)"
  >
    Войти
  </Digital>
  <Digital
      class="grid-element-control md:m-3 max-md:m-1"
      :is-active="newUser"
      @click="toggleForm(true)"
  >
    Создать аккаунт
  </Digital>
  <Auth v-if="showForm" :is-new-user="newUser" class="grid-element-auth"></Auth>
</template>

<script setup lang="ts">
import Digital from "~/components/Digital.vue";

definePageMeta({
  layout: 'base',
})

import Auth from "~/components/Auth.vue";
import { ref } from "vue";

const showForm = ref(false);
const newUser = ref(false);

const toggleForm = (isNewUser: boolean) => {
  newUser.value = isNewUser;
  showForm.value = true;
};
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
</style>

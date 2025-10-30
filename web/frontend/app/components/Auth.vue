<template>
  <section class="mx-auto max-w-sm max-h-sm overflow-x-hidden overflow-y-auto">
    <label class="block text-sm mb-2" for="email">E-mail</label>
    <input
        id="email"
        type="email"
        inputmode="email"
        autocomplete="email"
        class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
        v-model.trim="email"
        placeholder="you@example.com"
        maxlength="254"
    />

    <label class="block text-sm mb-2 mt-4" for="password">Пароль</label>
    <input
        id="password"
        type="password"
        autocomplete="password"
        class="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
        v-model.trim="password"
        placeholder="*********"
        maxlength="100"
    />

    <p v-if="error" class="text-sm text-red-600 mt-2">{{ error }}</p>

    <button
        class="mt-4 w-full rounded-md bg-blue-600 text-white py-2 disabled:opacity-50"
        :disabled="loading || !isEmailValid"
        @click="submit"
    >
        {{ loading ? 'Отправляю...' : 'Продолжить с email' }}
    </button>
  </section>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const loading = ref(false)

const email = ref('')
const error: Ref<string|null> = ref(null)

const password = ref('')

const props = defineProps({
  isNewUser: {
    type: Boolean,
    default: false,
  }
})

/** Проверка email (RFC-5322-lite) */
const isEmailValid = computed(() => {
  const v = (email.value || '').toLowerCase()
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/
  return re.test(v) && v.length <= 254
})

async function submit() {
  error.value = null
  if (!isEmailValid.value) {
    error.value = 'Проверьте корректность e-mail.'
    return
  }
  if (password.value.trim().length < 6) {
    error.value = 'Пароль должен состоять хотя бы из шести символов.'
    return
  }

  loading.value = true
  try {
    props.isNewUser
        ? await apiRequestRegisterByEmail(email.value.trim(), password.value.trim())
        : await apiRequestLoginByEmail(email.value.trim(), password.value.trim());
  } finally {
    loading.value = false
  }
}

/** --- Плейсхолдеры API --- */
async function apiRequestRegisterByEmail(email: string, password: string) {
  return fetch('/api/auth/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  })
}

async function apiRequestLoginByEmail(email: string, password: string) {
  return fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  })
}
</script>

<style scoped>
</style>

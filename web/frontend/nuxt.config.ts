import { config as loadEnv } from 'dotenv'
import { resolve } from 'path'

loadEnv({ path: resolve(__dirname, '.env') })

export default defineNuxtConfig({
    compatibilityDate: '2025-10-15',
    devtools: { enabled: true },

    devServer: {
        host: '0.0.0.0'
    },

    modules: [
        '@nuxt/eslint',
        '@nuxt/fonts',
        '@nuxt/image',
        '@nuxt/ui',
        '@nuxt/icon',
    ]
})

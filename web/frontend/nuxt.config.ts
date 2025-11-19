import { config as loadEnv } from 'dotenv'
import { resolve } from 'path'
import tailwindcss from "@tailwindcss/vite";

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
    ],

    css: ['./app/assets/css/main.css'],

    vite: {
        plugins: [
            tailwindcss()
        ],
    },

    fonts: {
        providers: {
            fontshare: false
        }
    },

    runtimeConfig: {
        public: {
            apiBaseUrl: process.env.NODE_ENV === 'development' ? 'http://localhost:8080' : 'http://flexgrid.com/api',
        }
    }
})

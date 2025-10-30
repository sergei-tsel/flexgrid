/** @type {import('tailwindcss').Config} */
export default {
    content: [
        './components/**/*.{vue,js,ts}',
        './layouts/**/*.{vue,js,ts}',
        './pages/**/*.{vue,js,ts}',
        './app.vue',
        './nuxt.config.ts',
        './entry.client.ts',
        './entry.server.ts',
    ],
    theme: {
        extend: {},
    },
    plugins: [],
};

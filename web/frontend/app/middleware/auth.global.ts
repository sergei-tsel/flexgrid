import type { RouteLocationNormalized } from 'vue-router';

export default defineNuxtRouteMiddleware(
    async (
        to: RouteLocationNormalized,
        from: RouteLocationNormalized
    ) => {
        const config = useRuntimeConfig();
        const apiBaseUrl = config.public.apiBaseUrl;

        const res = await fetch(`${apiBaseUrl}/auth/me`, {
            method: 'GET',
            headers: {
                cookie: useRequestEvent()?.req.headers.cookie || '',
            },
            credentials: 'include',
        })

        if (!res.ok && to.path !== '/') {
            return navigateTo('/');
        }

        if (res.ok && to.path === '/') {
            return navigateTo('/cabinet');
        }
    });

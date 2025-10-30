import type { RouteLocationNormalized } from 'vue-router';

export default defineNuxtRouteMiddleware(
    (
        to: RouteLocationNormalized,
        from: RouteLocationNormalized
    ) => {
    const nuxtApp = useNuxtApp();
    const cookie = useCookie('cookie_name', { maxAge: 60 * 60 * 24 });

    const isAuthenticated = Boolean(cookie.value);

    if (!isAuthenticated && to.path !== '/') {
        return navigateTo('/');
    }

    if (isAuthenticated && to.path === '/') {
        return navigateTo('/cabinet');
    }
});

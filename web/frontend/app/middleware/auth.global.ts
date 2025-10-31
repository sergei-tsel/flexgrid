import type { RouteLocationNormalized } from 'vue-router';

export default defineNuxtRouteMiddleware(
    (
        to: RouteLocationNormalized,
        from: RouteLocationNormalized
    ) => {
    const cookie = useCookie('fg_id');

    const isAuthenticated = Boolean(cookie.value);

    if (!isAuthenticated && to.path !== '/') {
        return navigateTo('/');
    }

    if (isAuthenticated && to.path === '/') {
        return navigateTo('/cabinet');
    }
});

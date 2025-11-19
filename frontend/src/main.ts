import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import App from './app.vue'

import '@fortawesome/fontawesome-free/css/all.min.css'
import '@styles/custom.css'
import '@styles/ant.css'
import '@styles/btn.css'
import '@styles/text.css'
import '@styles/table.css'
import '@styles/badge.css'
import '@styles/portlet.css'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/base/user/edit' },
    { path: '/:path(.*)*', component: App },
  ],
})

createApp(App).use(router).mount('#app')

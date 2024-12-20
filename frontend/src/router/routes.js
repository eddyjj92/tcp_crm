const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') }
    ]
  },
  {
    path: '/products',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/products/index.vue') }
    ]
  },
  {
    path: '/suppliers',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/suppliers/index.vue') }
    ]
  },
  {
    path: '/purchases',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/purchases/index.vue') },
      { path: 'create', component: () => import('pages/purchases/create.vue') }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes

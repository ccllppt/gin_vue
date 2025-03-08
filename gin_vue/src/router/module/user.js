const userRouter = [
  {
    path: '/register',
    name: 'RegisterView',
    component: () => import('@/views/register/RegisterView.vue'),
  },
  {
    path: '/login',
    name: 'LoginView',
    component: () => import('@/views/login/LoginView.vue'),
  },
  {
    path: '/profile',
    name: 'ProfileView',
    meta: {
      auth: true,
    },
    component: () => import('@/views/profile/ProfileView.vue'),
  },
];
export default userRouter;

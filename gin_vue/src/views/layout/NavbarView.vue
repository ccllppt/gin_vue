<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-container>
        <b-navbar-brand @click="$router.push({name:'home'})">导航</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse id="nav-collapse" is-nav>
          <b-nav-item-dropdown right v-if="userInfo && userInfo.name">
            <template v-slot:button-content>
              <em>{{ userInfo.name }}</em>
            </template>
            <b-dropdown-item @click="$router.push({name:'ProfileView'})">个人主页</b-dropdown-item>
            <b-dropdown-item @click="signOut">注销</b-dropdown-item>
          </b-nav-item-dropdown>
          <b-navbar-nav class="ml-auto">
            <div v-if="!userInfo || !userInfo.name">
              <b-nav-item v-if="$route.name !='LoginView'"
              @click="$router.replace({name:'LoginView'})">登录</b-nav-item>
              <b-nav-item v-if="$route.name !='RegisterView'"
              @click="$router.replace({name:'RegisterView'})">注册</b-nav-item>
            </div>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script lang="ts">
import { mapState, mapActions } from 'vuex';

export default {
  computed: {
    ...mapState('userModule', ['userInfo']),
  },
  methods: {
    ...mapActions('userModule', ['sign_out']),
    async signOut() {
      await this.sign_out();
      // 检查当前路由名称，避免重复导航
      if (this.$route.name !== 'LoginView') {
        this.$router.replace({ name: 'LoginView' });
      } else {
        this.$router.go(0); // 刷新页面
      }
    },
  },
};
</script>

<style scoped>
</style>

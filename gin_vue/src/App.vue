<template>
  <div id="app">
    <b-container>
      <NavbarView/>
    </b-container>
    <router-view/>
  </div>
</template>
<script>

import NavbarView from './views/layout/NavbarView.vue';

export default ({
  components: { NavbarView },
  data() {
    return {};
  },
});
</script>

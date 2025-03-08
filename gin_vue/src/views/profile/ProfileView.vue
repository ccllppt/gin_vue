<template>
<h3>
hello<span v-if="userInfo">{{ userInfo.name }}</span>
</h3>
</template>

<script lang="ts">

import { mapState } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>

<style scoped>

</style>

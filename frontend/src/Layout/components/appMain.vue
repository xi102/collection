<template>
  <div>
    <router-view :key="key" />
  </div>
</template>

<script>
export default {
  name: "appMain",
  computed: {
    key() {
      return this.$route.path;
    },
  },
};
</script>

<style>
</style>
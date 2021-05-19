<template>
  <div>
    <appMain />
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link> |
      <router-link to="/testfiles">Upload Files</router-link> |
      <router-link to="/files_show">Files</router-link> |
      <router-link to="/login">Login</router-link>
    </div>
  </div>
</template>
<script>
import appMain from "./components/appMain.vue";
export default {
  name: "Layout",
  components: {
    appMain,
  },
  computed: {
    key() {
      return this.$route.path;
    },
  },
};
</script>
<style lang="scss">
#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
  position: absolute;
  width: 100%;
  align-items: center;
  bottom: 0;
}
</style>

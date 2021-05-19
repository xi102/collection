<template>
  <div id="layout">
    <v-app id="inspire">
      <v-app id="inspire">
        <v-navigation-drawer
          v-model="drawer"
          app
        >
        <v-list dense>
            <v-list-item link>
              <v-list-item-action>
                <v-icon>mdi-home</v-icon>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>Home</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-list-item link>
              <v-list-item-action>
                <v-icon>mdi-email</v-icon>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>Contact</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-navigation-drawer>
        <v-app-bar
          app
          color="#AADDAA"
          dark
        >
          <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
          <v-toolbar-title>仓库核心</v-toolbar-title>
        </v-app-bar>
        <v-main>
          <appMain />
        </v-main>
        <v-footer app>
          <div id="nav">
            <router-link to="/">Home</router-link> |
            <router-link to="/about">About</router-link> |
            <router-link to="/testfiles">Upload Files</router-link> |
            <!-- <router-link to="/files_show">Files</router-link> | -->
            <router-link to="/login">Login</router-link>
          </div>
        </v-footer>
      
      </v-app>
    </v-app>
  </div>
</template>
<script>
import appMain from "./components/appMain.vue";
//import Vue from 'vue/dist/vue.esm.js';
import Vuetify from 'vuetify';



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
  //el: '#app',
  vuetify: new Vuetify(),
  props: {
    source: String,
  },

  data: () => ({
    drawer: null,
  }),
};
</script>
<style lang="scss">
#nav {
  padding: 30px;
  margin: auto;
  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>

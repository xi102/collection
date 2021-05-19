<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-img
          :src="require('../assets/base.png')"
          class="my-3"
          contain
          height=500
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          欢迎来到 主楼西102仓库
        </h1>

        <p class="subheading font-weight-regular">
          这里有师兄师姐们的毕业资料和数据
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'HelloWorld',

    data: () => ({
      
    }),
  }
</script>

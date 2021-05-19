<template>
  <div>
    <h1 style="padding-top:120px;padding-bottom:20px;">师兄师姐们把文件放到这里来</h1>
    <v-container>
        <div style="margin:60px;">
            <v-file-input 
            filled
            show-size 
            counter 
            chips 
            multiple 
            label="点击此处选择文件" 
            ref="myfile" 
            v-model="files"
            />
            <v-btn text @click="submitFiles" class="mt-10"><h3>提交上传</h3></v-btn>
        </div>
    </v-container>
    <v-dialog
        v-model="dialog"
        width="500"
    >
        <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="red lighten-2"
            dark
            v-bind="attrs"
            v-on="on"
        >
            Click Me
        </v-btn>
        </template>

        <v-card>
        <v-card-title class="headline grey lighten-2">
            提示
        </v-card-title>

        <v-card-text>
            恭喜你，上传成功了
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
            color="primary"
            text
            @click="dialog = false"
            >
            好的
            </v-btn>
        </v-card-actions>
        </v-card>
    </v-dialog>
  </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            files: null,
        }
    },
    methods: {
        submitFiles(){
            let formData = new FormData()

            if(this.files){

                for( let file of this.files){
                    formData.append("file", file, file.name)
                }

                console.log(formData.getAll("file"))
                console.log(this.files)
                //POST请求向后端用form方式传文件
                axios.post('/upload/', formData)
                      .then( response => {
                                console.log('Success!')
                                console.log({response})
                            })
                      .catch(error => {
                                console.log({error})
                            })
            }
            else {
                console.log('there are no files.')
            }
        }
    }
}

</script>
<style scoped>
.container {
    max-width: 800px;
}
</style>
<template>
  <div>
    <v-flex>
      <v-file-input show-size counter chips multiple label="Arquivo Geral" ref="myfile" v-model="files"></v-file-input>
      <v-btn color="primary" text @click="submitFiles">test</v-btn>
    </v-flex>
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
                axios.post('/upload', formData)
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
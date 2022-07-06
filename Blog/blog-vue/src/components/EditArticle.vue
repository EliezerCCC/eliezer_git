<template>
    <div class="hello">
      <el-dialog title="编辑文章信息" :visible.sync="dialogEdit.show">
        <el-form :model="form" :rules="rules" ref="formEdit" label-width="100px" >
          <el-form-item label="标题" >
            <el-input v-model="form.recordtitle"></el-input>
          </el-form-item>
          <el-form-item label="内容" >
            <el-input v-model="form.recordcontent"   type="textarea"
            :rows="20"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogEdit.show = false;canel();">取消</el-button>
          <el-button type="primary" @click="dialogEdit.show = false;confirm(); ">确定</el-button>
        </div>
      </el-dialog>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HelloWorld',
    props:{
      dialogEdit:Object,
      form:Object,
    },
    data () {
      return {
        rules: {
            recordtitle: [
            { required: true, message: '请输入标题', trigger: 'blur' },
            { min: 2, max: 20, message: '长度在 2到 20 个字符', trigger: 'blur' }
          ]
        }
      }
    },
    methods:{
            confirm(){
                this.axios({
              url:this.global.apiUrl+'/updata', 
              data:this.form, 
              method:"PUT", 
              headers:{
                'Content-Type':'application/json',
                'Authorization': "Bearer " + localStorage.getItem('token'),
              }
            })
            .then(res=>{
                console.log(res.data);
                if(res.data.code==2003||res.data.code==2004||res.data.code==2005){
                    alert('请先完成登录!');
                    this.$router.push('/')
                }
            })
            .catch(Error=>{
                console.log(Error)
            })
            
               
            },
            canel(){
                this.$emit('updateEdit')
            }
    }
  }
  </script>
  

  <style scoped>
  
  </style>
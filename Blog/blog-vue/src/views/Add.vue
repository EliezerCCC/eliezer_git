<template>
    <div>
        <HeaderTop></HeaderTop>

        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-top: 120px; margin-left: auto;margin-right: auto;max-width: 550px;">
            <el-form-item label="标题" prop="recordtitle">
              <el-input v-model="ruleForm.recordtitle"></el-input>
            </el-form-item>
            <el-form-item label="内容" prop="recordcontent" >
                <el-input v-model="ruleForm.recordcontent"   type="textarea"
                :rows="20"
                placeholder="请输入内容"
                ></el-input>
              </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="Add()" style="margin-left:180px;">发表</el-button>
            </el-form-item>
          </el-form>

    </div>
</template>

<script>
export default {
    data() {
      return {
        ruleForm: {
            recordtitle: '',
            recordcontent:'',
            recorduserid:'',
            recordusername:'',
            createtime:'',
        },
        rules: {
            recordtitle: [
            { required: true, message: '请输入标题', trigger: 'blur' },
            { min: 2, max: 20, message: '长度在 2到 20 个字符', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
        Add(){
            this.axios({
              url:this.global.apiUrl+'/add', 
              data:this.ruleForm, 
              method:"POST", 
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
                }else{
                    alert('发表成功!');
                    this.$router.push('/home')
                }
            })
            .catch(Error=>{
                console.log(Error)
            })
            
        },
        created(){
    	//页面加载时就从本地通过localstorage获取存储的token值
        this.token =  localStorage.getItem('token')
    },
    }
  }
</script>

<style>

</style>
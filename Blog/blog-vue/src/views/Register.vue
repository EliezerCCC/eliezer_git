<template>
    <div>
        <el-container>
            <el-header>Register</el-header>
            <el-main>
                <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-top: 200px; margin-left: auto;margin-right: auto;max-width: 350px;">
                    <el-form-item label="账号" prop="userid">
                      <el-input v-model="ruleForm.userid"></el-input>
                    </el-form-item>
                    <el-form-item label="昵称" prop="username">
                        <el-input v-model="ruleForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="userpassword">
                        <el-input v-model="ruleForm.userpassword"></el-input>
                    </el-form-item>
                    <el-form-item label="确认密码" prop="secondpassword">
                        <el-input v-model="ruleForm.secondpassword"></el-input>
                    </el-form-item>
                    <el-form-item>    
                      <el-button type="primary" @click="submitForm('ruleForm')">注册</el-button>
                      <el-button @click="ToLogin()">返回</el-button>
                    </el-form-item>
                  </el-form>
            </el-main>
          </el-container>
    </div>
</template>

<script>
export default {
    data() {
      return {
        ruleForm: {
          userid: '',
          userpassword:'',
          username:'',
          secondpassword:'',
        },
        rules: {
            userid: [
            { required: true, message: '请输入账号', trigger: 'blur' },
            { min: 6, max: 11, message: '长度在 6 到 11 个字符', trigger: 'blur' }
          ],
          userpassword: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
          ],
          username: [
          { required: true, message: '请输入昵称', trigger: 'blur' },
          { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if(this.ruleForm.userpassword!=this.ruleForm.secondpassword){
                alert('密码不一致!');
              }else{
                this.axios({
              url:this.global.apiUrl+'/register', 
              data:{
                userid:this.ruleForm.userid,
                username:this.ruleForm.username,
                userpassword:this.ruleForm.userpassword,
              }, 
              method:"POST", 
              header:{
                'Content-Type':'application/json' 
              }
            })
            .then(res=>{
                console.log(res.data)
                if(res.data.result=="账号已存在"){
                  alert('账号已存在!');
                }else{
                  alert('注册成功!');
                  this.$router.push('/')
                }
            })
            .catch(Error=>{
                console.log(Error)
            })
              }
            
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      ToLogin() {
        this.$router.push('/')
      }
    }
  }
</script>

<style>
    .el-header, .el-footer {
      background-color: #a3fdf1;
      color: #333;
      text-align: center;
      line-height: 60px;
    }
    
    .el-aside {
      background-color: #D3DCE6;
      color: #333;
      text-align: center;
      line-height: 200px;
    }
    
    .el-main {
      background-color: #ffffff;
      color: #333;
      text-align: center;
      line-height: 800px;
    }
    
    body > .el-container {
      margin-bottom: 40px;
    }
    
    .el-container:nth-child(5) .el-aside,
    .el-container:nth-child(6) .el-aside {
      line-height: 260px;
    }
    
    .el-container:nth-child(7) .el-aside {
      line-height: 320px;
    }
  </style>
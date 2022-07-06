<template>
    <div>
        <el-container>
            <el-header>Login</el-header>
            <el-main>
                <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-top: 200px; margin-left: auto;margin-right: auto;max-width: 350px;">
                    <el-form-item label="账号" prop="userid">
                      <el-input v-model="ruleForm.userid"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="userpassword" >
                        <el-input v-model="ruleForm.userpassword" show-password></el-input>
                      </el-form-item>
                    <el-form-item>
                      <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
                      <el-button @click="ToRegister()">注册</el-button>
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
        },
        rules: {
            userid: [
            { required: true, message: '请输入账号', trigger: 'blur' },
            { min: 6, max: 11, message: '长度在 6 到 11 个字符', trigger: 'blur' }
          ],
          userpassword: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.axios({
              url:this.global.apiUrl+'/login', 
              data:this.ruleForm, 
              method:"POST", 
              header:{
                'Content-Type':'application/json' 
              }
            })            
            .then(res=>{
                console.log(res.data)
                if(res.data.result=="账号或密码错误"){
                  alert('账号或密码错误!');
                }else{
                  alert('登录成功!');
                  localStorage.setItem("token",res.data.token)
                  this.$router.push('/home')
                }
            })
            .catch(Error=>{
                console.log(Error)
            })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      ToRegister() {
        this.$router.push('/register')
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
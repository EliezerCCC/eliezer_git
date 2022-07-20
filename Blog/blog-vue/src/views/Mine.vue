<template>
    <div>
        <HeaderTop></HeaderTop>
        <el-container direction="vertical"  >
            <el-row type="flex" justify="center" >
              <el-col >
                <div class="blog-box" v-for="blog in blogs"  style="background-color: rgb(220, 235, 245);width: 800px;margin-left: auto;margin-right: auto;border-radius: 12px;" >
                  <h1 style="text-align:center" @click="goChange(blog)"><router-link :to="'/article'">{{blog.recordtitle}}</router-link></h1>
                  <span style="margin-left:20px;display :inline-block;width:90%;overflow:hidden;text-Overflow:ellipsis;white-Space:nowrap;">
                    {{blog.recordcontent}}
                  </span>
                  <el-row style="font-size:small;margin-left: 650px;">
                    作者：{{blog.recordusername}}
                  </el-row>
                  <el-row style="font-size:small;margin-left: 650px;">
                    创建时间：{{blog.createtime.slice(0, 10)}}
                  </el-row>
                </div>
              </el-col>
            </el-row>
          </el-container>
    </div>
</template>


<script>
    export default {
        data() {
            return {
                blogs: [],
            }
        },
        methods:{

// 跳转到传递参数页
goChange(row) {
  this.global.hid = true
  this.global.rid=row.recordid
  console.log(this.global.rid)
  this.$router.push({
    path: '/article',
    query: {
      blog:row,
      hid:true,
    }// 要传递的参数
  })
    }
  },
        created:function() {//创建完成
		this.axios({
            method: "GET",
            url: this.global.apiUrl+'/oneall',
            headers:{
                'Authorization': "Bearer " + localStorage.getItem('token'),
              }
        }).then(res => {
            console.log(res)
            if(res.data.code==2003||res.data.code==2004||res.data.code==2005){
                    alert('请先完成登录!');
                    this.$router.push('/')
                }else{
                  localStorage.setItem("token",res.data.token)
                    this.blogs=res.data.recordlist;
                }

        }).catch(err => {

        })
  },
        
}
</script>

<style>

</style>
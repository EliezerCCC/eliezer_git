<template>
    <div>
        <HeaderTop></HeaderTop>
        <el-row> <h1 style="text-align: center;">{{this.blog.recordtitle}}</h1></el-row>
        <el-row style=" width: 80%;margin-left: auto;margin-right: auto;">{{this.blog.recordcontent}}</el-row>
        <el-row style="margin-left:1600px;">作者:{{this.blog.recordusername}}</el-row>
        <el-row style="margin-left:1600px;">{{this.blog.createtime.slice(0,10)}}</el-row>
        <el-row v-if="this.global.hid" style="margin-left:1600px;">
            <el-button 
            size="mini"
            @click="handleEdit()">编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete()">删除</el-button>
        </el-row>
        <EditArticle :dialogEdit="dialogEdit" :form="form" @updateEdit="getInfo" ></EditArticle>
    </div>
</template>

<script>
export default {
    data() {
        return {
            blog: {},
            buttonhid: true,
            form: {},
            dialogEdit: {
                show: false,
            },
            rid:3,
        };
    },
    created:function() {//创建完成

  },
    watch: {
    '$route': 'gettingData'
  },
    methods: {
        getInfo(){
            this.blog.recordtitle= localStorage.getItem('title')
            this.blog.recordcontent= localStorage.getItem('content')
      },
        gettingData() {
      var queryCardNoFirst = this.$route.query.blog
      var his = this.$route.query.hid
      localStorage.setItem("title",queryCardNoFirst.recordtitle)
      localStorage.setItem("content",queryCardNoFirst.recordcontent)
      this.blog = queryCardNoFirst
      this.form = queryCardNoFirst
      this.buttonhid=this.$route.query.hid
           
      this.axios({
            method: "POST",
            url: this.global.apiUrl+'/getone',
            headers:{
                'Authorization': "Bearer " + localStorage.getItem('token'),
              },
            data:{
                recordid:this.global.rid,
            }
        }).then(res => {
            console.log(res)
            if(res.data.code==2003||res.data.code==2004||res.data.code==2005){
                    alert('请先完成登录!');
                    this.$router.push('/')
                }else{
                  localStorage.setItem("token",res.data.token)
                this.blog=res.data.record;
                }

        }).catch(err => {

        })

        
    },
        handleEdit() {
            this.dialogEdit.show = true;
        },
        handleDelete() {
            this.blog = this.$route.query.blog;
            console.log(localStorage.getItem("token"));
            this.$confirm("此操作将永久删除该文章,是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            }).then(() => {
                this.axios({
                    url: this.global.apiUrl + "/delete",
                    data: this.blog,
                    method: "DELETE",
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + localStorage.getItem("token"),
                    }
                })
                    .then(res => {
                        if(res.data.code==2003||res.data.code==2004||res.data.code==2005){
                    alert('请先完成登录!');
                    this.$router.push('/')
                }else{
                    localStorage.setItem("token",res.data.token)
                                    this.$message({
                    type: "success",
                    message: "删除成功!"
                });
                this.$router.push("/mine");
                }
                    console.log(res.data);
                })
                    .catch(Error => {
                    console.log(Error);
                });

            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "已取消删除"
                });
            });
        },
    }
}
</script>

<style>

</style>
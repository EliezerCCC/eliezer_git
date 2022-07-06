<template>
<div>
      <el-menu
        class="el-menu-demo"
        mode="horizontal"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b">
        <el-submenu index="1" >
          <template slot="title">选项</template>
          <el-menu-item index="1-1" @click="ToAll()">全部博客</el-menu-item>
          <el-menu-item index="1-2" @click="ToMine()">我的博客</el-menu-item>
          <el-menu-item index="1-3" @click="ToAdd()">发表博客</el-menu-item>
        </el-submenu>
      </el-menu>
</div>
</template>


<script>
  export default {
    methods:{
      ToAll(){
        this.$router.push('/home')
      },
      ToMine(){
        this.$router.push('/mine')
      },
      ToAdd(){
        this.$router.push('/add')
      }
    }
  }
</script>

<style>

</style>
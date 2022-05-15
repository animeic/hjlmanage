<template>
  <div class="login-container">
      <el-form
      ref="loginFrom"
      :model="LoginFrom"
      :rules="rules"
      label-width="80px"
      >
        <p>login</p>
        <el-form-item label="账号：" prop="username">
            <el-input v-model="LoginFrom.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码：" prop="password">
            <el-input v-model="LoginFrom.password" type="password" placeholder="请输入密码" show-password></el-input>
        </el-form-item>
        <el-form-item class="login-button">
          <el-button type="success" @click="submitForm('loginFrom')">登录</el-button>
        </el-form-item>
        <el-form-item class="register-button">
          <el-button type="danger" @click="toregister">去注册</el-button>
        </el-form-item>
          <el-row>
          <el-col el-col :span="20"><div class="grid-content bg-purple-dark"></div></el-col>
          <el-col el-col :span="4"><el-link href="/repass" type="primary" style="float: right;font-size: 12px;">忘记密码 ？</el-link></el-col>
        </el-row>
      </el-form>
    <!-- </div> -->
  </div>
</template>

<script setup>

// import http from '@/http/index'
import { reactive,getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage,ElLoading } from 'element-plus'
// import $ from '@/'
// 注意：1. 双向绑定必须ref,2. prop属性是验证属性 2.model rules ref
const LoginFrom = reactive({
  username: '',
  password: '',
})

// prop
const rules = reactive({
  username: [
    {required:true,message: '必须填写账号',trigger: 'blur'},
    {min: 6,max: 12,message: '账号长度为6~12个字符',trigger: 'blur'}
  ],
  password: [
    {required:true,message: '必须填写密码',trigger: 'blur'},
    {min: 6,max: 12,message: '密码长度为6~12个字符',trigger: 'blur'}
  ],
})
const router = useRouter()
const cxt = getCurrentInstance()

const toregister = () => {
  router.push('/register')
}
const {data:result,pending:delay} = await useFetch("https://api.animeii.tech/auth/login",{
  method: 'POST',
  params: LoginFrom,
})




let loading
const startLoading = () => {
    const options = {
        lock:true,
        text: "加载中。。。",
        background: 'rgba(0,0,0,0.7)'
    }
    loading = ElLoading.service(options)
}
const endLoading = () => {
    loading.close()
}



const  submitForm = (cf) => {
  // 表单验证，成功提交数据
  
  cxt.refs[cf].validate((valid) => {
      if (valid) {
        console.log(result.value)
        if (delay){
          startLoading()
        }
        endLoading()
        if (result.value.code > 0) {
          ElMessage.success(result.value.message)
          router.push('/')
        }
       
        // http.post("/auth/login",LoginFrom).then(res=>{
        //   if(res.data.code === 0){
        //     const resdata = res.data.data
        //     const access_token = resdata.token_type + " " + resdata.access_token
        //     // const token_type = resdata.token_type
        //     // const expires_in = resdata.expires_in
        //     localStorage.setItem("token",access_token)
        //     router.push("/")
        //     return
        //   }
        //   return
        // })

      } else {
      //  $.dispatch("getToken")
      // console.log($.state.tokenInfo)
        return false
      }
    })
}

</script>

<style lang="scss" scoped>
  .el-form{
    display: grid;
    p{ 
      justify-self: center;
      color: green;
      font-size: 20px;
    }
    margin:60px 40px 40px 20px; // 边框间距效果
    grid-area: 2/1/3/2;
    :deep(.el-form-item){
      margin-top: 5px;
    }

    :deep(.el-form-item:nth-child(-n+2)){
      display: grid;
      grid-template-columns: 1fr 5fr;
      .el-form-item__label{
        grid-area: 1/1/2/2;
      }
      .el-form-item__content{
        grid-area: 1/2/2/3;
      }
    }
    .login-button{
      margin-bottom: 5px;
      :deep(.el-form-item__content){
        display: grid;
      }
    }
    .register-button{
      margin-bottom: 10px;
      :deep(.el-form-item__content){
        display: grid;
        
      }
    }

  }
  
</style>

<!-- <template>
    <Header></Header>
    <div id="register-container">
      <div class="register-form">
        <div class="register-text"><span>注册账号</span></div>
        <el-form
        ref="registerFrom"
        :model="RegisterFrom"
        :rules="rules"
        label-width="80px"
        >
          <el-form-item label="账号：" prop="username">
              <el-input v-model="RegisterFrom.username"></el-input>
          </el-form-item>
          
          <el-form-item label="邮箱：" prop="email">
              <el-input v-model="RegisterFrom.email">
              </el-input>
          </el-form-item>
           <el-form-item label="验证码：" prop="ecode">
              <el-input v-model="RegisterFrom.ecode">
                <template #append>
                    <el-button @click="sendCode">发送验证码</el-button>
                </template>
              </el-input>
          </el-form-item>
          <el-form-item label="密码：" prop="password">
              <el-input v-model="RegisterFrom.password" type="password" show-password></el-input>
          </el-form-item>
          <el-form-item class="register-button">
            <el-button type="success" @click="submitForm('registerFrom')">注册</el-button>
          </el-form-item>
          <el-form-item class="login-button">
            <el-button type="danger" @click="tologin">去登录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <Footer></Footer>
</template>

<script setup>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import http from '@/http/index'
import { ref,reactive,getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'
import { resultProps } from 'element-plus'

const RegisterFrom = reactive({
  ecode: '',
  username: '',
  password: '',
  email: '',
  etype: 'register'
})

// prop
const rules = reactive({
  ecode: [
    {required:true,message: '必须填写验证码',trigger: 'blur'},
  ],
  username: [
    {required:true,message: '必须填写账号',trigger: 'blur'},
  ],
  password: [
    {required:true,message: '必须填写密码',trigger: 'blur'},
  ],
  email: [
    {required:true,message: '必须填写邮箱',trigger: 'blur'},
    // {min: 6,max: 6,message: '口令必须6个字符',trigger: 'blur'}
  ],
})
const router = useRouter()
const cxt = getCurrentInstance()

// todo 调试很久找到这个原因，大坑：reactive 如何赋值给响应对象
const sendCode = () =>{
 http.post("/auth/sendCode",RegisterFrom)
}
const tologin = () => {
  router.push('/login')
}

const  submitForm = (cf) => {
  cxt.refs[cf].validate((valid) => {
    if (valid) {
      http.post("/auth/register",RegisterFrom).then(res=>{
        if(res.data.code === 0){
          const resdata = res.data.data
          const access_token = resdata.token_type + " " + resdata.access_token
          // const token_type = resdata.token_type
          // const expires_in = resdata.expires_in

          localStorage.setItem("token",access_token)
          router.push("/")
        }
      })
    }
  })
}

</script>

<style lang="scss" scoped>
// :deep() nth-child(-n+x) !important
#register-container{
  // debug
  min-height: 900px;
  // grid
  display: grid;
  justify-content: center;
  .register-form{
    display: grid;
    grid-template-rows: 80px 1fr;
    .register-text{
      margin-top: 50px;
      margin-bottom: 20px;
      grid-area: 1/1/2/2;
      place-self: center;
    }
    .el-form{
      grid-area: 2/1/3/2;
      :deep(.el-form-item:nth-child(-n+4)){
        display: grid;
        grid-template-columns: 1fr 5fr;
        .el-form-item__label{
          grid-area: 1/1/2/2;
        }
        .el-form-item__content{
          grid-area: 1/2/2/3;
        }
      }
      .register-button{
        margin-bottom: 5px;
        :deep(.el-form-item__content){
          display: grid;
        }
      }
      .login-button{
        :deep(.el-form-item__content){
          display: grid;
        }
      }

    }
  }
  

  

}
</style> -->
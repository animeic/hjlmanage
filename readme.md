# hjl-manage

## frontend

安装nuxt3 element-plus scss  

```shell
# 安装nuxt3
npx nuxi init nuxt-app
# 安装依赖
yarn install
# 开发运行
yarn dev

# 安装element-plus
yarn add element-plus
# 全局引入，按需引入参考:https://element-plus.org/zh-CN/guide/quickstart.html
    # 在 plugins/element.ts中添加
    `
    import ElementPlus from 'element-plus/dist/index.full'

    export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.use(ElementPlus)
    })
    `
    # app.vue中引入样式
    `
    <style>
    @import 'element-plus/dist/index.css'
    </style>
    `

# 安装scss
yarn add sass node-sass sass-loader

# 编译
npx nuxi generate

```

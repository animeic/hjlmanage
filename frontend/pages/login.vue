<template>
    <el-container>
        <el-header class="header-con">
            <Header class="hjl-header" />
        </el-header>

        <el-main class="main-container">
           <login-form class="hjl-main"/>
        </el-main>
        
    </el-container>
</template>

<script setup>


</script>

<style lang="scss" scoped>
// .el-container {
//     background-color: #2d90c3;
// }
.header-con {
    display: grid;
    align-items: center;
    border: 1px solid red;
}
.hjl-header{
    // display: grid;
    border:1px solid green;
}


.main-container{
    display:grid;
    place-items: center; //定位子元素
    min-height: 800px; // 保证撑满屏幕
}

.hjl-main{// 和组件是同一个
    display: grid;
    align-items: center;
    border-radius: 8px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

</style>

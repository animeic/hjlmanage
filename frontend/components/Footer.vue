<template>
    <el-row class="header" justify="center">
        <el-col :span="12" class="item1">
            <p>footer</p>
        </el-col>

    </el-row>

</template>


<style lang="scss" scoped>
.header{
    .item1{
        display:grid;
        border:1px solid red;
        place-items:center;
    }
    border-top: 1px solid black;
}

</style>

<script setup>

</script>
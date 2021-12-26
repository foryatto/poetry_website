<template>
    <div class="header">
        <a-row>
            <a-col :xs="2" :sm="2" :md="2" :lg="2" :xl="2">
                <router-link :to="{name:'Home'}">
                    <a-icon type="home" theme="twoTone" :style="{ fontSize: '20px'}" />
                </router-link>
            </a-col>
            <a-col :xs="16" :sm="12" :md="19" :lg="19" :xl="19">
                <a-input-search placeholder="搜索：名称" style="width: 200px" @search="onSearch" />
            </a-col>
            <!-- <a-col :xs="2" :sm="4" :md="3" :lg="3" :xl="3">
                <a-button type="primary" @click="login">
                    登录
                </a-button>
            </a-col> -->
        </a-row>
    </div>
</template>

<script>
    export default {
        data(){
            return {
                keyword: ''
            }
        },
        methods:{
            login(){
                this.$message.info('登录功能即将上线，登录之后您将可以收藏诗词，发布诗词赏析。');
            },
            onSearch(keyword){
                this.$router.push({path:"/search_result", query:{keyword: keyword}});
            }
        }
    }
</script>

<style>
    .header {
        margin-left: 16px;
    }
</style>
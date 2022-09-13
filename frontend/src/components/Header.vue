<script setup>
import { ref, inject } from 'vue'
import { useRouter } from 'vue-router';
import { HomeOutlined } from '@ant-design/icons-vue';
import { InputSearch, Row, Col } from 'ant-design-vue'

const keyword = ref('')
const router = useRouter()
function gotoSearchPage() {
    if (keyword.value != '') {
        router.push({ path: 'search', query: { 'keyword': keyword.value } })
        // get useRoute().query.keyword
        keyword.value = ""
    }
}

const { apiBaseUrl } = inject('appConfig')
function searchPoetry() {
    fetch(api + "/poem/search", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'keyword': keyword.value,
            'page': 1,
            'pageSize': 50,
        })
    })
        .then(response => response.json())
        .then(json => {
            console.log(json)
        })
        .catch(error => console.log(error))
}


</script>

<template>
    <div class="header">
        <div class="homeButton">
            <RouterLink to="/">
                <home-outlined :style="{'fontSize':'20px'}" />
            </RouterLink>
        </div>
        <div class="searchButton">
            <InputSearch v-model:value="keyword" placeholder="搜索诗词" size="large" @search="gotoSearchPage" />
        </div>
    </div>
</template>

<style scoped>
.header {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.header .homeButton {
    margin-left: 20px;
    margin-right: 10px;
}

.header .searchButton {
    margin-right: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>
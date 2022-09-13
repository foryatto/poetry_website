<script setup>
import { inject, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router';

import { Menu, MenuItem } from 'ant-design-vue'
import { ReadOutlined } from '@ant-design/icons-vue';


// 获取朝代列表
const dynasties = ref([''])
const { apiBaseUrl } = inject('appConfig')

function getDynasties() {
    fetch(apiBaseUrl + '/dynasty')
        .then(response => response.json())
        .then(json => {
            dynasties.value = json.data.list
        })
        .catch(error => console.log(error))
}

const router = useRouter()
function gotoPoetPage(item) {
    router.push({ path: 'poets', query: { 'dynasty_id': item.key } })
}


onMounted(() => {
    getDynasties()
})

</script>

<template >
    <div class="logo">
        中华诗词
    </div>
    <Menu mode="inline" @click="gotoPoetPage">
        <MenuItem v-for="dynasty in dynasties" :key="dynasty.id">
        <ReadOutlined />
        <span class="nav-text">{{ dynasty.name }}</span>
        </MenuItem>
    </Menu>
</template>

<style scoped>
.logo {
    height: 32px;
    background: rgba(182, 21, 21, .2);
    margin: 16px;
    text-align: center;
    line-height: 32px;
    font-weight: bolder;
}
</style>
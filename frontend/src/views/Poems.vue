<script setup>
import { ref, inject, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';

import { Divider, List, ListItem, ListItemMeta, Row, Col, Pagination } from 'ant-design-vue';

const { apiBaseUrl } = inject('appConfig')
const route = useRoute()

const pageCurrent = ref(1)
const pageSize = ref(10)

var poetId = route.query.poet_id
var poetName = route.query.poet_name
const poems = ref([
    {
        'profile': '',
        'dynasty': { 'name': '' }
    }
])
const poemCount = ref(0)
function getPoems() {
    fetch(apiBaseUrl + '/poem/brief', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "poetId": poetId,
            "page": 1,
            "pageSize": 10,
        })
    })
        .then(response => response.json())
        .then(json => {
            poems.value = json.data.list
            poemCount.value = json.data.total
        })
        .catch(error => console.log(error))
}

function onShowSizeChange(current, newPageSize) {
    pageSize.value = newPageSize
    fetch(apiBaseUrl + '/poem/brief', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "poetId": poetId,
            "page": current,
            "pageSize": newPageSize,
        })
    })
        .then(response => response.json())
        .then(json => {
            poems.value = json.data.list
            poemCount.value = json.data.total
        })
        .catch(error => console.log(error))
}

function pageChange(page, pageSize) {
    pageCurrent.value = page
    fetch(apiBaseUrl + '/poem/brief', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "poetId": poetId,
            "page": page,
            "pageSize": pageSize,
        })
    })
        .then(response => response.json())
        .then(json => {
            poems.value = json.data.list
            poemCount.value = json.data.total
        })
        .catch(error => console.log(error))
}

watch(() => route.query.poet_id, () => {
    poetId = route.query.poet_id
    poetName = route.query.poet_name
    getPoems()
})

onMounted(() => {
    getPoems()
})
</script>
    
<template>
    <div>
        <Divider>
            {{ poetName }}
        </Divider>
        <list item-layout="horizontal">
            <ListItem v-for="poem in poems">
                <ListItemMeta :description="poem.content">
                    <template #title>
                        <RouterLink :to="{path:'poem', query:{poem_id: poem.id} }">
                            {{ poem.name }}
                        </RouterLink>
                    </template>
                </ListItemMeta>
            </ListItem>
        </list>
        <br /><br />
        <Row>
            <Col :span="24">
            <Pagination show-size-changer :default-current="1" :total="poemCount" :current="pageCurrent"
                :page-size.sync="pageSize" @showSizeChange="onShowSizeChange" @change="pageChange" />
            </Col>
        </Row>

    </div>
</template>
    
    
<style scoped>

</style>
<script setup>
import { ref, inject, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';

import { Divider, List, ListItem, ListItemMeta, Row, Col, Pagination } from 'ant-design-vue';

const { apiBaseUrl } = inject('appConfig')
const route = useRoute()

const pageCurrent = ref(1)
const pageSize = ref(10)

var dynastyId = route.query.dynasty_id
const poets = ref([
    {
        'profile': '',
        'dynasty': { 'name': '' }
    }
])
const poetCount = ref(0)
function getPoets(page, pageSize) {
    fetch(apiBaseUrl + '/poet', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "dynastyId": dynastyId,
            "page": page,
            "pageSize": pageSize,
        })
    })
        .then(response => response.json())
        .then(json => {
            poets.value = json.data.list
            poetCount.value = json.data.total
        })
        .catch(error => console.log(error))
}

function onShowSizeChange(current, newPageSize) {
    pageSize.value = newPageSize
    getPoets(current, newPageSize)
}

function pageChange(page, pageSize) {
    pageCurrent.value = page
    getPoets(page, pageSize)
}

watch(() => route.query.dynasty_id, () => {
    dynastyId = route.query.dynasty_id
    getPoets(1, 10)
})

onMounted(() => {
    getPoets(1, 10)
})
</script>

<template>
    <div>
        <Divider>
            {{ poets[0].dynasty.name }}
        </Divider>
        <List item-layout="horizontal">
            <ListItem slot="renderItem" v-for="poet in poets">
                <ListItemMeta :description="poet.profile">
                    <template #title>
                        <RouterLink :to="{path:'poems', query:{poet_id: poet.id, poet_name: poet.name} }">
                            {{ poet.name }}
                        </RouterLink>
                    </template>
                </ListItemMeta>
            </ListItem>
        </List>
        <br /><br />
        <Row>
            <Col :span="24">
            <Pagination show-size-changer :default-current="1" :total="poetCount" :current="pageCurrent"
                :page-size.sync="pageSize" @showSizeChange="onShowSizeChange" @change="pageChange"
                :pageSizeOptions="['10','20','30','40','50']" />
            </Col>
        </Row>
    </div>
</template>


<style scoped>

</style>
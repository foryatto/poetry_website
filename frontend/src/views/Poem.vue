<script setup>
import { onMounted, ref, inject, watch } from 'vue';
import { useRoute } from 'vue-router';
import { Row, Col, Divider, Card } from 'ant-design-vue';
import { EditOutlined } from '@ant-design/icons-vue';

const { apiBaseUrl } = inject('appConfig');
const route = useRoute()

var poemId = route.query.poem_id
const poem = ref({
    'poet': { 'name': '', 'dynasty': { 'name': '' } },
});

function getPoem() {
    fetch(apiBaseUrl + "/poem/detail", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "poemId": poemId,
        })
    })
        .then(response => response.json())
        .then(json => {
            poem.value = json.data.poem
        })
        .catch(error => console.log(error))
}

watch(() => route.query.poem_id, () => {
    poemId = route.query.poem_id
    getPoem()
})

onMounted(() => {
    getPoem()
})

</script>

<template>
    <div>
        <Divider>
            {{ poem.name }}
        </Divider>
        <Row :gutter="16">

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Card :title="poem.name" :bordered="true" style="margin-bottom:20px">
                <p>{{ poem.content }}</p>
            </Card>
            </Col>

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Card :bordered="true" :title="poem.poet.name">
                <p>{{ poem.poet.dynasty.name }}诗人</p>
                {{ poem.poet.profile }}
                <a>
                    <EditOutlined />纠错
                </a>
            </Card>
            </Col>

        </Row>
    </div>
</template>

<style scoped>

</style>
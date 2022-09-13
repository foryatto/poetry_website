<script setup>
import { onMounted, ref, inject } from 'vue';
import { Row, Col, Statistic, Divider, Card } from 'ant-design-vue';
import { EditOutlined } from '@ant-design/icons-vue';

const randomPoem = ref({
    'poet': { 'name': '', 'dynasty': { 'name': '' } },
});

// 获取一首随机诗词
const { apiBaseUrl } = inject('appConfig');
function getPoemByRandom() {
    fetch(apiBaseUrl + "/poem/today")
        .then(response => response.json())
        .then(json => {
            randomPoem.value = json.data.poem
        })
        .catch(error => console.log(error))
}

onMounted(() => {
    getPoemByRandom()
})

</script>

<template>
    <div style="white-space: pre-line;">
        <Row :gutter="16">

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Statistic title="已收录诗词" :value="53075" style="margin-right: 50px">
                <template #suffix>
                    <span> 首</span>
                </template>
            </Statistic>
            </Col>

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Statistic title="诗人" :value="4459" class="demo-class">
                <template #suffix>
                    <span> 人</span>
                </template>
            </Statistic>
            </Col>
        </Row>

        <Divider orientation="left">今日诗词</Divider>
        <Row :gutter="16">

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Card :title="randomPoem.name" :bordered="true" style="margin-bottom:20px">
                <p>{{ randomPoem.content }}</p>
            </Card>
            </Col>

            <Col :span="12" :xs='24' :sm='24' :md='12'>
            <Card :bordered="true" :title="randomPoem.poet.name">
                <p>{{ randomPoem.poet.dynasty.name }}诗人</p>
                {{ randomPoem.poet.profile }}
                <a>
                    <EditOutlined />纠错
                </a>
            </Card>
            </Col>

        </Row>
    </div>
</template>
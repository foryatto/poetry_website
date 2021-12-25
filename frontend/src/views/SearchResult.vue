<template>
    <div>
        <a-divider>
            搜索结果
        </a-divider>
        <a-list item-layout="horizontal" :data-source="poems">
            <a-list-item slot="renderItem" slot-scope="poem">
                <a-list-item-meta :description="poem.content">
                    <router-link :to="{name:'PoemDetail', query:{poem_id: poem.id} }" slot="title">{{ poem.name }}
                    </router-link>
                </a-list-item-meta>
            </a-list-item>
        </a-list>
        <!-- <br /><br />
        <a-row>
            <a-col :span="24">
                <a-pagination show-size-changer :default-current="1" :total="poemCount" :current="pageCurrent"
                    @showSizeChange="onShowSizeChange" @change="pageChange" />
            </a-col>
        </a-row> -->

    </div>
</template>


<script>
    export default {
        data() {
            return {
                keyword: this.$route.query.keyword,
                poems: {},
            }
        },
        methods: {
            getPoemsByKeyWord() {
                this.axios.get('/poems/search?keyword=' + this.keyword)
                    .then(response => this.poems = response.data)
                    .catch(error => console.log(error))
            }
        },
        mounted() {
            this.getPoemsByKeyWord()
        },
        watch: {
            $route: {
                immediate: true,
                handler() {
                    this.keyword = this.$route.query.keyword
                    this.getPoemsByKeyWord()
                }
            }
        }
    }
</script>

<style scoped>

</style>
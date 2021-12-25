<template>
    <div>
        <a-divider>
            {{ poems[0].poet.name }}
        </a-divider>
        <a-list item-layout="horizontal" :data-source="poems">
            <a-list-item slot="renderItem" slot-scope="poem">
                <a-list-item-meta :description="poem.content">
                    <router-link :to="{name:'PoemDetail', query:{poem_id: poem.id} }" slot="title">{{ poem.name }}
                    </router-link>
                </a-list-item-meta>
            </a-list-item>
        </a-list>
        <br /><br />
        <a-row>
            <a-col :span="24">
                <a-pagination show-size-changer :default-current="1" :total="poemCount" :current="pageCurrent"
                    :page-size.sync="pageSize" @showSizeChange="onShowSizeChange" @change="pageChange" />
            </a-col>
        </a-row>

    </div>
</template>


<script>
    export default {
        data() {
            return {
                poet_id: this.$route.query.poet_id,
                poems: {},
                poemCount: 0,
                pageCurrent: 1,
                pageSize: 10,
            }
        },
        methods: {
            getPoemsById() {
                this.axios.get('/poems/' + this.poet_id + '?skip=0&limit=10')
                    .then(response => this.poems = response.data)
                    .catch(error => console.log(error))
            },
            getPoemCount() {
                this.axios.get('/poems/count/' + this.poet_id)
                    .then(response => this.poemCount = response.data)
                    .catch(error => console.log(error))
            },
            onShowSizeChange(current, pageSize) {
                var skip = (current - 1) * pageSize
                this.pageSize = pageSize
                this.axios.get('/poems/' + this.poet_id + '?skip=' + skip + '&limit=' + pageSize)
                    .then(response => this.poems = response.data)
                    .catch(error => console.log(error))
            },
            pageChange(page, pageSize) {
                this.pageCurrent = page
                var skip = (page - 1) * pageSize
                this.axios.get('/poems/' + this.poet_id + '?skip=' + skip + '&limit=' + pageSize)
                    .then(response => this.poems = response.data)
                    .catch(error => console.log(error))
            }
        },
        mounted() {
            this.getPoemsById()
            this.getPoemCount()
        },
        watch: {
            $route: {
                immediate: true,
                handler() {
                    this.poet_id = this.$route.query.poet_id
                    this.pageCurrent = 1
                    this.pageSize = 10
                    this.getPoemsById()
                    this.getPoemCount()
                }
            }
        }
    }
</script>

<style scoped>

</style>
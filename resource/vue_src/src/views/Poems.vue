<template>
    <div>
        <a-divider>
            {{ poet_name }}
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
                poet_name: this.$route.query.poet_name
            }
        },
        methods: {
            getPoemsById() {
                this.axios.post('/poem/brief', {
                    "poetId": this.poet_id,
                    "page": 1,
                    "pageSize": 10
                })
                    .then(response => {
                        this.poems = response.data.data.list
                        this.poemCount = response.data.data.total
                    })
                    .catch(error => console.log(error))
            },
            onShowSizeChange(current, pageSize) {
                this.pageSize = pageSize
                this.axios.post('/poem/brief', {
                    "poetId": this.poet_id,
                    "page": current,
                    "pageSize": pageSize
                })
                    .then(response => {
                        this.poems = response.data.data.list
                        this.poemCount = response.data.data.total
                    })
                    .catch(error => console.log(error))
            },
            pageChange(page, pageSize) {
                this.pageCurrent = page
                this.axios.post('/poem/brief', {
                    "poetId": this.poet_id,
                    "page": page,
                    "pageSize": pageSize
                })
                    .then(response => {
                        this.poems = response.data.data.list
                        this.poemCount = response.data.data.total
                    })
                    .catch(error => console.log(error))
            }
        },
        mounted() {
            this.getPoemsById()
        },
        watch: {
            $route: {
                immediate: true,
                handler() {
                    this.poet_id = this.$route.query.poet_id
                    this.pageCurrent = 1
                    this.pageSize = 10
                    this.getPoemsById()
                }
            }
        }
    }
</script>

<style scoped>

</style>
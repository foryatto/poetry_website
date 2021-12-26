<template>
    <div>
        <a-divider>
            {{ poets[0].dynasty.name }}
        </a-divider>
        <a-list item-layout="horizontal" :data-source="poets">
            <a-list-item slot="renderItem" slot-scope="poet">
                <a-list-item-meta :description="poet.profile">
                    <router-link :to="{name:'Poems', query:{poet_id: poet.id, poet_name: poet.name} }" slot="title">{{ poet.name }}
                    </router-link>
                </a-list-item-meta>
            </a-list-item>
        </a-list>
        <br /><br />
        <a-row>
            <a-col :span="24">
                <a-pagination show-size-changer :default-current="1" :total="poetCount" :current="pageCurrent"
                    :page-size.sync="pageSize" @showSizeChange="onShowSizeChange" @change="pageChange" />
            </a-col>
        </a-row>
    </div>
</template>


<script>
    export default {
        data() {
            return {
                dynasty_id: this.$route.query.dynasty_id,
                poets: {},
                poetCount: 0,
                pageCurrent: 1,
                pageSize: 10,
            }
        },
        methods: {
            getPoets() {
                this.axios.post('/poets/', {
                    "dynastyId": this.dynasty_id,
                    "page": 1,
                    "pageSize": 10,
                })
                    .then(response => {
                        this.poets = response.data.data.items
                        this.poetCount = response.data.data.total
                    })
                    .catch(error => console.log(error))
            },
            onShowSizeChange(current, pageSize) {
                this.pageSize = pageSize
                this.axios.post('/poets/', {
                    "dynastyId": this.dynasty_id,
                    "page": current,
                    "pageSize": pageSize,
                })
                    .then(response => {
                        this.poets = response.data.data.items
                    })
                    .catch(error => console.log(error))
            },
            pageChange(page, pageSize) {
                this.pageCurrent = page
                this.axios.post('/poets/', {
                    "dynastyId": this.dynasty_id,
                    "page": page,
                    "pageSize": pageSize,
                })
                    .then(response => {
                        this.poets = response.data.data.items
                    })
                    .catch(error => console.log(error))
            }
        },
        mounted() {
            this.getPoets()
        },
        watch: {
            $route: {
                immediate: true,
                handler() {
                    this.dynasty_id = this.$route.query.dynasty_id
                    this.pageCurrent = 1
                    this.pageSize = 10
                    this.getPoets()
                }
            }
        }
    }
</script>

<style scoped>

</style>
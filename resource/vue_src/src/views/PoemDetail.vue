<template>
    <div>
        <a-divider>
            {{ poem.name }}
        </a-divider>
        <a-row :gutter="16">
      <a-col :span="12" :xs='24' :sm='24' :md='12'>
        <a-card :title="poem.name" :bordered="true" style="margin-bottom:20px">
          <p>{{ poem.content }}</p>
          <template slot="actions" class="ant-card-actions">
            <a-tooltip>
              <template slot="title">
                收藏
              </template>
              <a-icon key="star" type="star" />
            </a-tooltip>

            <a-tooltip>
              <template slot="title">
                纠错
              </template>
              <a-icon key="edit" type="edit" />
            </a-tooltip>
          </template>
        </a-card>
      </a-col>
      <a-col :span="12" :xs='24' :sm='24' :md='12'>
        <a-card :title="poem.poet.name" :bordered="true">
          <p>{{ poem.poet.dynasty.name }}诗人</p>
          {{ poem.poet.profile }}
          <a><a-icon key="edit" type="edit" />纠错</a>
        </a-card>
      </a-col>
    </a-row>
    </div>
</template>


<script>
    export default {
        data() {
            return {
                poem_id: this.$route.query.poem_id,
                poem: {},
            }
        },
        methods: {
            getPoemById() {
                this.axios.post('/poem/detail/', {
                  "poemId": this.poem_id,
                })
                    .then(response => this.poem = response.data.data.poem)
                    .catch(error => console.log(error))
            }
        },
        mounted() {
            this.getPoemById()
        },
        watch: {
            $route: {
                immediate: true,
                handler() {
                    this.poem_id = this.$route.query.poem_id
                    this.getPoemById()
                }
            }
        }
    }
</script>

<style scoped>

</style>
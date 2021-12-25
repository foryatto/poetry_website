<template>
  <div class="home" style="white-space: pre-line;">
    <a-row :gutter="16">
      <a-col :span="12" :xs='24' :sm='24' :md='12'>
        <a-statistic title="已收录诗词" :value="53075" style="margin-right: 50px">
          <template #suffix>
            <span> 首</span>
          </template>
        </a-statistic>
      </a-col>
      <a-col :span="12" :xs='24' :sm='24' :md='12'>
        <a-statistic title="诗人" :value="4459" class="demo-class">
          <template #suffix>
            <span> 人</span>
          </template>
        </a-statistic>
      </a-col>
    </a-row>

    <a-divider orientation="left">随机诗词</a-divider>
    <a-row :gutter="16">
      <a-col :span="12" :xs='24' :sm='24' :md='12'>
        <a-card :title="poemtoday.name" :bordered="true" style="margin-bottom:20px">
          <p>{{ poemtoday.content }}</p>
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
        <a-card :title="poemtoday.poet.name" :bordered="true">
          <p>{{ poemtoday.poet.dynasty.name }}诗人</p>
          {{ poemtoday.poet.profile }}
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
        poemtoday: {},
      }
    },
    methods: {
      getPoemtoday() {
        this.axios.get('/poems/random')
          .then(response => this.poemtoday = response.data)
          .catch(error => console.log(error))
      },
    },
    mounted() {
      this.getPoemtoday()
    }
  }
</script>
<template>
  <a-layout id="components-layout-demo-responsive">
    <a-layout-sider
      breakpoint="lg"
      collapsed-width="0"
      @collapse="onCollapse"
      @breakpoint="onBreakpoint"
      theme="light"
    >
    <SiderMenu />
    </a-layout-sider>
    <a-layout>
      <a-layout-header :style="{ background: '#fff', padding: 0 }" >
        <Header/>
      </a-layout-header>
      <a-layout-content :style="{ margin: '24px 16px 0'}" style="min-height:700px">
        <div :style="{ padding: '24px', background: '#fff', minHeight: '360px' }">
          <router-view/>
        </div>
      </a-layout-content>
      <a-layout-footer style="textAlign: center; padding-left:0px; padding-right:0px">
        <Footer />
      </a-layout-footer>
      <a-back-top />
    </a-layout>
  </a-layout>
</template>
<script>
import SiderMenu from './components/SiderMenu.vue'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'

export default {
  components:{
    SiderMenu,
    Header,
    Footer
  },
  methods: {
    onCollapse(collapsed, type) {
      console.log(collapsed, type);
    },
    onBreakpoint(broken) {
      console.log(broken);
    },
  },
};
</script>

<style>
#components-layout-demo-responsive .logo {
  height: 32px;
  background: rgba(182, 21, 21, 0.2);
  margin: 16px;
  text-align: center;
  line-height: 32px;
  font-weight: bolder;
}
</style>

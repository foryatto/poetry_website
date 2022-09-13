<script setup>
import { RouterView } from 'vue-router'

import { BackTop, Layout, LayoutSider, LayoutHeader, LayoutContent, LayoutFooter } from 'ant-design-vue'

import SideMenuVue from './components/SideMenu.vue';
import Footer from './components/Footer.vue';
import Header from './components/Header.vue';

const onCollapse = (collapsed, type) => {
  // console.log(collapsed, type);
};

const onBreakpoint = (broken) => {
  // console.log(broken);
};

</script>

<template>
  <Layout class="app">

    <LayoutSider class="sideBar" theme="light" breakpoint="lg" collapsed-width="0" @collapse="onCollapse" @breakpoint="onBreakpoint">
      <SideMenuVue />
    </LayoutSider>
    
    <Layout>
      <LayoutHeader :style="{ background: '#fff', padding: 0 }">
        <Header />
      </LayoutHeader>

      <LayoutContent :style="{ margin: '24px 16px 0', background: '#fff', padding: '20px 20px'}">
        <RouterView />
      </LayoutContent>

      <LayoutFooter style="text-align: center">
        <Footer />
      </LayoutFooter>
    </Layout>
    <BackTop />
  </Layout>
</template>

<style scoped>
.app {
  min-height: 100vh;
}
</style>

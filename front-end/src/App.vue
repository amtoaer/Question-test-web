<template>
  <div id="app">
    <mu-row class="set">
      <mu-appbar style="width: 100%;" color="red">
        <mu-button icon slot="left" @click="ifOpen=!ifOpen">
          <mu-icon value="menu"/>
        </mu-button>
        {{title}}
      </mu-appbar>
    </mu-row>
    <mu-drawer :open.sync="ifOpen" :docked=false :right=false>
      <mu-list>
        <mu-list-item button to="/" @click="ifOpen=!ifOpen">
          <mu-list-item-action><mu-icon value="home"></mu-icon></mu-list-item-action>
          <mu-list-item-title>主页</mu-list-item-title>
        </mu-list-item>
        <mu-divider></mu-divider>
        <mu-list-item button to="/马原" @click="ifOpen=!ifOpen">
          <mu-list-item-action><mu-icon value="book"></mu-icon></mu-list-item-action>
          <mu-list-item-title>马原</mu-list-item-title>
        </mu-list-item>
        <mu-list-item button to="/毛概" @click="ifOpen=!ifOpen">
          <mu-list-item-action><mu-icon value="book"></mu-icon></mu-list-item-action>
          <mu-list-item-title>毛概</mu-list-item-title>
        </mu-list-item>
        <mu-divider></mu-divider>
        <mu-list-item button to="/关于" @click="ifOpen=!ifOpen">
          <mu-list-item-action><mu-icon value="info"></mu-icon></mu-list-item-action>
          <mu-list-item-title>关于</mu-list-item-title>
        </mu-list-item>
      </mu-list>
    </mu-drawer>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
  data: function () {
    return {
      title: '主页',
      ifOpen: false
    }
  },
  created () {
    this.title = this.$route.path.replace('/', '')
    if (this.title === '') {
      this.title = '主页'
    }
  },
  watch: {
    $route (to) {
      this.title = to.path.replace('/', '')
      if (this.title === '') {
        this.title = '主页'
      }
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
</style>

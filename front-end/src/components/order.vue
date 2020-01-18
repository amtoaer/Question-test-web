<style>
  .footer {
    font-size: 0.8rem;
    line-height: 1rem;
    margin: 1rem 0 .5rem 0;
  }
</style>

<template>
  <mu-container>
    <mu-paper class="demo-paper" :z-depth="4">
      <div>
        <question v-for="i in 10" v-if="(current-1)*10+i<=total" v-bind:num="(current-1)*10+i" v-bind:type="type"></question>
      </div>
    </mu-paper>
    <mu-flex justify-content="center" style="margin-top: 20px;">
      <mu-pagination raised circle :total="total" :current.sync="current"></mu-pagination>
    </mu-flex>
    <mu-flex justify-content="center" style="margin-top: 10px;">
      <mu-footer class="footer">Designed by <a href="https://github.com/jeasonlau/">JeasonLau</a><br/>
      Powered by <a href="https://cn.vuejs.org/">Vue.js</a></mu-footer>
    </mu-flex>
  </mu-container>
</template>

<script>
  import question from "./question"
  export default {
    components: {
      "question":question
    },
    mounted () {
      if (this.type==="/马原/顺序刷题"){
        this.total=1329
      }else{
        this.total=1449
      }
    },
    data: function () {
      return {
        type: this.$route.path,
        current:1,
        total:1000,
      }
    }
  }
</script>

<template>
  <mu-container>
    <mu-paper class="demo-paper" :z-depth="4">
      <div>
        <mu-load-more @refresh="refresh" :refreshing="refreshing" :loaded-all="loadedall" :loading="loading" @load="load">
          <mu-list>
            <template v-for="i in num">
                <question :num="i" v-bind:type="type"></question>
            </template>
          </mu-list>
        </mu-load-more>
      </div>
    </mu-paper>
    <mu-flex justify-content="center" style="margin-top: 10px;">
      <mu-footer class="footer">Designed by <a href="https://github.com/jeasonlau/">JeasonLau</a><br/>
        Powered by <a href="https://cn.vuejs.org/">vue.js</a></mu-footer>
    </mu-flex>
  </mu-container>
</template>

<script>
  import question from "./question"
  export default {
    components: {
      "question":question
    },
    data: function () {
      return {
        type: this.$route.path,
        num: 10,
        refreshing: false,
        loading: false,
        loadedall: false
      }
    },
    methods: {
      refresh () {
        this.refreshing = true;
        this.$refs.container.scrollTop = 0;
        setTimeout(() => {
          this.refreshing = false;
          this.num = 10;
        }, 2000)
      },
      load () {
        if (this.num < 50) {
          this.loading = true;
          setTimeout(() => {
            this.loading = false;
            this.num += 10;
          }, 2000)
        }else{
          this.loadedall=true
        }
      }
    }
  }
</script>

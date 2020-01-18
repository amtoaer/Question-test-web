<template>
  <mu-container>
    <h2>请选择你的刷题模式：</h2>
    <mu-select label="" v-model="choice">
      <mu-option v-bind:key="option.index" v-for="option in options" :label="option" :value="option"/>
    </mu-select>
    <mu-button @click="changeLocation(choice)">确定</mu-button>
  </mu-container>
</template>

<script>
export default {
  data: function () {
    return {
      options: ['顺序刷题',  '考试模拟'],
      choice: ''
    }
  },
  methods: {
    changeLocation (path) {
      window.location.href = window.location.href + '/' + path
    }
  }
}
</script>

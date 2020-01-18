<style scoped>
  .article{
    font-size: 22px;
    margin-left: 2%;
    margin-right: 2%;
    text-indent: 2em;
  }
  .list{
    font-size: 22px;
    margin-left: 6%;
  }
</style>
<template>
  <mu-container>
    <mu-paper :z-depth="4">
      <h1 align="center">- 马原、毛概刷题工具 -</h1>
      <p class="article" align="left">
        欢迎使用！该网站提供马原、毛概的顺序刷题与考试模拟功能。其中马原题目<span style="color: #4caf50">1329</span>道
        ，毛概题目<span style="color: #4caf50">1449</span>道，可以满足大部分学生的练习需要。</p>
      <p class="article" align="left">
        如果有问题或者建议，欢迎反馈到我的邮箱<a href="mailto:b1361974998@gmail.com">b1361974998@gmail.com</a>
        或给<a href="https://github.com/jeasonlau/Question-test-web/">该项目</a>提issue。</p>
      <mu-divider></mu-divider>
      <p class="article" align="left">
        使用说明：</p>
      <ol class="list" align="left">
        <li>点击左上角的图标选择科目</li>
        <li>在跳转到的界面选择刷题方式</li>
        <li>开始您的刷题之旅！</li>
      </ol>
      <mu-flex style="margin-bottom: 20px"></mu-flex>
    </mu-paper>
    <mu-flex justify-content="center" style="margin-top: 10px;">
      <mu-footer class="footer">Designed by <a href="https://github.com/jeasonlau/">JeasonLau</a><br/>
        Powered by <a href="https://cn.vuejs.org/">vue.js</a></mu-footer>
    </mu-flex>
  </mu-container>
</template>

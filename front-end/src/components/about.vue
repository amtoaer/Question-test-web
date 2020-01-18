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
      <p class="article" align="left">
        本项目由<a href="https://github.com/jeasonlau">JeasonLau</a>开发,前后端代码已经开源，托管在github上。
      <p class="article" align="left">
      题库来源于同学们分享的word文档，本人对其进行正则匹配后以json格式存储了下来以便开发使用。
    </p>
      <p class="article" align="left">
      本项目使用的技术主要有</p>
      <ul class="list" align="left">
        <li>gin:go语言的轻量web框架</li>
        <li>vue.js:构建用户界面的渐进式框架</li>
        <li>muse-ui:基于vue2.0的Material Design UI 组件库</li>
      </ul>

      <p class="article" align="left">
        除此之外，本网站的界面对<a href="https://neumathe.xyz">扶朕起来朕还能学</a>进行了诸多参考，在此一并表示感谢。</p>
      <mu-flex style="margin-bottom: 20px"></mu-flex>
    </mu-paper>
    <mu-flex justify-content="center" style="margin-top: 10px;">
      <mu-footer class="footer">Designed by <a href="https://github.com/jeasonlau/">JeasonLau</a><br/>
        Powered by <a href="https://cn.vuejs.org/">vue.js</a></mu-footer>
    </mu-flex>
  </mu-container>
</template>

<style>
  .description{
    font-size: 16px;
  }
</style>

<template>
  <div>
    <p align="left" class="description">{{num}}. {{question['Description']}}</p>
    <div v-if="isRadio===true">
    <mu-flex style="margin-bottom: 30px; margin-top: 30px" v-if="isRadio===true" v-bind:key="choice.index" v-for="choice in question['Choice']">
      <mu-radio :value="choice[0]" color="red" :ripple="false"
                v-model="yourAnswer" :label="choice" @click="ifRight"/>
    </mu-flex>
      </div>
    <div v-if=""isRadio="==false">
    <mu-flex style="margin-bottom: 30px; margin-top: 30px" v-if="isRadio===false" v-bind:key="choice.index" v-for="choice in question['Choice']">
      <mu-checkbox :value="choice[0]" color="red" :ripple="false" :label-left="false"
                v-model="yourAnswer" :label="choice" @click="ifRight"/>
    </mu-flex>
    </div>
    <mu-divider/>
    <mu-snackbar :open.sync="isRight" color="red">
      <mu-icon left value="check_circle"/>
      题目{{num}}答案正确！
      <mu-button flat slot="action" color="#fff" @click="isRight = false">Close</mu-button>
    </mu-snackbar>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data: function () {
      return {
        question: {"Answer":"ABCD","Choice":["A.发展科技的战略基点，走中国特色自主创新道路，推动科技跨越式发展","B.调整经济结构，转变经济增长方式的中心环节，建设资源节约型.环境友好型社会，推动国民经济又好又快发展","C.国家战略，贯穿到现代化建设各个方面，激发全民族创新精神，培养高水平创新人才，形成有利于自主创新的体制机制","D.国家发展战略的核心，大力推进理论创新.制度创新.科技创新，不断巩固和发展社会主义伟大事业"],"Description":"建设创新型国家的主要要求或根本途径是以增强自主创新能力作为()"},
        isRight: false,
        isRadio: false,
        yourAnswer: [],
        url: ''
      }
    },
    props: ["num", "type"],
    mounted () {
      this.getQuestion()
    },
    watch: {
      num: function () {
        this.getQuestion()
      }
    },
    methods: {
      ifRight: function () {
        let answer
        if (this.isRadio === false) {
          answer = this.yourAnswer.sort().join("")
        } else {
          answer = this.yourAnswer
        }
        if (answer === this.question["Answer"]) {
          this.isRight = true
          let that = this
          if (this.timer)
            clearTimeout(this.timer)
          this.timer = setTimeout(function () {
            that.isRight = false
          }, 2000)
        } else {
          this.isRight = false
        }
      },
      getQuestion () {
        this.yourAnswer=[]
        if (this.type === "/马原/顺序刷题") {
          this.url = "https://jeasonlau.xyz:10000/api/mayuan/position/" + (this.num - 1)
        } else if (this.type === "/毛概/顺序刷题") {
          this.url = "https://jeasonlau.xyz:10000/api/maogai/position/" + (this.num - 1)
        } else if (this.type === "/马原/考试模拟") {
          if (this.num <= 40) {
            this.url = "https://jeasonlau.xyz:10000/api/mayuan/random/radio"
          } else {
            this.url = "https://jeasonlau.xyz:10000/api/mayuan/random/checkbox"
          }
        } else if (this.type === "/毛概/考试模拟") {
          if (this.num <= 40) {
            this.url = "https://jeasonlau.xyz:10000/api/maogai/random/radio"
          } else {
            this.url = "https://jeasonlau.xyz:10000/api/maogai/random/checkbox"
          }
        }
        axios
          .get(this.url)
          .then(response => {
            this.question = response.data
            if (this.question['Answer'].length > 1) {
              this.isRadio = false
            } else {
              this.isRadio = true
            }
          })
      }
    }
  }
</script>

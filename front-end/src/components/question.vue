<style>
  .description{
    font-size: 16px;
  }
</style>

<template>
  <div>
    <p align="left" class="description">{{num}}. {{question['Description']}}</p>
    <div v-if="isRadio===true">
    <mu-row v-if="isRadio===true" v-bind:key="choice.index" v-for="choice in question['Choice']">
      <mu-radio :value="choice[0]" color="red" :ripple="false"
                v-model="yourAnswer" :label="choice" @click="ifRight"/>
    </mu-row>
      </div>
    <div v-if=""isRadio="==false">
    <mu-row v-if="isRadio===false" v-bind:key="choice.index" v-for="choice in question['Choice']">
      <mu-checkbox :value="choice[0]" color="red" :ripple="false"
                v-model="yourAnswer" :label="choice" @click="ifRight"/>
    </mu-row>
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
        question: {},
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
          if (this.num <= 39) {
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

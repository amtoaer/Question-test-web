<template>
  <mu-container>
    <mu-paper class="demo-paper" :z-depth="4">
    <div v-for="item in questionList">
      <p align="left">{{item['count']}}.{{item['Description']}}</p>
      <mu-row v-bind:key="it.index" v-for="it in item['Choice']">
          <mu-radio :value="it"
                  v-model="item['chose']" :label="it" @click="judgeAnswer(item)"/>
      </mu-row>
      <mu-divider></mu-divider>
      <mu-snackbar :color="Success" :open.sync="item['judge']">
        <mu-icon left :value="icon"></mu-icon>
        题目{{item['count']}}答案正确！
        <mu-button flat slot="action" color="#fff" @click="item['judge'] = false">Close</mu-button>
      </mu-snackbar>
    </div>
    </mu-paper>
  </mu-container>
</template>

<script>
export default {
  data: function() {
    return {
      count: 1,
      questionList: [
        {
          "Description": "作为中国共产党和社会主义事业指导思想的马克思主义",
          "Answer": "D",
          "Choice": [
            "A.指马克思恩格斯创立的基本理论.基本观点和学说的体系",
            "B.是非正宗的马克思主义",
            "C.是实用主义的马克思主义",
            "D.是继承和发展了的马克思主义"
          ]
        },
        {
          "Description": "马克思主义三大组成部分中，列宁将其称之为“科学思想中的最大成果”的是",
          "Answer": "A",
          "Choice": [
            "A.马克思主义哲学",
            "B.马克思主义政治经济学",
            "C.马克思主义科学社会主义",
            "D.唯物史观与剩余价值学说"
          ]
        },
        {
          "Description": "提出“价值是由劳动创造”的学者是",
          "Answer": "D",
          "Choice": [
            "A.马克思",
            "B.恩格斯",
            "C.马克思.恩格斯",
            "D.亚当·斯密和大卫·李嘉图"
          ]
        }]
    }
  },
  created() {
    let cnt=1
    for (let item of this.questionList){ //设置编号
      item.count=cnt
      cnt=cnt+1
    }
  },
  methods: {
    judgeAnswer: function (item) {
      if (item["Answer"].length===1){
      if (item["Answer"]===item["chose"][0]){
        item['judge']=true
      }else{
        item['judge']=false
      }
      }
    }
  }
}
</script>

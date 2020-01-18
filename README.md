# 马原、毛概刷题工具（web version）
马原、毛概刷题工具的web版本

## 组成

本项目使用的技术主要有：

1. `gin`: go语言的轻量web框架
2. `vue.js`: 构建用户界面的渐进式框架
3. `muse-ui`: 基于`vue2.0`的Material Design UI 组件库

## 使用

> **注意：自行部署需要把`./front-end/src/components/question.vue`中请求的地址修改为自己的后端地址。**

下载该项目：

```bash
git clone https://github.com/jeasonlau/Question-test-web.git
```

+ 后端：

  ```bash
  cd ./back-end
  go run main.go
  ```

+ 前端：

  ```bash
  cd ./front-end
  npm install
  npm run build
  ```

  将`dist`文件夹下的内容放到web服务器上。

## api

| 请求                                  | 返回                                  |
| ------------------------------------- | ------------------------------------- |
| /api/(mayuan\|maogai)/position/(num)  | (mayuan\|maogai)题库中索引为num的题目 |
| /api/(mayuan\|maogai)/random/         | (mayuan\|maogai)题库中的随机一道题目  |
| /api/(mayuan\|maogai)/random/radio    | (mayuan\|maogai)题库中随机一道单选题  |
| /api/(mayuan\|maogai)/random/checkbox | (mayuan\|maogai)题库中随机一道多选题  |
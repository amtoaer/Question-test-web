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

## 已知BUG

1. 由于部分题目题干/选项过长，导致文字重叠。（目前还没有找到更合适的容纳题目的组件，就先不修复了）

   > 临时解决办法：使用大屏幕设备，如：电脑。
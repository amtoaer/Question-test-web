# 马原、毛概刷题工具（web version）
> 目前并未完成，只是先占个坑！

马原、毛概刷题工具的web版本（该项目为任务驱动型，本人对`html/css/js/go`使用并不熟练）

## 组成

本项目使用的技术主要有：

1. `gin`: go语言的轻量web框架
2. `vue.js`: 构建用户界面的渐进式框架
3. `muse-ui`: 基于`vue2.0`的Material Design UI 组件库

## 使用

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

  将`dist`文件夹下的内容部署到web服务器上
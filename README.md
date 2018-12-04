# WMuTong_vue_go

WMuTong网站重构，前端使用vue，后台使用go语言

## 运行

``` bash
# 安装依赖
npm install

# 启动项目
npm run dev

# 编译项目
npm run build
```

## 目录结构介绍 ##

    |-- build                            // webpack配置文件
    |-- config                           // 项目打包路径
    |-- src                              // 源码目录
    |   |-- components                   // 公共组件
    |   |-- view                         // 页面组件
    |       |--                          // 
    |           |--                      //
    |   |-- router                       // 路由分配
    |   |-- store                        // vuex
    |   |-- App.vue                      // 页面入口文件
    |   |-- main.js                      // 程序入口文件，加载各种公共组件
    |-- .babelrc                         // ES6语法编译配置
    |-- .editorconfig                    // 代码编写规格
    |-- .gitignore                       // 忽略的文件
    |-- index.html                       // 入口html文件
    |-- package.json                     // 项目及工具的依赖配置文件
    |-- README.md                        // 说明


## 第三方依赖 ##

```
样式： element-ui
css : sass
http: axios
```
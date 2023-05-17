# pen
<div align=center><img width="450" height="280" src="https://raw.githubusercontent.com/rr13k/pen/main/static/pen.jpg"/></div>
<p align="center">
<a href="https://coding.jd.com/cherry/cherry-core/"><img alt="Test Dependencies" src="https://badgen.net/badge/pen/pen/yellow?icon=github" /></a>
<a href="https://coding.jd.com/cherry/cherry-core/"><img alt="Test Dependencies" src="https://badgen.net/badge/web/framework/red?icon=github" /></a>
<a href="https://coding.jd.com/cherry/cherry-core/"><img alt="NPM Version" src="https://badgen.net/github/status/micromatch/micromatch/4.0.1" style="max-width:100%;"></a>
<a href="https://coding.jd.com/cherry/cherry-core/"><img alt="NPM Version" src="https://badgen.net/badge/license/MIT/blue" style="max-width:100%;"></a>
</p>

## 简介

`pen` 通过[专属模版](https://github.com/rr13k/pen-tmpl)构建一套快速起手的web框架，并提供一定的代码封装引导，帮助人们更快速的
将工作投入到核心编码当中。于其他框架不同的是，我们提倡轻量及原生并且该框架已在大量项目中取得了不错的功能验证。

## 文档
查看在线文档，请访问 [rr13k.github.io/pen-doc](https://rr13k.github.io/pen-doc/).

## 特性

一. 纯净、原生

- 使用go原生模块进行基础构建，并经过大量项目进行可靠性验证。

二. 简单、快速
- 提供一键项目生成，避免重复的项目复刻工作
- 提供友好专业的基础套件
- 自研模块提供丰富功能、示例

三. 专业
- 符合go社区的项目规范
- 提供完备的单元测试
- 轻度代码洁癖


## Start

```shell
    # 在GoPath目录下运行
    go get -u github.com/rr13k/pen  # 安装pen

    go install github.com/rr13k/pen  # 编译

    # 通常用户的GoPath为 ～/go 如果你的安装位置有更改需要手动替换        
    export PATH=$PATH:~/go/bin # 将go/bin(包含pen) 添加至环境变量

    pen new  # 通过pen新建项目
```

## 本地更新

```sh
    # 安装本地开发版本(在pen目录下执行)
    go install github.com/rr13k/pen

```

## 问题跟踪

使用我们的 `github Issues` 页面 [报告错误](https://github.com/rr13k/pen/issues) 并 [提出改进建议](https://github.com/rr13k/pen/issues)

## 执照
>Code released under the [MIT license](LICENSE).
# 基本介绍
gobrief-tools是一个gobrief的脚手架。直接init快速构建拉出一个可以运行的web服务，供您快速开发。

# 目录结构
```
➜  gobrief-tool tree .               
gobrief-tool   项目
├── README.md  
├── build.sh   编译脚本
├── cmd        命令行代码
│   ├── init.go
│   ├── root.go
│   └── version.go
├── codeGen   代码自动生成
│   └── gen.go   
├── go.mod
├── go.sum
├── main.go   入口文件
└── template  模板文件
    ├── project
    │   ├── app
    │   │   ├── api
    │   │   │   ├── dashboard.template
    │   │   │   └── sys_user.template
    │   │   ├── dao
    │   │   │   └── sys_user.template
    │   │   ├── model
    │   │   │   └── sys_user.template
    │   │   └── service
    │   ├── cmd
    │   │   ├── config.template
    │   │   ├── root.template
    │   │   ├── run.template
    │   │   └── version.template
    │   ├── go.template
    │   ├── main.template
    │   └── project
    │       ├── dbs
    │       │   ├── init.template
    │       │   └── migrate.template
    │       ├── form_validation
    │       │   ├── form_validation.template
    │       │   ├── init.template
    │       │   └── sys_user.template
    │       ├── logger
    │       │   └── logger.template
    │       ├── middleware
    │       │   ├── init.template
    │       │   ├── logger.template
    │       │   └── recover.template
    │       ├── result
    │       │   ├── result.template
    │       │   └── result_code.template
    │       └── router
    │           ├── init.template
    │           └── v1
    │               ├── dashboard.template
    │               ├── router.template
    │               ├── sys_base.template
    │               └── sys_user.template
    └── template.go  处理模板关系
```

# gobrief思想
我本人致力于不依赖框架开发，架构只是简单的构建出来结构。并不做实际内容的输入，本框架会很轻量。并不会像python里面的Django。
对于框架不同人有不同的想法，有人喜欢使用像Django一样的自带管理的快速开发框架。可我认为框架本身会限制自己很多的想象与实现。
所以，我本人也一直致力于非常轻量的框架。但我认为Gin过于轻量，Gin很适合一个团队用来构建自己的框架体系，但并不是所有团队都
有时间去一点一点研究MVC或MTV等这种架构与目录结构，甚至还要为日志和路由分组浪费时间。所以带着这样的想法gobrief就诞生了。

# 使用说明
- golang版本 >= v1.14.14
- IDE推荐：Goland
- 推荐使用：GoMod

# 使用方法
`./gobrief-tool init sysCmdb`
```
Please gobrief-tool init <projectName>

Usage:
gobrief-tool [flags]
gobrief-tool [command]

Available Commands:
help        Help about any command
init        example: gobrief-tool init <projectName>
version     Print the version number of gobrief-tool

Flags:
-h, --help   help for gobrief-tool

Use "gobrief-tool [command] --help" for more information about a command.
```
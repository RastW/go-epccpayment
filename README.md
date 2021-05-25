# go-epccpayment
This project well refactor the payment farmework of bancstone

first you need clone and switch the branch to main, execute this git commands;
```
git remote add origin https://github.com/RastW/go-epccpayment.git

git pull

git checkout main
```

## 说明
1. 此项目分为两个模块，启动文件以及main.go位于cmd对应模块下

2. service模块包含对外接收交易流程

3. job模块包含定时任务、以及相关请求（手工触发对账等处理）

4. apis文件夹下为模版、分为epcc以及esb，分别存储对外以及对内提供服务模版，均适用struct结构体模式

5. common为公共代码

6. 各模块均含有一个config.json配置，为总配置文件，json内通过数组 + map对象模式映射配置，不允许再新建配置文件

7. 各模块根目录模块名/文件夹下续自建server.go以及config.json

## 层级信息
<details>
<summary>展开查看</summary>
<pre><code>
.
├── README.md
├── apis
│   ├── epcc
│   │   └── epcc10100101.go
│   └── esb
├── cmd
│   ├── job
│   └── service
│       ├── main
│       └── main.go
├── common
│   ├── config
│   │   └── json.go
│   ├── db
│   ├── log
│   │   └── log.go
│   ├── pack
│   │   └── pack.go
│   └── server
│       └── server.go
├── go.mod
├── job
├── logfile
│   └── epcc.log
└── service
    ├── config.json
    └── http.go

15 directories, 12 files
</code></pre>
</details>

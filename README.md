# go-epccpayment
This project well refactor the payment farmework of bancstone

first you need clone and switch the branch to main, execute this git commands;
```
git remote add origin https://github.com/RastW/go-epccpayment.git

git pull

git checkout main
```

## 层级信息
<details>
<summary>展开查看</summary>
<pre><code>
.
├── README.md
├── apis 								# 接收以及发送对外服务的模版
│   ├── epcc
│   │   └── epcc10100101.go
│   └── esb
├── business 							# 业务流程
├── common 								# 通过组件
│   ├── db
│   ├── log
│   │   └── log.go
│   ├── pack
│   │   └── pack.go
│   └── server
│       └── http.go
├── config.json
├── go.mod
├── logfile 							# 日志存放
│   └── epcc.log
└── main.go

10 directories, 9 files
</code></pre>
</details>

package log

import (
	golog "log"
	"os"
)

// log对象
var log *golog.Logger

// 打印文件地址
const filePath = "/Users/wangjiawei/go/src/go-epccpayment/logfile/epcc.log"

type filelog string

// 初始化log对象
func init() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {

	}
	log = golog.New(file, "go-epcc", golog.LstdFlags)
}

func Info(data string) {
	log.Println(data)
}

package main

import (
	"getcharzp.cn/router"
)

func main() {
	r := router.Router()
	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}

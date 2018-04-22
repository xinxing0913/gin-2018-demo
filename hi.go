package main

import "github.com/gin-gonic/gin"

func main() {
    // 获取应用路由
    r := gin.Default()

    // 添加路由
    r.GET("/", func(c *gin.Context) {
        c.String(200, "hello gin框架!")
    })

    // 运行
    r.Run()
}

package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main() {
   //创建一个服务
   ginServer := gin.Default()

   //接受前端获取的参数
   //第一种 url?方式
   ginServer.GET("/userinfo", func(context *gin.Context) {
      userid := context.Query("userid")
      username := context.Query("username")
      context.JSON(http.StatusOK, gin.H{
         "userid":   userid,
         "username": username,
      })
   })
//    //第二种 路径方式
//    ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
//        userid := context.Param("userid")
//        username := context.Param("username")
//        context.JSON(http.StatusOK, gin.H{
//           "userid":   userid,
//           "username": username,
//        })
//     })
   //服务器端口
   ginServer.Run(":8082")
}
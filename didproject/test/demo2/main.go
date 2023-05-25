package main

import (
   "encoding/json"
   "github.com/gin-gonic/gin"
   "net/http"
)

func main() {
   //创建一个服务
   ginServer := gin.Default()
   //前端给后端传递JSON
   ginServer.POST("/json", func(context *gin.Context) {
      // request.body
      // []byte
      data, _ := context.GetRawData()
      var m map[string]interface{}
      //包装为JSON数据，[]byte
      _ = json.Unmarshal(data, &m)
      context.JSON(http.StatusOK, m)
   })
   //服务器端口
   ginServer.Run(":8082")
}
package main

import (
	"log"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	_ "didproject/docs"
    "github.com/swaggo/gin-swagger"

	"didproject/api/handlers"
	"didproject/api/utils"
	"didproject/api/mydb"
	"didproject/docs"
	 
)

//	@title			DIDapi
//	@version		1.0
//	@description	分布式数字身份系统
//	@host			localhost:8000
//	@BasePath		/
func main() {
	// 读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}



	//连接数据库
	dbConfig := "host=" + viper.GetString("postgresql.host") +
		" port=" + viper.GetString("postgresql.port") +
		" user=" + viper.GetString("postgresql.user") +
		" password=" + viper.GetString("postgresql.password") +
		" dbname=" + viper.GetString("postgresql.dbname") +
		" sslmode=disable"
	db, err := sql.Open("postgres", dbConfig)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	mydb.InitRedis()
	defer mydb.CloseRedis()


	// 设置 gin 模式
	gin.SetMode(viper.GetString("gin.mode"))

	// 创建 gin 实例
	r := gin.Default()
	// 添加跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 允许的域名，可以设置为 * 表示允许所有域名
		AllowMethods: []string{"*"}, // 允许的 HTTP 方法
		AllowHeaders: []string{"*"}, // 允许的请求头
	}))

	//中间件
	r.Use(utils.Logger())
	// programatically set swagger info
	docs.SwaggerInfo.Schemes = []string{"http"}	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//报表
	r.GET("/dash/count", utils.GetTableCounts(db))

	//个人信息
	r.GET("/information/select", handlers.InformationSelect(db))
	r.POST("/information/update", handlers.InformationUpdate(db))
	r.POST("/system/register", handlers.Register(db))
	r.POST("/system/login",handlers.Login(db))

	//数字身份
	r.POST("/did/create",handlers.DIDCreate)
	r.GET("/did/scan", handlers.DIDScan(db))
	
	//智能合约
	r.POST("/contract/create", handlers.ContractCreate(db))
	r.GET("/contract/select", handlers.ContractSelect(db))
	r.GET("/contract/select/sum", handlers.ContractSum(db))
	r.GET("/contract/select/:id", handlers.ContractSelectOne(db))
	r.POST("/contract/update/:id", handlers.ContractUpdate(db))
	r.DELETE("/contract/delete/:id", handlers.ContractDelete(db))
	
	//注册环境
	r.GET("/issuer/select", handlers.IssuerSelectAll(db))
	r.GET("/issuer/select/:id", handlers.IssuerSelectOne(db))

	//应用程序
	r.POST("/application/create", handlers.ApplicationCreate)
	r.POST("/application/add", handlers.ApplicationAdd(db))
	r.DELETE("/application/delete/:id", handlers.ApplicationDelete(db))
	r.POST("/application/update/:id", handlers.ApplicationUpdate(db))
	r.GET("/application/select", handlers.GetApplications)
	r.GET("/application/select/:id", handlers.ApplicationSelectOne(db))
	r.GET("/application/select/sum", handlers.ApplicationSum(db))


	// 启动 gin 服务
	port := viper.GetString("server.port")
    r.Run(":" + port)
}

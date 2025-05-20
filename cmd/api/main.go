package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "calculation-service/docs" // замените на ваш путь
)

// @title Gin Swagger API
// @version 1.0
// @description Example API for Gin + Swagger
// @host localhost:8080
// @BasePath /api/v1
//	@securityDefinitions.apikey ApiKeyAuth
//	@in header
//	@name Autorization
func main() {
    r := gin.Default()

	api := r.Group("/api/v1")
	{
        api.GET("/hello", helloHandler)
        api.GET("/hell", helloHandler2)

//         группа ручек для авторизации
//         тут нужно вставить middleware auth(yes/no)
//         тут ручки для авторизованного
	}

    logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DataBasePostgress.host"),
		Port:     viper.GetString("DataBasePostgress.port"),
		Username: viper.GetString("DataBasePostgress.username"),
		Password: viper.GetString("DataBasePostgress.password"),
		DBName:   viper.GetString("DataBasePostgress.dbname"),
		SSLMode:  viper.GetString("DataBasePostgress.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)


     // Добавьте роут для Swagger
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   r.Run(":8080")
}

// helloHandler godoc
// @Summary Получить приветствие
// @Description Возвращает простое приветственное сообщение
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "H2el222lo, 3!",
	})
}

// helloHandler godoc
// @Summary Получить приветствие
// @Description Возвращает простое приветственное сообщение
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hell [get]
func helloHandler2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "H2el11111222lo, 3!",
	})
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
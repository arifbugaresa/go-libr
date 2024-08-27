package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-libr/configs"
	"go-libr/databases/connection"
	"go-libr/databases/migration"
	"go-libr/modules/user"
	"go-libr/utils/logger"
	"go-libr/utils/rabbitmq"
	"go-libr/utils/redis"
	"go-libr/utils/swagger"
)

// @title Swagger Documentation
// @version 1.1.2
// @description This is documentation go-libr.
// @host localhost:8080
func main() {
	// initiate file configuration
	configs.Initiator()

	// initiate logger
	logger.Initiator()

	// initiate scheduler
	//scheduler.Initiator()

	// initiate redis
	redis.Initiator()

	// initiate database connection
	dbConnection, _ := connection.Initiator()
	defer dbConnection.Close()

	// initiate sql migration
	migration.Initiator(dbConnection)

	// initiate rabbitmq publisher
	//rabbitMqConn := rabbitmq.Initiator()
	//defer rabbitMqConn.Channel.Close()
	//defer rabbitMqConn.Conn.Close()

	// initiate rabbitmq consumer
	//_ = rabbitMqConn.Consume()

	// initiate router
	InitiateRouter(dbConnection, nil)
}

func InitiateRouter(dbConnection *sql.DB, rabbitMqConn *rabbitmq.RabbitMQ) {
	router := gin.Default()

	// initiate swagger docs
	swagger.Initiator(router)

	user.Initiator(router, rabbitMqConn, dbConnection)

	router.Run(viper.GetString("app.port"))
}

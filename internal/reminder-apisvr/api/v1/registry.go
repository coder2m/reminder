/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:33
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories/impl"
	server "github.com/myxy99/reminder/internal/reminder-apisvr/server"
	"github.com/myxy99/reminder/pkg/client/database"
	R "github.com/myxy99/reminder/pkg/response"
	"github.com/myxy99/reminder/pkg/validator"
)

func InitRouter(db *database.Client, validator *validator.Validator) *gin.Engine {
	service := server.NewWebService(
		impl.NewUserRepository(db.DB()),
		impl.NewRemindRepository(db.DB()))
	handler := NewUserHandler(service, validator)

	app := gin.Default()
	api := app.Group("/api/v1")
	{
		api.GET("/", func(context *gin.Context) {
			R.Ok(context, R.MSG_OK, nil)
		})

		api.POST("/login", handler.Login)
		api.PUT("/user", Auth(), handler.SetUser)
	}
	return app
}

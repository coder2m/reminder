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
	"net/http"
)

func InitRouter(db *database.Client, validator *validator.Validator) *gin.Engine {
	service := server.NewWebService(
		impl.NewUserRepository(db.DB()),
		impl.NewRemindRepository(db.DB()))
	userHandler := NewUserHandler(service, validator)
	reminderHandler := NewReminderHandler(service, validator)

	app := gin.Default()
	app.NoRoute(func(context *gin.Context) {
		R.Response(context, http.StatusNotFound, "Not Found", nil, http.StatusNotFound)
		return
	})
	api := app.Group("/api/v1")
	{
		api.POST("/login", userHandler.Login)
		api.PUT("/user", Auth(), userHandler.SetUser)

		api.GET("/reminder", Auth(), reminderHandler.GetUserReminder)
	}
	return app
}

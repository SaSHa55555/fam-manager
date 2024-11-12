package cmd

import (
	"github.com/labstack/echo/v4"
)

func Run() {
	managerApi, err := ConfigureApi()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/", managerApi.ApiTransport.HandleMainPage)

	e.GET("/register", managerApi.ApiTransport.GetCreateFamily)
	e.POST("/register", managerApi.ApiTransport.HandleCreateFamily)

	e.GET("/login", managerApi.ApiTransport.GetLogInFamily)
	e.POST("/login", managerApi.ApiTransport.HandleLogInFamily)

	e.GET("/family/:id", managerApi.ApiTransport.ShowFamilyTasks)

	e.GET("/add-task/:id", managerApi.ApiTransport.GetAddTask)
	e.POST("/add-task/:id", managerApi.ApiTransport.HandleAddTask)

	e.GET("/add-member/:id", managerApi.ApiTransport.GetAddMember)
	e.POST("/add-member/:id", managerApi.ApiTransport.HandleAddMember)

	e.POST("/update-status/:id/:name", managerApi.ApiTransport.EditTaskStatus)

	e.POST("/delete-task/:id/:name", managerApi.ApiTransport.DeleteTask)

	e.Logger.Fatal(e.Start(":5051"))
}

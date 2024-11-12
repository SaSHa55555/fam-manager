package api

import "github.com/labstack/echo/v4"

type IApiTransport interface {
	GetAddTask(context echo.Context) error
	HandleAddTask(context echo.Context) error
	EditTaskStatus(context echo.Context) error
	GetAddMember(context echo.Context) error
	HandleAddMember(context echo.Context) error
	OpenFamily(context echo.Context) error
	GetLogInFamily(context echo.Context) error
	HandleLogInFamily(context echo.Context) error
	GetCreateFamily(context echo.Context) error
	HandleCreateFamily(context echo.Context) error
	HandleMainPage(context echo.Context) error
	ShowFamilyTasks(context echo.Context) error
	DeleteTask(context echo.Context) error
}

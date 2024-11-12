package http

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/SaSHa55555/fam-manager/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type Handler struct {
	service api.IApiService
}

func NewHandler(service api.IApiService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAddTask(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	data := map[string]interface{}{
		"FamilyID": id,
	}

	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/add_task.html"))
	errTmpl := tmpl.Execute(context.Response().Writer, data)
	if errTmpl != nil {
		log.Error(err)

		return errTmpl
	}

	return nil
}

func (h *Handler) HandleAddTask(context echo.Context) error {
	var (
		task api.Task
		err  error
	)

	task.Name = context.FormValue("name")
	task.Description = context.FormValue("description")
	task.Assignee = context.FormValue("assignee")
	task.Priority, err = api.ConvertPriorityToDomain(context.FormValue("priority"))
	task.Status, err = api.ConvertStatusToDomain(context.FormValue("status"))

	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	pointsParam := context.FormValue("points")
	task.Points, err = strconv.Atoi(pointsParam)
	if err != nil {
		log.Error(err)

		return err
	}

	err = h.service.AddTask(id, task)
	if err != nil {
		if errors.Is(err, api.ErrTaskExists) || errors.Is(err, api.ErrNoSuchMember) {
			data := map[string]interface{}{
				"FamilyID": id,
			}

			tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/add_task.html"))
			errTmpl := tmpl.Execute(context.Response().Writer, data)
			if errTmpl != nil {
				log.Error(err)

				return errTmpl
			}

			if errors.Is(err, api.ErrTaskExists) {
				context.Response().Write([]byte("<script>alert('Задача с таким названием уже существует');</script>"))
			}
			if errors.Is(err, api.ErrNoSuchMember) {
				context.Response().Write([]byte("<script>alert('Указанного пользователя не существует');</script>"))
			}
		}

		log.Error(err)

		return err
	}

	return context.Redirect(http.StatusFound, "/family/"+strconv.Itoa(id))
}

func (h *Handler) ShowTask(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) EditTaskStatus(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	name := context.Param("name")
	status, err := api.ConvertStatusToDomain(context.FormValue("status"))

	err = h.service.EditTaskStatus(id, name, status)
	if err != nil {
		log.Error(err)

		return err
	}

	return context.Redirect(http.StatusFound, "/family/"+strconv.Itoa(id))
}

func (h *Handler) GetAddMember(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	data := map[string]interface{}{
		"FamilyID": id,
	}

	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/add_member.html"))
	errTmpl := tmpl.Execute(context.Response().Writer, data)
	if errTmpl != nil {
		log.Error(err)

		return errTmpl
	}

	return nil
}

func (h *Handler) HandleAddMember(context echo.Context) error {
	name := context.FormValue("name")

	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	err = h.service.AddMember(id, name)
	if err != nil {
		if errors.Is(err, api.ErrMemberExists) {
			data := map[string]interface{}{
				"FamilyID": id,
			}

			tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/add_member.html"))
			errTmpl := tmpl.Execute(context.Response().Writer, data)
			if errTmpl != nil {
				log.Error(err)

				return errTmpl
			}

			context.Response().Write([]byte("<script>alert('Имя Занято');</script>"))
		}

		log.Error(err)

		return err
	}

	return context.Redirect(http.StatusFound, "/family/"+strconv.Itoa(id))
}

func (h *Handler) OpenFamily(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) ShowFamilyTasks(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	tasks, err := h.service.ShowFamilyTasks(id)
	if err != nil {
		log.Error(err)

		return err
	}

	data := map[string]interface{}{
		"ReadyForWork": tasks[api.StatusReadyForWork],
		"InProgress":   tasks[api.StatusInProgress],
		"Done":         tasks[api.StatusDone],
		"FamilyID":     id,
	}
	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/family_tasks.html"))
	err = tmpl.Execute(context.Response().Writer, data)
	if err != nil {
		log.Error(err)

		return err
	}

	return nil
}

func (h *Handler) ShowMemberTasks(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) DeleteTask(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err)

		return err
	}

	name := context.Param("name")

	err = h.service.DeleteTask(id, name)
	if err != nil {
		log.Error(err)

		return err
	}

	return nil
}

func (h *Handler) HandleCreateFamily(context echo.Context) error {
	name := context.FormValue("name")
	pswd := context.FormValue("pswd")

	id, err := h.service.CreateFamily(name, pswd)
	if err != nil {
		if errors.Is(err, api.ErrFamilyExists) {
			tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/reg.html"))
			errTmpl := tmpl.Execute(context.Response().Writer, nil)
			if errTmpl != nil {
				log.Error(err)

				return errTmpl
			}

			context.Response().Write([]byte("<script>alert('Имя Занято');</script>"))
		}

		log.Error(err)

		return err
	}

	return context.Redirect(http.StatusFound, "/family/"+strconv.Itoa(id))
}

func (h *Handler) GetCreateFamily(context echo.Context) error {
	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/reg.html"))
	err := tmpl.Execute(context.Response().Writer, nil)
	if err != nil {
		log.Error(err)

		return err
	}

	return nil
}

func (h *Handler) GetLogInFamily(context echo.Context) error {
	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/login.html"))
	err := tmpl.Execute(context.Response().Writer, nil)
	if err != nil {
		log.Error(err)

		return err
	}

	return nil
}

func (h *Handler) HandleLogInFamily(context echo.Context) error {
	name := context.FormValue("name")
	pswd := context.FormValue("pswd")

	id, err := h.service.CheckFamily(name, pswd)
	if err != nil {
		if errors.Is(err, api.ErrWrongCreds) {
			tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/login.html"))
			errTmpl := tmpl.Execute(context.Response().Writer, nil)
			if errTmpl != nil {
				log.Error(err)

				return errTmpl
			}

			context.Response().Write([]byte("<script>alert('Неверный пароль!');</script>"))
		}

		log.Error(err)

		return err
	}

	return context.Redirect(http.StatusFound, "/family/"+strconv.Itoa(id))
}

func (h *Handler) HandleMainPage(context echo.Context) error {
	tmpl := template.Must(template.ParseFiles("internal/api/transport/tmpl/index.html"))
	err := tmpl.Execute(context.Response().Writer, nil)
	if err != nil {
		log.Error(err)

		return err
	}

	return nil
}

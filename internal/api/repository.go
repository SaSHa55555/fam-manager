package api

type IApiRepository interface {
	ShowFamilyTasks(familyID int) ([]Task, error)
	AddTask(familyID int, task Task) error
	AddMember(familyID int, name string) error
	CreateFamily(name string, pswd string) (int, error)
	EditTaskStatus(familyID int, taskName string, status Status) error
	CheckFamily(name string, pswd string) (int, error)
	DeleteTask(familyID int, taskName string) error
}

package service

import (
	"github.com/pkg/errors"

	"github.com/SaSHa55555/fam-manager/internal/api"
)

type Service struct {
	repository api.IApiRepository
}

func NewService(r api.IApiRepository) *Service {
	return &Service{repository: r}
}

func (s Service) ShowFamilyTasks(familyID int) (map[api.Status][]api.Task, error) {
	tasks, err := s.repository.ShowFamilyTasks(familyID)
	if err != nil {
		return nil, errors.Wrap(err, "show family tasks")
	}

	tasksByStatus := map[api.Status][]api.Task{
		api.StatusReadyForWork: make([]api.Task, 0),
		api.StatusInProgress:   make([]api.Task, 0),
		api.StatusDone:         make([]api.Task, 0),
	}

	for _, task := range tasks {
		tasksByStatus[task.Status] = append(tasksByStatus[task.Status], task)
	}

	return tasksByStatus, nil
}

func (s Service) AddTask(familyID int, task api.Task) error {
	err := s.repository.AddTask(familyID, task)
	if err != nil {
		return errors.Wrap(err, "add task")
	}

	return nil
}

func (s Service) AddMember(familyID int, name string) error {
	err := s.repository.AddMember(familyID, name)
	if err != nil {
		return errors.Wrap(err, "add member")
	}

	return nil
}

func (s Service) CreateFamily(name string, pswd string) (int, error) {
	id, err := s.repository.CreateFamily(name, pswd)
	if err != nil {
		return 0, errors.Wrap(err, "create family")
	}

	return id, nil
}

func (s Service) DeleteTask(familyID int, taskName string) error {
	err := s.repository.DeleteTask(familyID, taskName)
	if err != nil {
		return errors.Wrap(err, "delete task")
	}

	return nil
}

func (s Service) EditTaskStatus(familyID int, taskName string, status api.Status) error {
	err := s.repository.EditTaskStatus(familyID, taskName, status)
	if err != nil {
		return errors.Wrap(err, "edit task")
	}

	return nil
}

func (s Service) CheckFamily(name string, pswd string) (int, error) {
	id, err := s.repository.CheckFamily(name, pswd)
	if err != nil {
		return 0, errors.Wrap(err, "check family")
	}

	return id, nil
}

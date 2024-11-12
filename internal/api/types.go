package api

import "github.com/pkg/errors"

type Status string

const (
	StatusDone         Status = "done"
	StatusInProgress   Status = "in progress"
	StatusReadyForWork Status = "ready for work"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
)

type Task struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Status      Status   `json:"status" db:"status"`
	Points      int      `json:"points" db:"points"`
	Priority    Priority `json:"priority" db:"priority"`
	Assignee    string   `json:"assignee" db:"assignee"`
	Description string   `json:"description" db:"description"`
	FamilyID    int      `json:"family_id" db:"family_id"`
}

type Member struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Family struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	PSWD string `json:"pswd" db:"pswd"`
}

var (
	ErrFamilyExists = errors.New("family already exists")
	ErrMemberExists = errors.New("member already exists")
	ErrNoSuchMember = errors.New("no such member")
	ErrTaskExists   = errors.New("task already exists")
	ErrWrongCreds   = errors.New("wrong creds")
)

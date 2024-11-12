package repository

import (
	"database/sql"

	"github.com/SaSHa55555/fam-manager/internal/api"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

const duplicateKeyViolationCode = "23505"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) ShowFamilyTasks(familyID int) ([]api.Task, error) {
	query := `SELECT name, description, points, priority, assignee, status from tasks WHERE family_id = $1`

	tasks := make([]api.Task, 0, 5)
	rows, err := r.db.Query(query, familyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task api.Task
		var id int
		err = rows.Scan(&task.Name, &task.Description, &task.Points, &task.Priority, &id, &task.Status)
		if err != nil {
			return nil, err
		}

		nameQuery := `SELECT name from members WHERE id = $1`
		err = r.db.QueryRow(nameQuery, id).Scan(&task.Assignee)
		if err != nil {
			return nil, err
		}

		task.FamilyID = familyID

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r Repository) AddTask(familyID int, task api.Task) error {
	assigneeQuery := `SELECT id FROM members WHERE family_id = $1 and name = $2`

	var id int

	err := r.db.QueryRow(assigneeQuery, familyID, task.Assignee).Scan(&id)
	if err != nil {
		log.Info(err)
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrNoSuchMember
		}

		return err
	}

	query := `INSERT INTO tasks (name, description, points, priority, assignee, status, family_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = r.db.QueryRow(query, task.Name, task.Description, task.Points, task.Priority, id, task.Status, familyID).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == duplicateKeyViolationCode {
			return api.ErrTaskExists
		}

		return err
	}

	return nil
}

func (r Repository) AddMember(familyID int, name string) error {
	query := `INSERT INTO members (name, family_id) VALUES ($1, $2) RETURNING id`

	var id int

	err := r.db.QueryRow(query, name, familyID).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == duplicateKeyViolationCode {
			return api.ErrMemberExists
		}

		return err
	}

	return nil
}

func (r Repository) CreateFamily(name string, pswd string) (int, error) {
	query := `INSERT INTO families (name, pswd) VALUES ($1, $2) RETURNING id`

	var id int

	err := r.db.QueryRow(query, name, pswd).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == duplicateKeyViolationCode {
			return 0, api.ErrFamilyExists
		}

		return 0, err
	}

	return id, nil
}

func (r Repository) EditTaskStatus(familyID int, taskName string, status api.Status) error {
	query := `UPDATE tasks SET status = $1 WHERE family_id = $2 and name = $3`

	_, err := r.db.Exec(query, status, familyID, taskName)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) CheckFamily(name string, pswd string) (int, error) {
	query := `SELECT id from families where name = $1 and pswd = $2`

	var id int

	err := r.db.QueryRow(query, name, pswd).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, api.ErrWrongCreds
		}

		return 0, err
	}

	return id, nil
}

func (r Repository) DeleteTask(familyID int, taskName string) error {
	query := `DELETE FROM tasks WHERE family_id = $1 and name = $2`

	_, err := r.db.Exec(query, familyID, taskName)
	if err != nil {
		return err
	}

	return nil
}

package userrepository

import (
	"context"
	"database/sql"
	"simple-go-restapi/app/models"
	repo "simple-go-restapi/app/repository"
)

func NewSQLRepo(Conn *sql.DB) repo.UserRepo {
	return &mysqlRepo{
		Conn: Conn,
	}
}

type mysqlRepo struct {
	Conn *sql.DB
}

func (m *mysqlRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.Name,
			&data.Age,
			&data.School,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlRepo) GetByName(ctx context.Context, name string) (*models.User, error) {

	query := "Select * From Users where Name=?"

	rows, err := m.fetch(ctx, query, name)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlRepo) Create(ctx context.Context, p *models.User) (int64, error) {
	panic("implement me")
}

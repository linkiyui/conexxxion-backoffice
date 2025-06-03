package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	shared_infra "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/shared/infra"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
)

type UserPostgreRepository struct {
	pg_repo *shared_infra.PostgreRepository
}

func NewUserPostgreRepository(repo *shared_infra.PostgreRepository) domain.IUserRepository {
	return &UserPostgreRepository{
		pg_repo: repo,
	}
}

func (r *UserPostgreRepository) WithUnitOfWork(uw uow.UnitOfWork) domain.IUserRepository {
	r.pg_repo.WithUnitOfWork(uw)
	return r
}

func (r *UserPostgreRepository) GetByEmail(email string) (*domain.User, error) {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return nil, err
	}
	defer r.pg_repo.ReleaseConn()
	u := domain.User{}
	q := `SELECT id, email, username, name, last_name, password, role, create_at, update_at FROM users where email = $1`
	err = c.QueryRow(context.Background(), q, email).Scan(&u.ID, &u.Email, &u.Username, &u.Name, &u.LastName, &u.Password, &u.Role, &u.CreateAt, &u.UpdateAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		clog.Error(err.Error(), nil)
		return nil, err
	}
	return &u, nil
}

func (r *UserPostgreRepository) GetByUsername(username string) (*domain.User, error) {

	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return nil, err
	}

	defer r.pg_repo.ReleaseConn()
	u := domain.User{}
	q := `SELECT id, email, username, name, last_name, password, role, create_at, update_at FROM users where username = $1`
	err = c.QueryRow(context.Background(), q, username).Scan(&u.ID, &u.Email, &u.Username, &u.Name, &u.LastName, &u.Password, &u.Role, &u.CreateAt, &u.UpdateAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		clog.Error(err.Error(), nil)
		return nil, err
	}
	return &u, nil
}

func (r *UserPostgreRepository) Create(user *domain.User) error {

	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `INSERT INTO users(id, email,  username, name, last_name, password, role, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	_, err = c.Exec(context.Background(), q, user.ID, user.Email, user.Username, user.Name, user.LastName, user.Password, user.Role, user.CreateAt, user.UpdateAt)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

func (r *UserPostgreRepository) GetByID(id string) (*domain.User, error) {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return nil, err
	}
	defer r.pg_repo.ReleaseConn()
	u := domain.User{}
	q := `SELECT id, email, username, name, last_name, password, role, create_at, update_at FROM users where id = $1`
	err = c.QueryRow(context.Background(), q, id).Scan(&u.ID, &u.Email, &u.Username, &u.Name, &u.LastName, &u.Password, &u.Role, &u.CreateAt, &u.UpdateAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		clog.Error(err.Error(), nil)
		return nil, err
	}
	return &u, nil
}

func (r *UserPostgreRepository) UpdatePassword(id string, password string) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `UPDATE users SET password = $1 WHERE id = $2`
	_, err = c.Exec(context.Background(), q, password, id)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

func (r *UserPostgreRepository) Save(user *domain.User) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `UPDATE users SET email = $1, username = $2, name = $3, last_name = $4, update_at = $5 WHERE id = $6`
	_, err = c.Exec(context.Background(), q, user.Email, user.Username, user.Name, user.LastName, user.UpdateAt, user.ID)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

func (r *UserPostgreRepository) Delete(id string) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `DELETE FROM users WHERE id = $1`
	_, err = c.Exec(context.Background(), q, id)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

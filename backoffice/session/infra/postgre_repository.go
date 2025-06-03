package infra

import (
	"context"
	"errors"

	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"

	"github.com/jackc/pgx/v5"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/domain"
	shared_infra "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/shared/infra"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
)

type SessionPostgreRepository struct {
	pg_repo *shared_infra.PostgreRepository
}

func NewSessionPostgreRepository(pg_repo *shared_infra.PostgreRepository) domain.ISessionRepository {
	return &SessionPostgreRepository{
		pg_repo: pg_repo,
	}
}

func (r *SessionPostgreRepository) WithUnitOfWork(uw uow.UnitOfWork) domain.ISessionRepository {
	r.pg_repo.WithUnitOfWork(uw)
	return r
}

func (r *SessionPostgreRepository) GetByID(id string) (*domain.Session, error) {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return nil, err
	}
	defer r.pg_repo.ReleaseConn()
	s := domain.Session{}
	q := `SELECT id, user_id, device_info, ip, last_access, created_at FROM session where id = $1`
	err = c.QueryRow(context.Background(), q, id).Scan(&s.ID, &s.UserID, &s.DeviceInfo, &s.IP, &s.LastAccess, &s.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		clog.Error(err.Error(), nil)
		return nil, err
	}
	return &s, nil
}

func (r *SessionPostgreRepository) GetByUserID(userId string) (*domain.Session, error) {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return nil, err
	}
	defer r.pg_repo.ReleaseConn()
	s := domain.Session{}
	q := `SELECT id, user_id, device_info, ip, last_access, created_at FROM session where user_id = $1`
	err = c.QueryRow(context.Background(), q, userId).Scan(&s.ID, &s.UserID, &s.DeviceInfo, &s.IP, &s.LastAccess, &s.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		clog.Error(err.Error(), nil)
		return nil, err
	}
	return &s, nil
}

func (r *SessionPostgreRepository) Create(s domain.Session) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `INSERT INTO session(id, user_id, device_info, last_access, created_at, ip) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err = c.Exec(context.Background(), q, s.ID, s.UserID, s.DeviceInfo, s.LastAccess, s.CreatedAt, s.IP)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

func (r *SessionPostgreRepository) Delete(id string) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `DELETE FROM session WHERE id = $1`
	_, err = c.Exec(context.Background(), q, id)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

func (r *SessionPostgreRepository) DeleteUserSessions(userID string) error {
	c, err := r.pg_repo.GetConn()
	if err != nil {
		clog.Error("getting connection: "+err.Error(), nil)
		return err
	}
	defer r.pg_repo.ReleaseConn()
	q := `DELETE FROM session WHERE user_id = $1`
	_, err = c.Exec(context.Background(), q, userID)
	if err != nil {
		clog.Error(err.Error(), nil)
	}
	return err
}

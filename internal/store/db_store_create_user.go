package store

import (
	"context"
	"errors"
	"github.com/eampleev23/gophkeeper/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

// cost - сложность хэша (чем сложнее, тем дольше хэширование и сравнение, но пи этом более надежное)
const cost = 14

// InsertUser занимается непосредственно запросом вставки строки в бд.
func (d DBStore) InsertUser(ctx context.Context, userRegReq models.UserRegReq) (newUser models.User, err error) {
	newUser = models.User{}
	bytesPassword := []byte(userRegReq.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, cost)
	if err != nil {
		return newUser, err
	}

	err = d.dbConn.QueryRow( // нужен скан
		`INSERT INTO
    users (login, password)
	VALUES($1, $2)
	RETURNING
	    id, login, password`,
		userRegReq.Login,
		hashedPassword).Scan(
		&newUser.ID,
		&newUser.Login,
		&newUser.Password)
	// Проверяем, что ошибка сигнализирует о потенциальном нарушении целостности данных
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		err = ErrConflict
	}
	return newUser, err
}

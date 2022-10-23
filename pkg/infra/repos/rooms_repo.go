package repos

import (
	"database/sql"

	"cinema-booking.ro/be/pkg/domain/model"
	"github.com/jmoiron/sqlx"
)

type RoomsRepo struct {
	db *sqlx.DB
}

var (
	getAllRoomsSQL = `SELECT id, name, seats FROM rooms`
)

func NewRoomsRepo(db *sqlx.DB) *RoomsRepo {
	return &RoomsRepo{db: db}
}

func (r *RoomsRepo) GetAll() ([]*model.Room, error) {

	result := []*model.Room{}
	if err := r.db.Select(&result, getAllRoomsSQL); err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}
	return result, nil

}

package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang_api/Ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByUserId(id int) (*types.User, error) {
	rows, err := s.db.Query("Select * from tbl_golang_user where id=?", id)
	if err != nil {
		log.Fatal(err)
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.USERNAME,
		&user.USER_ID,
		&user.INSERTED_ON,
		&user.UPDATED_ON,
	)

	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *Store) CreateUser(user types.User) error {
	return nil
}

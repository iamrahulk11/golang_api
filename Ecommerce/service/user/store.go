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

func (s *Store) GetAllUser() (*[]types.User, error) {
	if err := s.db.Ping(); err != nil {
		return nil, err
	}

	query := `select id, user_id, username,inserted_on,updated_on from tbl_golang_user`

	rows, err := s.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	var au []types.User

	for rows.Next() {
		db_user := types.User{}
		err = rows.Scan(&db_user.ID, &db_user.USERNAME, &db_user.USER_ID, &db_user.INSERTED_ON, &db_user.UPDATED_ON)
		if err != nil {
			// handle this error
			panic(err)
		}
		au = append(au, db_user)
	}

	// if u.ID == 0 {
	// 	return nil, fmt.Errorf("user not found")
	// }

	return &au, nil
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

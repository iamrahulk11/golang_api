package v1

import "time"

type userModal struct {
	USER_ID  string `json:"user_id"`
	USERNAME string `json:"username"`
	USER_IMG string `json:"user_img"`
}

type db_userModal struct {
	ID          int        `json:"id"`
	USER_ID     string     `json:"user_id"`
	USERNAME    string     `json:"username"`
	INSERTED_ON time.Time  `json:"inserted_on"`
	UPDATED_ON  *time.Time `json:"updated_on"`
}

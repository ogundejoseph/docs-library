package model

type User struct {
	ID         uint64 `json:"id"`
	FName      string `json:"first_name"`
	LName      string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func CreateUser(user *User) error {
	statement := `INSERT INTO users(first_name, last_name, username, email, phone, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.Exec(statement, user.FName, user.LName, user.Username, user.Email, user.Phone, user.Password, user.Created_at, user.Updated_at)
	return err
}

func GetUser(id string) (User, error) {
	var user User

	statement := `SELECT * FROM users WHERE id=$1;`

	rows, err := db.Query(statement, id)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FName, &user.LName, &user.Username, &user.Email, &user.Phone, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			return User{}, err
		}
	}

	return user, nil

}

func CheckEmail(email string, user *User) bool {
	statement := `SELECT id, name, email FROM users WHERE email=$1 LIMIT 1;`

	rows, err := db.Query(statement, email)
	if err != nil {
		return false
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FName, &user.LName, &user.Username, &user.Email, &user.Phone, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			return false
		}
	}
	return true
}

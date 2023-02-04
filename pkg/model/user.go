package model

type User struct {
	Id       string `db:"student_id"`
	Username string `db:"student_username"`
	Password string `db:"student_password"`
}

// Get Method
func (m *User) GetId() string {
	return m.Id
}

func (m *User) GetUsername() string {
	return m.Username
}

func (m *User) GetPassword() string {
	return m.Password
}

// set Method
func (m *User) SetId(id string) {
	m.Id = id
}

func (m *User) SetUsername(username string) {
	m.Username = username
}

func (m *User) SetPassword(password string) {
	m.Password = password
}

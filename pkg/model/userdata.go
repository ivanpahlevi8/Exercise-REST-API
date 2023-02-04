package model

type UserData struct {
	Id       string `json:"student_id"`
	Username string `json:"student_username"`
	Password string `json:"student_password"`
}

// Get Method
func (m *UserData) GetId() string {
	return m.Id
}

func (m *UserData) GetUsername() string {
	return m.Username
}

func (m *UserData) GetPassword() string {
	return m.Password
}

// set Method
func (m *UserData) SetId(id string) {
	m.Id = id
}

func (m *UserData) SetUsername(username string) {
	m.Username = username
}

func (m *UserData) SetPassword(password string) {
	m.Password = password
}

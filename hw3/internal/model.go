package model

type User struct {
	userId   int
	name     string
	login    string
	password string
}

type Notify struct {
	notifyId   int
	message    string
	dataCreate string
	tag        string
	ownerId    int
}

func NewUser(userId int, name string, login string, password string) *User {
	return &User{
		userId:   userId,
		name:     name,
		login:    login,
		password: password,
	}
}

func (u *User) GetUserId() int {
	return u.userId
}

func (u *User) GetUsername() string {
	return u.name
}

func (u *User) GetUserLogin() string {
	return u.login
}

func (u *User) GetUserPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

package main

type User struct {
	name    string
	id      int
	session Session
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetSession(sess Session) {
	u.session = sess
}

func CreateUser(userName string, session Session) *User {
	var newUser *User
	newUser.SetName(userName)
	newUser.SetSession(session)
	return newUser
}

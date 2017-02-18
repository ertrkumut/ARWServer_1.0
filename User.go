package main

type User struct {
	name    string
	pass    string
	id      int
	session Session
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetPassword(pass string) {
	u.pass = pass
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetSession(sess Session) {
	u.session = sess
}

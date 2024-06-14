package models

type IUser interface {
	SetAccessLevel(al int)
	SetUID(id string)
	SetFname(fname string)
	SetLname(lname string)
	SetEmail(email string)
	SetPhone(phone string)
	SetPassword(phone string)
	GetUID() string
	GetFname() string
	GetLname() string
	GetEmail() string
	GetPhone() string
	GetPassword() string
}

type User struct {
	AccessLevel int    `json:"access"`
	UID         string `json:"id"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	Uname       string `json:"uname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

// GetID
func (u *User) GetID() string {
	return u.UID
}

// GetEmail implements IUser.
func (u *User) GetEmail() string {
	return u.Email
}

// GetFname implements IUser.
func (u *User) GetFname() string {
	return u.Fname
}

// GetLname implements IUser.
func (u *User) GetLname() string {
	return u.Lname
}

// GetUname implements IUser.
func (u *User) GetUname() string {
	return u.Uname
}

// GetPhone implements IUser.
func (u *User) GetPhone() string {
	return u.Phone
}

// GetPassword implements IUser.
func (u *User) GetPassword() string {
	return u.Password
}

// SetEmail implements IUser.
func (u *User) SetEmail(email string) {
	u.Email = email
}

// SetFname implements IUser.
func (u *User) SetFname(fname string) {
	u.Fname = fname
}

// SetID implements IUser.
func (u *User) SetID(id string) {
	u.UID = id
}

// SetLname implements IUser.
func (u *User) SetLname(lname string) {
	u.Lname = lname
}

// SetUname implements IUser.
func (u *User) SetUname(uname string) {
	u.Uname = uname
}

// SetPhone implements IUser.
func (u *User) SetPhone(phone string) {
	u.Phone = phone
}

// SetPassword implements IUser.
func (u *User) SetPassword(password string) {
	u.Password = password
}

// SetAccessLevel implements IUser.
func (u *User) SetAccessLevel(al int) {
	u.AccessLevel = al
}

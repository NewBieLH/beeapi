package models

var (
	UserList map[string]*User
)

func (u *User) TableName() string {
	return GetUserTableName()
}

type User struct {
	Id          int
	UserName    string
	IsSuper 	int8
	UserPwd 	string
	Status 		int
	Mobile		string
	Email		string
	Avatar 		string
}


func AddUser(u User) string {
	return ""
}

func GetUser(uid string) (u *User, err error) {
	return u,err
}

//func GetAllUsers() map[string]*User {
//}
func GetAllUsers() string {
	return ""
}
func UpdateUser(uid string, uu *User) (a *User, err error) {
	return a,err
}

func Login(username, password string) bool {
	return true
}

func DeleteUser(uid string) {

}

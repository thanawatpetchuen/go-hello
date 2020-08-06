package user

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

func (u User) GetPassword() string {
	return u.Name + u.ID
}

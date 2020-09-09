package models

type Role int

const (
	ADMIN Role = iota
	MEMBER
)

func (r Role) String() string {
	return [...]string{"ADMIN", "MEMBER"}[r]
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

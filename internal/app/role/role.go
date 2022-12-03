package role

type Role int

const (
	User    Role = iota // 0
	Manager             // 1
	Admin               // 2
)

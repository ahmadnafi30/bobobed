package entity

type User struct {
	ID        int64  `json:"id"`         // ID sebagai primary key, biasanya tipe data int64
	FirstName string `json:"first_name"` // FirstName
	LastName  string `json:"last_name"`  // LastName
	Email     string `json:"email"`      // Email
	Password  string `json:"password"`   // Password
}

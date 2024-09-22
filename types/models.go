package types

type UserModel struct {
	ID       int64  `db:"id"`       // Assuming ID is auto-incremented
	Username string `db:"username"` // Add fields according to your schema
	Email    string `db:"email"`    // Add fields according to your schema
}

package datastruct

type SessionUserClient struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	LoggedIn bool   `json:"logged_in"`
}

package auth

type SignUpInput struct {
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"required"`
	Password          string `json:"password" validate:"required,min=8"`
	PasswordConfirm   string `json:"password_confirm" validate:"required,min=8"`
	EncryptPassword   string `json:"encrypt_password"`
	Photo             string `json:"photo"`
	NotificationToken string `json:"notification_token"`
}

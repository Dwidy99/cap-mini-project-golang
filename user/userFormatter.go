package user

type UserFormatter struct {
	ID         int    `json:"user_id"`
	Name       string `json:"name"`
	Occupasion string `json:"occupasion"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupasion: user.Occupasion,
		Token:      token,
	}

	return formatter
}
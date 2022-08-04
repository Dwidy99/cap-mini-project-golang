package user

type RegisterUserInput struct {
	// validasi
	Name       string `json:"name" binding:"required"`        // Validation
	Occupasion string `json:"occupasion" binding:"required"`  // Validation
	Email      string `json:"email" binding:"required,email"` // Validation
	Password   string `json:"password" binding:"required"`    // Validation
}
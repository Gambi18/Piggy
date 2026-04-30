package models

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type CreateTransactionPayload struct {
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
}

type Transaction struct {
	ID        *int32 `json:"id"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"createdAt"`
}

const (
	TypeSaving     = "saving"
	TypeWithdrawal = "withdrawal"
)

type SignUpPayload struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

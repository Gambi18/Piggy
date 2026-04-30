package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Balance  int32  `json:"balance"`
}

type CreateTransactionPayload struct {
	UserID    string `json:"userId"`
	Amount    int32  `json:"amount"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
}

type Transaction struct {
	ID        *int32 `json:"id"`
	Amount    int32  `json:"amount"`
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

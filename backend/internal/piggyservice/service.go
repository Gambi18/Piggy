package piggyservice

import (
	"context"
	"fmt"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) *Service {
	return &Service{repo: repo}
}

// define service methods here

func (s *Service) CreateTransaction(ctx context.Context, payload models.CreateTransactionPayload) (*models.Transaction, error) {
	userID, err := stringToUUID(payload.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id format: %w", err)
	}

	// Start a database transaction
	q, tx, err := s.repo.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// 1. Get current user to check balance
	user, err := q.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Calculate new balance
	currentBalance := numericToInt32(user.Balance)
	var newBalance int32
	if payload.Type == models.TypeSaving {
		newBalance = currentBalance + payload.Amount
	} else if payload.Type == models.TypeWithdrawal {
		if currentBalance < payload.Amount {
			return nil, fmt.Errorf("insufficient balance: current %d, requested %d", currentBalance, payload.Amount)
		}
		newBalance = currentBalance - payload.Amount
	} else {
		return nil, fmt.Errorf("invalid transaction type: %s", payload.Type)
	}

	// 3. Update user balance
	_, err = q.UpdateUserBalance(ctx, sqlc.UpdateUserBalanceParams{
		ID:      userID,
		Balance: int32ToNumeric(newBalance),
	})
	if err != nil {
		fmt.Printf("[ERROR] Failed to update balance for user %s: %v\n", userID, err)
		return nil, fmt.Errorf("failed to update balance: %w", err)
	}

	// 4. Create the transaction record
	transaction, err := q.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		UserID: userID,
		Amount: int32ToNumeric(payload.Amount),
		Type:   &payload.Type,
		Reason: &payload.Reason,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	// 5. Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		fmt.Printf("[ERROR] Failed to commit transaction for user %s: %v\n", userID, err)
		return nil, err
	}

	fmt.Printf("[SUCCESS] Transaction record created and committed for user %s (%s of %d)\n", userID, payload.Type, payload.Amount)

	return sqlCToAppTransaction(transaction), nil
}

func (s *Service) GetTransactions(ctx context.Context, userID string) (*[]models.Transaction, error) {
	uid, err := stringToUUID(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id format: %w", err)
	}
	txns, err := s.repo.Do().GetTransactions(ctx, uid)
	if err != nil {
		return nil, err
	}
	transactions := []models.Transaction{}
	for _, v := range txns {
		transactions = append(transactions, *sqlCToAppTransaction(v))
	}
	return &transactions, nil
}

func (s *Service) SignUp(ctx context.Context, payload models.SignUpPayload) (*models.User, error) {
	user, err := s.repo.Do().CreateUser(ctx, sqlc.CreateUserParams{
		Username: payload.Username,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Balance:  pgtype.Numeric{Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return sqlCToAppUser(user), nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.repo.Do().GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return sqlCToAppUser(user), nil
}

func (s *Service) GetBalance(ctx context.Context, userID string) (*models.UserBalanceResponse, error) {
	uid, err := stringToUUID(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id format: %w", err)
	}

	totals, err := s.repo.Do().GetTransactionTotals(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction totals: %w", err)
	}

	res := &models.UserBalanceResponse{}

	for _, t := range totals {
		val := int32(t.Total)
		if t.Type != nil && *t.Type == models.TypeSaving {
			res.TotalSavings = val
		} else if t.Type != nil && *t.Type == models.TypeWithdrawal {
			res.TotalWithdrawals = val
		}
	}
	
	// Calculate balance as the difference
	res.Balance = res.TotalSavings - res.TotalWithdrawals

	return res, nil
}

func (s *Service) Login(ctx context.Context, payload models.SignInPayload) (*models.User, error) {
	user, err := s.repo.Do().GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	// Simple password check (In a real app, use bcrypt)
	if user.Password != payload.Password {
		return nil, fmt.Errorf("invalid password")
	}

	return sqlCToAppUser(user), nil
}

func stringToUUID(s string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	err := uid.Scan(s)
	return uid, err
}

func int32ToNumeric(n int32) pgtype.Numeric {
	var num pgtype.Numeric
	num.Scan(fmt.Sprintf("%d", n))
	return num
}

func numericToInt32(n pgtype.Numeric) int32 {
	var val int32
	n.Scan(&val)
	return val
}

func sqlCToAppUser(u sqlc.User) *models.User {
	id := fmt.Sprintf("%x-%x-%x-%x-%x", u.ID.Bytes[0:4], u.ID.Bytes[4:6], u.ID.Bytes[6:8], u.ID.Bytes[8:10], u.ID.Bytes[10:16])
	return &models.User{
		ID:       id,
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
		Balance:  numericToInt32(u.Balance),
	}
}

func sqlCToAppTransaction(t sqlc.Transaction) *models.Transaction {
	id := int32(t.ID)
	res := &models.Transaction{
		Amount:    numericToInt32(t.Amount),
		ID:        &id,
		CreatedAt: t.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
	if t.Reason != nil {
		res.Reason = *t.Reason
	}
	if t.Type != nil {
		res.Type = *t.Type
	}
	return res
}

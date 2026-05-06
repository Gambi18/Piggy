package piggyservice

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"

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

	const errFailedToUpdateBalance = "failed to update balance: %w"

	// 2. Calculate new balance using database balance directly
	// For withdrawal, use calculated balance from GetBalance API to ensure consistency
	// For savings, update balance normally
	if payload.Type == models.TypeWithdrawal {
		// Get current calculated balance from transaction history
		balanceResp, err := s.GetBalance(ctx, userID.String())
		if err != nil {
			return nil, fmt.Errorf("failed to get balance: %w", err)
		}
		currentBalance := balanceResp.Balance

		if currentBalance < payload.Amount {
			return nil, fmt.Errorf("insufficient balance: current %d, requested %d", currentBalance, payload.Amount)
		}
		newBalance := currentBalance - payload.Amount

		// Update user balance with calculated value
		_, err = q.UpdateUserBalance(ctx, sqlc.UpdateUserBalanceParams{
			ID:      userID,
			Balance: int32ToNumeric(newBalance),
		})
		if err != nil {
			return nil, fmt.Errorf(errFailedToUpdateBalance, err)
		}
	} else if payload.Type == models.TypeSaving {
		// For savings, get current balance from database and update
		currentBalance := numericToInt32(user.Balance)
		newBalance := currentBalance + payload.Amount

		_, err = q.UpdateUserBalance(ctx, sqlc.UpdateUserBalanceParams{
			ID:      userID,
			Balance: int32ToNumeric(newBalance),
		})
		if err != nil {
			return nil, fmt.Errorf(errFailedToUpdateBalance, err)
		}
	} else {
		return nil, fmt.Errorf("invalid transaction type: %s", payload.Type)
	}

	// 3. Create transaction record
	transaction, err := q.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		UserID: userID,
		Amount: fmt.Sprintf("%d", payload.Amount),
		Type:   &payload.Type,
		Reason: &payload.Reason,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	// 4. Commit the transaction
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

func (s *Service) GetTransactionsByType(ctx context.Context, userID string, transactionType string) (*[]models.Transaction, error) {
	uid, err := stringToUUID(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id format: %w", err)
	}
	txns, err := s.repo.Do().GetTransactionsByType(ctx, uid, transactionType)
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
		// Convert int64 to pgtype.Numeric first, then to int32
		numericTotal := pgtype.Numeric{Int: big.NewInt(t.Total), Valid: true}
		val := numericToInt32(numericTotal)
		if t.Type != nil && *t.Type == models.TypeSaving {
			res.TotalSavings = val
		} else if t.Type != nil && *t.Type == models.TypeWithdrawal {
			res.TotalWithdrawals = val
		}
	}

	// Calculate balance as difference
	fmt.Printf("DEBUG: GetBalance - TotalSavings: %d, TotalWithdrawals: %d\n", res.TotalSavings, res.TotalWithdrawals)
	res.Balance = res.TotalSavings - res.TotalWithdrawals
	fmt.Printf("DEBUG: GetBalance - Final Balance: %d\n", res.Balance)
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
	return pgtype.Numeric{Int: big.NewInt(int64(n)), Valid: true}
}

func numericToInt32(n pgtype.Numeric) int32 {
	if !n.Valid {
		return 0
	}

	fmt.Printf("DEBUG: numericToInt32 called with Valid=%v, Int=%v, Exp=%v\n", n.Valid, n.Int, n.Exp)

	// For PostgreSQL numeric types, handle the internal representation
	if n.Int != nil {
		// If there's no exponent, return the integer value directly
		if n.Exp == 0 {
			result := int32(n.Int.Int64())
			fmt.Printf("DEBUG: returning %d (no exponent)\n", result)
			return result
		}

		// If there's an exponent, we need to adjust the value
		if n.Exp < 0 {
			// Negative exponent means decimal places
			// For amounts like 85000.00, we want to return 85000
			divisor := int64(math.Pow10(int(-n.Exp)))
			result := int32(n.Int.Int64() / divisor)
			fmt.Printf("DEBUG: returning %d (with negative exponent %d)\n", result, n.Exp)
			return result
		} else {
			// Positive exponent means multiply by power of 10
			multiplier := int64(math.Pow10(int(n.Exp)))
			result := int32(n.Int.Int64() * multiplier)
			fmt.Printf("DEBUG: returning %d (with positive exponent %d)\n", result, n.Exp)
			return result
		}
	}

	// Try to get the value as a string (fallback)
	if val, err := n.Value(); err == nil {
		fmt.Printf("DEBUG: n.Value() returned: %v (type: %T)\n", val, reflect.TypeOf(val))
		if str, ok := val.(string); ok {
			if parsed, err := strconv.ParseFloat(str, 64); err == nil {
				result := int32(parsed)
				fmt.Printf("DEBUG: parsed string value %d\n", result)
				return result
			}
		}
	}

	fmt.Printf("DEBUG: returning 0 (no valid or other methods worked)\n")
	return 0
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

	// Parse amount from string field
	var amount int32
	if t.Amount != "" {
		fmt.Printf("DEBUG: t.Amount = '%s' (length: %d)\n", t.Amount, len(t.Amount))
		if parsed, err := strconv.ParseFloat(t.Amount, 64); err == nil {
			amount = int32(parsed)
			fmt.Printf("DEBUG: parsed amount = %d\n", amount)
		} else {
			fmt.Printf("DEBUG: failed to parse amount: %v\n", err)
		}
	} else {
		fmt.Printf("DEBUG: t.Amount is empty\n")
	}

	res := &models.Transaction{
		Amount:    amount,
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

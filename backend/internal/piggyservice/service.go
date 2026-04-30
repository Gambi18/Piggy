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
	userID := stringToUUID(payload.UserID)
	// create the transaction
	transaction, err := s.repo.Do().CreateTransaction(ctx, sqlc.CreateTransactionParams{
		UserID: userID,
		Amount: int32ToNumeric(payload.Amount),
		Type:   &payload.Type,
		Reason: &payload.Reason,
	})
	if err != nil {
		return nil, err
	}
	return sqlCToAppTransaction(transaction), nil
}

func (s *Service) GetTransactions(ctx context.Context, userID string) (*[]models.Transaction, error) {
	uid := stringToUUID(userID)
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

func stringToUUID(s string) pgtype.UUID {
	var uid pgtype.UUID
	uid.Scan(s)
	return uid
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
	return &models.Transaction{
		Amount:    numericToInt32(t.Amount),
		Reason:    *t.Reason,
		Type:      *t.Type,
		ID:        &id,
		CreatedAt: t.CreatedAt.Time.String(),
	}
}

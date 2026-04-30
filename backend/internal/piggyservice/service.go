package piggyservice

import (
	"context"
	"fmt"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/models"
)

type Service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) *Service {
	return &Service{repo: repo}
}

// define service methods here

func (s *Service) CreateTransaction(ctx context.Context, payload models.CreateTransactionPayload) (*models.Transaction, error) {
	// create the transaction
	transaction, err := s.repo.Do().CreateTransaction(ctx, sqlc.CreateTransactionParams{
		Amount: payload.Amount,
		Type:   &payload.Type,
		Reason: &payload.Reason,
	})
	if err != nil {
		return nil, err
	}
	return sqlCToAppTransaction(transaction), nil
}

func (s *Service) GetTransactions(ctx context.Context) (*[]models.Transaction, error) {
	txns, err := s.repo.Do().GetTransactions(ctx)
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
	})
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

func sqlCToAppUser(u sqlc.User) *models.User {
	return &models.User{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
	}
}

func sqlCToAppTransaction(t sqlc.Transaction) *models.Transaction {
	return &models.Transaction{
		Amount:    t.Amount,
		Reason:    *t.Reason,
		Type:      *t.Type,
		ID:        &t.ID,
		CreatedAt: t.CreatedAt.Time.String(),
	}
}

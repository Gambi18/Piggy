-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, reason, type) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransactions :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (username, name, email, password, balance) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;
-- name: CreateTransaction :one
INSERT INTO transactions (amount, reason, type) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransactions :many
SELECT * FROM transactions;

-- name: CreateUser :one
INSERT INTO users (username, name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;
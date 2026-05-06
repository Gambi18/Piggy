-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, reason, type) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransactions :many
SELECT id, user_id, amount::text as amount, reason, created_at, type FROM transactions WHERE user_id = $1;

-- name: GetTransactionsByType :many
SELECT id, user_id, amount::text as amount, reason, created_at, type FROM transactions WHERE user_id = $1 AND type = $2;

-- name: CreateUser :one
INSERT INTO users (username, name, email, password, balance) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserBalance :one
UPDATE users SET balance = $2 WHERE id = $1 RETURNING *;

-- name: GetTransactionTotals :many
SELECT type, SUM(amount) as total FROM transactions WHERE user_id = $1 GROUP BY type;
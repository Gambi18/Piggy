CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  balance NUMERIC(18,2) NOT NULL DEFAULT 0.00
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  amount NUMERIC(18,2) NOT NULL,
  reason TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  type VARCHAR(25)
);
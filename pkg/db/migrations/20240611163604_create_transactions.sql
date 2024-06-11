-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  name VARCHAR(100),
  amount NUMERIC(12, 2),
  type VARCHAR(20),
  category_id int NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id),
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  account_id int NOT NULL,
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd

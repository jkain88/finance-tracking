-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS budgets (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  label VARCHAR(100),
  amount NUMERIC(12, 2),
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  category_id int NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE budgets;
-- +goose StatementEnd

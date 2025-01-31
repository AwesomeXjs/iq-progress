-- +goose Up
SELECT 'up SQL query';
CREATE TABLE users (
                    id SERIAL PRIMARY KEY,
                    username VARCHAR(50) UNIQUE NOT NULL,
                    balance DECIMAL(10, 2) NOT NULL
);

-- Создание таблицы транзакций
CREATE TABLE transactions (
                    id SERIAL PRIMARY KEY,
                    from_user_id INT NOT NULL,
                    to_user_id INT NOT NULL,
                    amount DECIMAL(10, 2) NOT NULL,
                    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    FOREIGN KEY (from_user_id) REFERENCES users(id),
                    FOREIGN KEY (to_user_id) REFERENCES users(id)
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE transactions;
DROP TABLE users;

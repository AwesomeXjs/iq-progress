-- +goose Up

SELECT 'up SQL query';
CREATE TABLE users (
                    id SERIAL PRIMARY KEY,
                    username VARCHAR(50) UNIQUE NOT NULL,
                    balance INT NOT NULL
);

-- Создание таблицы транзакций
CREATE TABLE transactions (
                    id SERIAL PRIMARY KEY,
                    from_user_id INT NOT NULL,
                    to_user_id INT NOT NULL,
                    amount INT NOT NULL,
                    type VARCHAR(20) NOT NULL,
                    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    FOREIGN KEY (from_user_id) REFERENCES users(id),
                    FOREIGN KEY (to_user_id) REFERENCES users(id)
);



-- +goose Down
DROP TABLE users;
DROP TABLE transactions;

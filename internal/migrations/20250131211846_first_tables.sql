-- +goose Up
INSERT INTO users (
    username, balance)
VALUES ('Vasya', 0),
       ('Dima', 0),
       ('Petya', 0),
       ('Vladimir', 0),
       ('Kolya', 0),
       ('Sasha', 0);

-- +goose Down


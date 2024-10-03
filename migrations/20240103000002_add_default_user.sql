-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, age, social, created_at, updated_at)
VALUES (1, 'Willy', 21, 'some-link.com', '2024-01-03 14:26:00+04', '2024-01-03 14:28:00+04');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM users
WHERE id = 1;
-- +goose StatementEnd

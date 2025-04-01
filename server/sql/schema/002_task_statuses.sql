-- +goose Up
CREATE TABLE task_statuses (
    status VARCHAR(100) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE task_statuses;
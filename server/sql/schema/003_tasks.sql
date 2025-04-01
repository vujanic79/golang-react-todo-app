-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(100) NOT NULL REFERENCES task_statuses(status),
    complete_deadline TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tasks;
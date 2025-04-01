-- name: CreateTaskStatus :one
INSERT INTO task_statuses (status)
VALUES ($1)
RETURNING *;

-- name: GetTaskStatuses :many
SELECT * FROM task_statuses;

-- name: GetTaskStatusByStatus :one
SELECT * FROM task_statuses ts
WHERE ts.status = $1;
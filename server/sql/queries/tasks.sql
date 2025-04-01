-- name: CreateTask :one
INSERT INTO tasks(id, created_at, updated_at, title, description, status, complete_deadline, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks t
WHERE t.id = $1;

-- name: UpdateTask :one
UPDATE tasks SET title = $2, description = $3, status = $4, complete_deadline = $5
WHERE id = $1
RETURNING *;

-- name: GetTasksByUserId :many
 SELECT * FROM tasks t
 WHERE t.user_id = $1;
-- name: CreateUser :one
INSERT INTO users (
        username,
        hashed_password,
        name,
        email,
        is_active,
        tenant_id
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY name
LIMIT $1
OFFSET $2;

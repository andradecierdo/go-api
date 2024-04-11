-- name: CreateUser :one
INSERT INTO "users"(
    firstName,
    lastName,
    email
) VALUES (
     $1, $2, $3
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM "users"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "users"
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "users"
SET
    firstName = $2,
    lastName = $3,
    email = $4
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

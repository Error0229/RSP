-- name: CreateUser :one

INSERT INTO
    users (name, email, password, role)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: QueryUserByEmail :one

SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: QueryAllUser :many

SELECT * FROM users;
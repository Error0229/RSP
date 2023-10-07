-- name: AddProduct :one

INSERT INTO products (name, price) VALUES ($1, $2) RETURNING *;

-- name: QueryProductByName :one

SELECT * FROM products WHERE name = $1 LIMIT 1;

-- name: QueryAllProduct :many

SELECT * FROM products;
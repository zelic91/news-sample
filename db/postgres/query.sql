
-- Query

-- users
-- name: FindUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM users
WHERE username LIKE '%' || $1 || '%' LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email LIKE '%' || $1 || '%' LIMIT 1;

-- name: FindAllUsers :many
SELECT * FROM users
ORDER BY created_at;

-- name: FindUsersByUsername :many
SELECT * FROM users
WHERE username LIKE '%' || $1 || '%'
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: FindUsersByEmail :many
SELECT * FROM users
WHERE email LIKE '%' || $1 || '%'
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    username,
    email,
    password_hashed,
    password_salt,
    status,
    created_by,
    updated_by,
    deleted_by,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, 
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    NOW(),
    NOW(),
    NOW()
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
    first_name = $2,
    last_name = $3,
    username = $4,
    email = $5,
    password_hashed = $6,
    password_salt = $7,
    status = $8,
    updated_by = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Devices
-- name: FindDeviceById :one
SELECT * FROM devices
WHERE id = $1 LIMIT 1;

-- name: FindAllDevices :many
SELECT * FROM devices
ORDER BY created_at;

-- name: CreateDevice :one
INSERT INTO devices (
    device_token,
    platform,
    status,
    created_by,
    updated_by,
    deleted_by,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW(),
    NOW(),
    NOW()
)
RETURNING *;

-- name: UpdateDevice :one
UPDATE devices
SET
    device_token = $2,
    platform = $3,
    status = $4,
    updated_by = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE id = $1;

-- Categories
-- name: FindCategoryById :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: FindAllCategories :many
SELECT * FROM categories
ORDER BY created_at;

-- name: CreateCategory :one
INSERT INTO categories (
    description,
    title,
    created_by,
    updated_by,
    deleted_by,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW(),
    NOW()
)
RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
SET
    description = $2,
    title = $3,
    updated_by = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- Objects
-- name: FindObjectById :one
SELECT * FROM objects
WHERE id = $1 LIMIT 1;

-- name: FindAllObjects :many
SELECT * FROM objects
ORDER BY created_at;

-- name: CreateObject :one
INSERT INTO objects (
    category_id,
    content,
    title,
    created_by,
    updated_by,
    deleted_by,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW(),
    NOW(),
    NOW()
)
RETURNING *;

-- name: UpdateObject :one
UPDATE objects
SET
    category_id = $2,
    content = $3,
    title = $4,
    updated_by = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteObject :exec
DELETE FROM objects
WHERE id = $1;
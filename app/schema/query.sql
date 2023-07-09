-- name: FindUserByID :one
SELECT `id`, `name`, `biography`, `email`, `created_at`, `updated_at`
FROM
    `user`
WHERE
    `id` = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT `id`, `name`, `biography`, `email`, `created_at`, `updated_at`
FROM
    `user`
ORDER BY
    `created_at` ASC;

-- name: CreateUser :execresult
INSERT INTO
    `user` ( `id`, `name`, `biography`, `email`, `created_at`, `updated_at`)
VALUES
    (?, ?, ?, ?, ?, ?);

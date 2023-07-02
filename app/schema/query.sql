-- name: GetUser :one
SELECT `id`, `name`, `created_at`, `updated_at`
FROM
    `user`
WHERE
    `id` = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT `id`, `name`, `created_at`, `updated_at`
FROM
    `user`
ORDER BY
    `created_at` ASC;

-- name: CreateUser :execresult
INSERT INTO
    `user` (`id`, `name`, `created_at`, `updated_at`)
VALUES
    (?, ?, ?, ?);

-- name: UpdateUserName :execresult
UPDATE
    `user`
SET
    `name` = ?,
    `updated_at` = ?
WHERE
    `id` = ?;

-- name: DeleteUser :execrows
DELETE FROM
    `user`
WHERE
    `id` = ?;
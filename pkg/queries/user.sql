-- name: ListUsers :many
-- get all users ordered by the username
SELECT * FROM webapp.users
ORDER BY user_name;

-- name: GetUser :one
-- get users of a particular user_id
SELECT * FROM webapp.users
WHERE user_id = $1;

-- name: DeleteUsers :exec
-- delete a particular user
DELETE FROM webapp.users
WHERE user_id = $1;

-- name: CreateUsers :one
-- insert new user
INSERT INTO webapp.users (User_Name, Pass_Word_Hash, name)
VALUES ($1,
        $2,
        $3) RETURNING *;


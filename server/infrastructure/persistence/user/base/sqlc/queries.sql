-- name: Insert :one
INSERT INTO "user".t_user("name", age, document)
VALUES (@name, @age, @document)
RETURNING "id", "created_at", "name", "age", "document";

-- name: FindByID :one
SELECT * FROM "user".t_user
WHERE id = @id;

-- name: Update :exec
UPDATE "user".t_user SET "name" = @name , age = @age, document = @document 
WHERE id = @id;

-- name: Delete :exec
UPDATE "user".t_user SET deleted_at = NOW()
WHERE id = @id;

-- name: FindAll :many
SELECT * FROM "user".t_user
ORDER BY "name" ASC;
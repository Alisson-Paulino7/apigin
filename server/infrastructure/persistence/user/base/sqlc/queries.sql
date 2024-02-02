-- name: Insert :one

INSERT INTO "user".t_user("name", age, document)
VALUES (@name, @age, @document)
RETURNING "id", "created_at", "name", "age", "document";

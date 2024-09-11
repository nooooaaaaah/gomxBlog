-- queries.sql

-- name: CreatePost :execresult
INSERT INTO posts (id, title, content, description, link, published_on)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateCategory :execresult
INSERT INTO categories (id, name)
VALUES (?, ?);

-- name: AddPostCategory :exec
INSERT INTO post_categories (post_id, category_id)
VALUES (?, ?);

-- name: GetPost :one
SELECT * FROM posts WHERE id = ?;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

-- name: ListPosts :many
SELECT * FROM posts ORDER BY published_on DESC LIMIT ?;

-- name: ListCategories :many
SELECT * FROM categories ORDER BY name;

-- name: GetPostCategories :many
SELECT c.* FROM categories c
JOIN post_categories pc ON c.id = pc.category_id
WHERE pc.post_id = ?;

-- name: GetCategoryPosts :many
SELECT p.* FROM posts p
JOIN post_categories pc ON p.id = pc.post_id
WHERE pc.category_id = ?;

-- name: CheckTablesExist :one
SELECT
    CASE
        WHEN (
            SELECT COUNT(*)
            FROM sqlite_master
            WHERE type = 'table'
            AND name IN ('posts', 'categories', 'post_categories')
        ) = 3 THEN 1
        ELSE 0
    END AS all_tables_exist;

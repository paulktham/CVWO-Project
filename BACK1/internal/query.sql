-- name: LoadCategory :many
SELECT DISTINCT "Category" FROM "Post"
ORDER BY "TimePosted";

-- name: LoadComments :many
SELECT "Comment"."Content", "Comment"."TimePosted", "Comment"."Username", "Comment"."Points"
FROM "Comment"
JOIN "Post" ON "Comment"."PostID" = "Post"."PostID";

-- name: CreatePost :one
INSERT INTO "Post" (
  "Title", "Content", "TimePosted", "ID", "Category"
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateComment :one
INSERT INTO "Comment" (
  "PostID", "Content", "TimePosted", "ID", "Username", "Points"
) VALUES (
  $1, $2, $3, $4, $5, 0
)
RETURNING *;

-- name: CreateUser :one
INSERT INTO "User" (
  "ID", "Username", "PasswordKey"
) VALUES (
  $1, $2, $3
) RETURNING "Username";

-- name: UpVote :one
UPDATE "Comment"
SET "Points" = "Points" + 1 
WHERE "CommentID" = $1
RETURNING "Points";

-- name: DownVote :one
UPDATE "Comment"
SET "Points" = "Points" - 1 
WHERE "CommentID" = $1
RETURNING "Points";

-- name: LoginUser :one
SELECT "Username" FROM "User"
WHERE "Username" = $1 AND "PasswordKey" = $2;
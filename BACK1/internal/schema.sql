CREATE TABLE "User" (
    "ID" SERIAL PRIMARY KEY,
    "Username" VARCHAR(20),
    "PasswordKey" VARCHAR(80),
    "Email" VARCHAR(50),
    "FullName" VARCHAR(50)
);

CREATE TABLE "Post" (
    "PostID" SERIAL PRIMARY KEY,
    "Title" VARCHAR(20),
    "Content" VARCHAR(1000),
    "TimePosted" TIMESTAMP,
    "ID" INT,
    "Category" VARCHAR(20), -- Assuming the ENUM values can be represented as VARCHAR
    CONSTRAINT "FK_Post_User" FOREIGN KEY ("ID") REFERENCES "User"("ID") ON DELETE CASCADE
);

CREATE TABLE "CommentLike" (
    "UserID" CHAR(15) REFERENCES "User"("ID") ON DELETE CASCADE, 
    "PostID" CHAR(15) REFERENCES "Post"("ID") ON DELETE CASCADE,
    UNIQUE ("UserID", "PostID")
);

CREATE TABLE "Comment" (
    "CommentID" SERIAL PRIMARY KEY,
    "PostID" INT,
    "Content" VARCHAR(1000),
    "TimePosted" TIMESTAMP,
    "ID" INT,
    "Username" VARCHAR(20),
    "Points" INT,
    CONSTRAINT "FK_Comment_Post" FOREIGN KEY ("PostID") REFERENCES "Post"("PostID")
);
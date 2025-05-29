CREATE TABLE "users" (
    "id" VARCHAR(256) PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "avatarurl" VARCHAR(1000),
    "bio" VARCHAR(100) NOT NULL,
    "birthday" DATE,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
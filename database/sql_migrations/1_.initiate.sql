-- +migrate Up
-- +migrate StatementBegin

DROP TABLE IF EXISTS "books";
DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "users";

CREATE TABLE "books" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "image_url" varchar,
  "release_year" integer,
  "price" integer,
  "total_page" integer,
  "thickness" varchar,
  "category_id" integer,
  "created_at" timestamp,
  "created_by" varchar,
  "modified_at" timestamp,
  "modified_by" varchar
);

CREATE TABLE "categories" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp,
  "created_by" varchar,
  "modified_at" timestamp,
  "modified_by" varchar
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "created_at" timestamp,
  "created_by" varchar,
  "modified_at" timestamp,
  "modified_by" varchar
);

ALTER TABLE "books" ADD FOREIGN KEY ("id") REFERENCES "categories" ("id");

-- Insert a user into the users table
INSERT INTO users (id, username, password, created_at, created_by, modified_at, modified_by)
VALUES (1, 'testuser', 'password123', NOW(), 'migration', NOW(), 'migration');

-- +migrate StatementEnd

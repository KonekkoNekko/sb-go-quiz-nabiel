-- +migrate Up
-- +migrate StatementBegin

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

-- +migrate StatementEnd
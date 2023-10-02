CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar UNIQUE NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "category_id" uuid NOT NULL,
  "title" varchar UNIQUE NOT NULL,
  "subtitle" varchar NOT NULL,
  "content" VARCHAR NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "publicated" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts-tags" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "post_id" uuid NOT NULL,
  "tag_id" uuid NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "categories" ("id");

CREATE INDEX ON "posts" ("id");

CREATE INDEX ON "posts-tags" ("id");

CREATE INDEX ON "posts-tags" ("post_id");

CREATE INDEX ON "posts-tags" ("tag_id");

ALTER TABLE "posts-tags" ADD UNIQUE ("post_id", "tag_id");

ALTER TABLE "posts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "posts-tags" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "posts-tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");
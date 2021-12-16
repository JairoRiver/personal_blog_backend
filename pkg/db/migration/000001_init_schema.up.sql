CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
                         "id" uuid PRIMARY KEY,
                         "role_id" uuid NOT NULL,
                         "username" varchar UNIQUE NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
                         "id" uuid PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
                         "id" uuid PRIMARY KEY,
                         "user_id" uuid NOT NULL,
                         "title" varchar NOT NULL,
                         "subtitle" varchar NOT NULL,
                         "content" text NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "roles" ("id");

CREATE INDEX ON "posts" ("id");

CREATE INDEX ON "posts" ("title");

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "username" varchar UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "is_active" boolean NOT NULL DEFAULT FALSE,
    "tenant_id" varchar
);

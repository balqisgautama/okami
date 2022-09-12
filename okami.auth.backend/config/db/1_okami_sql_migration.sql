-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS user_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "users" (
    "user_id" BIGINT DEFAULT nextval('user_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "username" VARCHAR(256) NOT NULL UNIQUE,
    "email" VARCHAR(256) NOT NULL UNIQUE,
    "password" VARCHAR(256) NOT NULL,
    "client_id" VARCHAR(256) NOT NULL UNIQUE,
    "status" SMALLINT NOT NULL,
    "locale" VARCHAR(20) NOT NULL DEFAULT 'id-ID',
    "additional_info" TEXT,
    "last_token" TIMESTAMP WITHOUT TIME ZONE,
    "created_by" BIGINT,
    "created_client" VARCHAR(256),
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_by" BIGINT,
    "updated_client" VARCHAR(256),
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_by" BIGINT,
    "deleted_client" VARCHAR(256),
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);

-- password hash dari aasdafds12A#
INSERT INTO users (user_id, username, email, password, client_id, created_by, created_client, status)
VALUES (1,'okami.admin.auth', 'info.okami.project@gmail.com',
        '$2a$14$Qt5mLU963E8la4L8lC5LQe47Qv92.j28Uz71oK/uBIUyOZmMNFi9q',
        'f23992cfdef34e1f9fcdd441d27d5cb7', 1, 'f23992cfdef34e1f9fcdd441d27d5cb7', 2);

-- +migrate StatementEnd
-- +migrate Down

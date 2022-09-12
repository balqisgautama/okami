-- +migrate Up
-- +migrate StatementBegin
--
-- ACTIVATION USERS TABLE
CREATE SEQUENCE IF NOT EXISTS user_activation_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "user_activation" (
    "activation_id" BIGINT DEFAULT nextval('user_activation_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "counter_regenerate" SMALLINT NOT NULL DEFAULT 0,
    "code" VARCHAR(256) NOT NULL,
    "expired_at" TIMESTAMP WITHOUT TIME ZONE,
    "status" INTEGER NOT NULL DEFAULT 1,
    "user_id" BIGINT NOT NULL UNIQUE,
    "email_to" VARCHAR(256) NOT NULL,
    "email_link_validate" TEXT NOT NULL,
    "email_link_resend" TEXT NOT NULL,
    CONSTRAINT "FK__users" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

-- +migrate StatementEnd
-- +migrate Down

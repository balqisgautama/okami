-- +migrate Up
-- +migrate StatementBegin

-- LOG ACTIVITY TABLE
CREATE SEQUENCE IF NOT EXISTS log_pkce_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "log_pkce" (
    "log_pkce_id" BIGINT DEFAULT nextval('log_pkce_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "step1" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "step2" TIMESTAMP WITHOUT TIME ZONE,
    "step3" TIMESTAMP WITHOUT TIME ZONE,
    "user_client_id" VARCHAR(256),
    "secret_code" VARCHAR(256),
    "code_challenger" VARCHAR(256) NOT NULL UNIQUE
);

-- +migrate StatementEnd
-- +migrate Down

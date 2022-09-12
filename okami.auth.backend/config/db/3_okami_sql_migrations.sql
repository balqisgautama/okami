-- +migrate Up
-- +migrate StatementBegin

-- RESOURCES TABLE
CREATE SEQUENCE IF NOT EXISTS resources_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "resources" (
    "resource_id" BIGINT DEFAULT nextval('resources_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "client_id" VARCHAR(256) NOT NULL,
    "surname" VARCHAR(256) NOT NULL,
    "nickname" VARCHAR(50) NOT NULL UNIQUE,
    "access_to" TEXT NULL DEFAULT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);

-- hanya resource Authentication Gate yang dapat akses langsung ke Authentication Server
INSERT INTO resources (resource_id, client_id, surname, nickname, access_to)
VALUES (1, 'f23992cfdef34e1f9fcdd441d27d5cb7', 'Authentication Gate', 'gate', 'auth');

-- +migrate StatementEnd
-- +migrate Down

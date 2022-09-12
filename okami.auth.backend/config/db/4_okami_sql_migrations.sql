-- +migrate Up
-- +migrate StatementBegin

-- LOG ACTIVITY TABLE
CREATE SEQUENCE IF NOT EXISTS log_activities_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "log_activities" (
    "activity_id" BIGINT DEFAULT nextval('log_activities_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "activity_time" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "activity_detail" TEXT NOT NULL,
    "resource_client_id" VARCHAR(256) NOT NULL
);

-- LOG AUDIT SYSTEM
CREATE SEQUENCE IF NOT EXISTS log_audit_system_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE "log_audit_system" (
    "audit_id" BIGINT DEFAULT nextval('log_audit_system_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "audit_time" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "audit_detail" TEXT NOT NULL,
    "resource_client_id" VARCHAR(256) NOT NULL,
    "data_old" TEXT NULL,
    "data_new" TEXT NOT NULL,
    "action" SMALLINT NOT NULL
);

-- +migrate StatementEnd
-- +migrate Down

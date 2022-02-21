BEGIN;

SET ROLE 'postgres';

CREATE TABLE schema_migrations_history (
    id SERIAL PRIMARY KEY NOT NULL,
    migration_version BIGINT NOT NULL,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION TRACK_APPLIED_MIGRATION()
RETURNS TRIGGER AS $$
DECLARE _current_version integer;
BEGIN
    SELECT COALESCE(MAX(migration_version),0)
    FROM schema_migrations_history
    INTO _current_version;
    IF new.dirty = 'f' AND new.version > _current_version THEN
        INSERT INTO schema_migrations_history(migration_version) VALUES (new.version);
    END IF;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- TRIGGER
CREATE TRIGGER track_applied_migrations AFTER INSERT ON schema_migrations
FOR EACH ROW EXECUTE PROCEDURE TRACK_APPLIED_MIGRATION();

COMMIT;

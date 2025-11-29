-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE if NOT EXISTS ecodes (
    id varchar(10) NOT NULL,
	description varchar(100) NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	CONSTRAINT pk_ecode PRIMARY KEY (id)
);

INSERT INTO ecodes (id, description) VALUES ('dip1', 'connection refuse') ON CONFLICT DO NOTHING;
INSERT INTO ecodes (id, description) VALUES ('dip2', 'connection timeout') ON CONFLICT DO NOTHING;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
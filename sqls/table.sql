CREATE EXTENSION btree_gist;

CREATE TABLE IF NOT EXISTS codes
(
    id serial,
    value character(8) COLLATE pg_catalog."default" NOT NULL,
    owner varchar not null,
    payload varchar not null,
    valid_between tstzrange NOT NULL,
    CONSTRAINT otc_pkey PRIMARY KEY (id),
    CONSTRAINT unique_code_between_time_window EXCLUDE USING gist (valid_between WITH &&, value WITH =)
);

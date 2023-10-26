CREATE TABLE password_reset_keys(
    id              VARCHAR(40)     PRIMARY KEY,
    user_id         VARCHAR(40)     NOT NULL REFERENCES users(id),
    reset_key       VARCHAR(255)    NOT NULL,
    used            BOOLEAN,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT clock_timestamp(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT clock_timestamp()
);
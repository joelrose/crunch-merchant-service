CREATE TABLE whitelist (
    identifier text UNIQUE PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
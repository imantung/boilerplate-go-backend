CREATE TABLE IF NOT EXISTS oauth2_tokens (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    expires_at TIMESTAMPTZ NOT NULL,
    code TEXT NOT NULL,
    access TEXT NOT NULL,
    refresh TEXT NOT NULL,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_oauth2_tokens_expires_at ON oauth2_tokens (expires_at);

CREATE INDEX IF NOT EXISTS idx_oauth2_tokens_code ON oauth2_tokens (code);

CREATE INDEX IF NOT EXISTS idx_oauth2_tokens_access ON oauth2_tokens (access);

CREATE INDEX IF NOT EXISTS idx_oauth2_tokens_refresh ON oauth2_tokens (refresh);
CREATE TABLE IF NOT EXISTS oauth2_clients (
    id SERIAL PRIMARY KEY,
    client_id TEXT UNIQUE NOT NULL,
    user_id INT NOT NULL,
    "secret" TEXT NOT NULL,
    "domain" TEXT NOT NULL,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_oauth2_clients_client_id ON oauth2_clients (client_id);
CREATE INDEX idx_oauth2_clients_user_id ON oauth2_clients (user_id);
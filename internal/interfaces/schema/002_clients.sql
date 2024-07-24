-- +goose  Up
CREATE TABLE
    clients (
        id UUID PRIMARY KEY,
        client_id VARCHAR(255) UNIQUE NOT NULL,
        client_secret VARCHAR(255) NOT NULL,
        redirect_uris TEXT NOT NULL,
        grant_types TEXT NOT NULL,
        scope TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose Down
DROP TABLE clients;
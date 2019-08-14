CREATE KEYSPACE IF NOT EXISTS auth WITH replication = {'class': 'SimpleStrategy', 'factor': 1};

CREATE TABLE IF NOT EXISTS api_token (
    id text,
    account_id text,
    api_token text,
    created_at timestamp,
    is_enabled boolean,
    deleted_at timestamp,
    PRIMARY KEY(token),
);
CREATE INDEX IF NOT EXISTS token_by_id on api_token (id);
CREATE INDEX IF NOT EXISTS token_by_acc_id on api_token (account_id);
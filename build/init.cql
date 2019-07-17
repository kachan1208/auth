CREATE KEYSPACE IF NOT EXISTS auth WITH replication = {'class': 'SimpleStrategy', 'factor': 1};

CREATE TABLE token (
    id text,
    account_id text,
    token text,
    created_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY(token),
)
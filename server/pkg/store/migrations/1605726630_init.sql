CREATE TABLE "patch" (
    id TEXT PRIMARY KEY,
    fingerprint TEXT UNIQUE,
    name TEXT,
    description TEXT,
    data JSONB,
    created_at TIMESTAMP
);
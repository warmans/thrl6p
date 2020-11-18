CREATE TABLE "patch" (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE,
    description TEXT,
    data JSONB,
    created_at TIMESTAMP
);
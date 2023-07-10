CREATE TABLE chat (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL
);
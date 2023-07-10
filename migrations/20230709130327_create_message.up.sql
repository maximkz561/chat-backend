CREATE TABLE message (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    chat_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    author_type smallint NOT NULL
);
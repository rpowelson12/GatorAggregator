-- +goose Up

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
    
);

CREATE TABLE IF NOT EXISTS sessions(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id int REFERENCES users (id) NOT NULL,
    rotated boolean DEFAULT FALSE,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL 
);

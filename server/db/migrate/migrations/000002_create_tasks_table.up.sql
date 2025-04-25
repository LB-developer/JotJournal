CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    monthly boolean DEFAULT FALSE NOT NULL,
    weekly boolean DEFAULT FALSE NOT NULL,
    daily boolean DEFAULT FALSE NOT NULL,
    deadline timestamp NOT NULL,
    description varchar(300),
    is_completed boolean DEFAULT FALSE NOT NULL,
    user_id int REFERENCES users (id)
);

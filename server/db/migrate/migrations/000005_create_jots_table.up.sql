CREATE TABLE IF NOT EXISTS jots (
    id SERIAL PRIMARY KEY,
    habit varchar(15) NOT NULL,
    date date NOT NULL,
    is_completed boolean DEFAULT FALSE NOT NULL,
    user_id int REFERENCES users (id),
    UNIQUE (user_id, habit, date)
);

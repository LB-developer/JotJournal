CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    email varchar (300) UNIQUE NOT NULL,
    password varchar(130) NOT NULL,
    created_at date DEFAULT CURRENT_DATE NOT NULL 
);

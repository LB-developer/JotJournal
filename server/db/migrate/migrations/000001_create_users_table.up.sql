CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    firstName varchar(50) UNIQUE NOT NULL,
    lastName varchar(50) UNIQUE NOT NULL,
    password varchar(130) NOT NULL,
    email varchar (300) UNIQUE NOT NULL
);

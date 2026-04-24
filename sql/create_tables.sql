CREATE TABLE IF NOT EXISTS genres (
    id         SERIAL PRIMARY KEY,
    genre      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS movies (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    runtime      INT NOT NULL,
    mpaa_rating  VARCHAR(10) NOT NULL,
    description  TEXT NOT NULL,
    image        VARCHAR(255),
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS movies_genres (
    id       SERIAL PRIMARY KEY,
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    genre_id INT NOT NULL REFERENCES genres(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Seed genres
INSERT INTO genres (genre, created_at, updated_at) VALUES
    ('Comedy',    '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Sci-Fi',    '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Horror',    '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Romance',   '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Action',    '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Thriller',  '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Drama',     '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Mystery',   '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Crime',     '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Animation', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Adventure', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Fantasy',   '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
    ('Superhero', '2022-09-23 00:00:00', '2022-09-23 00:00:00');

-- Seed movies
INSERT INTO movies (title, release_date, runtime, mpaa_rating, description) VALUES
    ('Highlander',             '1986-03-07', 116, 'R',     'He fought his first battle on the Scottish Highlands'),
    ('Raiders of the Lost Ark','1981-06-12', 115, 'PG-13', 'Archaeology professor Indiana Jones ventures to stop the Ark of the Covenant from falling into Nazi hands'),
    ('The Godfather',          '1972-03-24', 175, '18A',   'The aging patriarch of an organized crime dynasty transfers control to his reluctant son');

-- Seed movies_genres (movie_id : genre_id)
-- Highlander (1): Action(5), Fantasy(12)
-- Raiders of the Lost Ark (2): Action(5), Adventure(11)
-- The Godfather (3): Crime(9), Drama(7)
INSERT INTO movies_genres (movie_id, genre_id) VALUES
    (1, 5),
    (1, 12),
    (2, 5),
    (2, 11),
    (3, 9),
    (3, 7);

-- Seed users (password is bcrypt hashed)
INSERT INTO users (first_name, last_name, email, password) VALUES
    ('Admin', 'User', 'admin@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOlEHpyh2D');

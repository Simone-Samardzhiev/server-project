CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL
);

CREATE TABLE events
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    date        TIMESTAMP    NOT NULL,
    address     VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL
);

CREATE TABLE users_events
(
    id      SERIAL PRIMARY KEY,
    userId  INT NOT NULL REFERENCES users (id),
    eventId INT NOT NULL REFERENCES events (id)
);




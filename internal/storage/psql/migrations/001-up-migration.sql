CREATE TABLE shoguns
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE daimyo
(
    username        VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname        VARCHAR(255) NOT NULL UNIQUE,
    cards_balance   FLOAT        NOT NULL DEFAULT 0, -- Остаток на картах под конец смены
    shogun_username VARCHAR(255) NOT NULL REFERENCES shoguns (username)
);

CREATE TABLE cards
(
    id              INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(50)  NOT NULL,
    last_digits     INTEGER      NOT NULL UNIQUE,
    daily_limit     INTEGER      NOT NULL DEFAULT 2000000,
    daimyo_username VARCHAR(255) NOT NULL REFERENCES daimyo (username)
);

CREATE TABLE samurai
(
    username           VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname           VARCHAR(255) NOT NULL UNIQUE,
    daimyo_username    VARCHAR(255) NOT NULL REFERENCES daimyo (username),
    turnover_per_shift FLOAT        NOT NULL DEFAULT 0
);

CREATE TABLE administrators
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

-- CREATE TABLE daimyo_cards
-- (
--     id              INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
--     daymio_username VARCHAR(255) NOT NULL REFERENCES daimyo (username),
--     card_id         INTEGER      NOT NULL REFERENCES cards (id)
-- );

CREATE TABLE cash_managers
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE
--     replenishment_request_id INTEGER      NOT NULL REFERENCES replenishment_requests (id)
);

CREATE TABLE replenishment_request_status
(
    id   INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(50) NOT NULL UNIQUE
);

CREATE TABLE replenishment_requests
(
    id                    INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    cash_manager_username VARCHAR(255) NOT NULL REFERENCES cash_managers (username),
    daimyo_username       VARCHAR(255) NOT NULL REFERENCES daimyo (username),
    card_id               INTEGER      NOT NULL REFERENCES cards (id),
    status_id             INTEGER      NOT NULL REFERENCES replenishment_request_status (id)
);

CREATE TABLE roles
(
    id   INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE user_roles
(
    id       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
--     user_id  INTEGER      NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    role_id  INTEGER      NOT NULL REFERENCES roles (id)
);

INSERT INTO roles(name)
VALUES ('Администратор'),
       ('Сёгун'),
       ('Даймё'),
       ('Самурай'),
       ('Инкассатор')
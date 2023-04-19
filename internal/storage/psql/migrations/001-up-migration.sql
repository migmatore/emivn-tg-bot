CREATE TABLE cards
(
    id                       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    issuing_bank_information VARCHAR(255) NOT NULL,
    daily_limit              INTEGER      NOT NULL DEFAULT 2000000,
    username                 VARCHAR(255) NOT NULL UNIQUE,
    nickname                 VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE shoguns
(
    id       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) NOT NULL UNIQUE,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE daimyo
(
    id            INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username      VARCHAR(255) NOT NULL UNIQUE,
    nickname      VARCHAR(255) NOT NULL UNIQUE,
    cards_balance FLOAT        NOT NULL, -- Остаток на картах под конец смены
    shogun_id     INTEGER      NOT NULL REFERENCES shoguns (id)
);

CREATE TABLE samurai
(
    id                 INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username           VARCHAR(255) NOT NULL UNIQUE,
    nickname           VARCHAR(255) NOT NULL UNIQUE,
    daimyo_id          INTEGER      NOT NULL REFERENCES daimyo (id),
    turnover_per_shift FLOAT        NOT NULL
);

CREATE TABLE administrators
(
    id       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) NOT NULL UNIQUE,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE daimyo_cards
(
    id        INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    daymio_id INTEGER NOT NULL REFERENCES daimyo (id),
    card_id   INTEGER NOT NULL REFERENCES cards (id)
);

CREATE TABLE replenishment_requests
(
    id        INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    daimyo_id INTEGER NOT NULL REFERENCES daimyo (id),
    card_id   integer NOT NULL REFERENCES cards (id)
);

CREATE TABLE cash_managers
(
    id                       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username                 VARCHAR(255) NOT NULL UNIQUE,
    nickname                 VARCHAR(255) NOT NULL UNIQUE,
    replenishment_request_id INTEGER      NOT NULL REFERENCES replenishment_requests (id)
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

INSERT INTO roles(name) VALUES ('Администратор'), ('Сёгун'), ('Даймё'), ('Самурай'), ('Инкассатор')
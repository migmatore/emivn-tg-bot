CREATE TABLE cards
(
    id                       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    issuing_bank_information VARCHAR(255) NOT NULL,
    daily_limit              INTEGER      NOT NULL DEFAULT 2000000,
    username                 VARCHAR(255) NOT NULL,
    nickname                 VARCHAR(255) NOT NULL
);

CREATE TABLE shoguns
(
    id       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL
);

CREATE TABLE daimyo
(
    id            INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username      VARCHAR(255) NOT NULL,
    nickname      VARCHAR(255) NOT NULL,
    cards_balance INTEGER      NOT NULL, -- Остаток на картах под конец смены
    shogun_id     INTEGER      NOT NULL REFERENCES shoguns (id)
);

CREATE TABLE samurai
(
    id                 INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username           VARCHAR(255) NOT NULL,
    nickname           VARCHAR(255) NOT NULL,
    daimyo_id          INTEGER      NOT NULL REFERENCES daimyo (id),
    turnover_per_shift INTEGER      NOT NULL
);

CREATE TABLE administrators
(
    id       INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL
);

CREATE TABLE daimyo_cards
(
    id      INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    daymio_id INTEGER NOT NULL REFERENCES daimyo(id),
    card_id INTEGER NOT NULL REFERENCES cards (id)
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
    username                 VARCHAR(255) NOT NULL,
    nickname                 VARCHAR(255) NOT NULL,
    replenishment_request_id INTEGER      NOT NULL REFERENCES replenishment_requests (id)
);
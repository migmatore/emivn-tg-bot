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

CREATE TABLE bank_types
(
    id   INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(50) NOT NULL UNIQUE
);

CREATE TABLE cards
(
    id              INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(50)  NOT NULL UNIQUE,
    daimyo_username VARCHAR(255) NOT NULL REFERENCES daimyo (username),
    last_digits     INTEGER      NOT NULL UNIQUE,
    daily_limit     INTEGER      NOT NULL DEFAULT 2000000,
    balance         FLOAT        NOT NULL DEFAULT 0,
    bank_type_id    INTEGER      NOT NULL REFERENCES bank_types (id)
);

CREATE TABLE samurai
(
    username           VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname           VARCHAR(255) NOT NULL UNIQUE,
    daimyo_username    VARCHAR(255) NOT NULL REFERENCES daimyo (username),
    turnover_per_shift FLOAT        NOT NULL DEFAULT 0,
    chat_id            BIGINT NULL
);

CREATE TABLE samurai_turnovers
(
    id               INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    samurai_username VARCHAR(255) NOT NULL REFERENCES samurai (username),
    start_date       DATE         NOT NULL DEFAULT now(),
    initial_amount   FLOAT        NOT NULL DEFAULT 0,
    final_amount     FLOAT        NOT NULL DEFAULT 0,
    turnover         FLOAT        NOT NULL DEFAULT 0,
    bank_type_id     INTEGER      NOT NULL REFERENCES bank_types (id)
);

CREATE TABLE administrators
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE cash_managers
(
    username        VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname        VARCHAR(255) NOT NULL UNIQUE,
    shogun_username VARCHAR(255) NOT NULL UNIQUE REFERENCES shoguns (username),
    chat_id         BIGINT NULL
--     replenishment_request_id INTEGER      NOT NULL REFERENCES replenishment_requests (id)
);

CREATE TABLE main_operators
(
    username        VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname        VARCHAR(255) NOT NULL UNIQUE,
    shogun_username VARCHAR(255) NOT NULL UNIQUE REFERENCES shoguns (username)
);

CREATE TABLE controllers
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE controller_turnovers
(
    id                  INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    controller_username VARCHAR(255) NOT NULL REFERENCES controllers (username),
    samurai_username    VARCHAR(255) NOT NULL REFERENCES samurai (username),
    start_date          DATE         NOT NULL DEFAULT now(),
    initial_amount      FLOAT        NOT NULL DEFAULT 0,
    final_amount        FLOAT        NOT NULL DEFAULT 0,
    turnover            FLOAT        NOT NULL DEFAULT 0,
    bank_type_id        INTEGER      NOT NULL REFERENCES bank_types (id)
);

CREATE TABLE replenishment_request_status_groups
(
    id   INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(50) NOT NULL UNIQUE
);

CREATE TABLE replenishment_requests
(
    id                    INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    cash_manager_username VARCHAR(255) NOT NULL REFERENCES cash_managers (username),
    owner_username        VARCHAR(255) NOT NULL,
    card_id               INTEGER      NOT NULL REFERENCES cards (id),
    amount                DECIMAL      NOT NULL DEFAULT 0,
    status_id             INTEGER      NOT NULL REFERENCES replenishment_request_status_groups (id)
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

CREATE TABLE tasks
(
    id           INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    alias        VARCHAR(50) NOT NULL,
    name         VARCHAR(50) NOT NULL,
    arguments    TEXT NULL,
    status       INT         NOT NULL DEFAULT 0,
    schedule     INT         NOT NULL,
    scheduled_at TIMESTAMP   NOT NULL,
    created_at   TIMESTAMP   NOT NULL,
    updated_at   TIMESTAMP   NOT NULL
);

INSERT INTO roles(name)
VALUES ('Администратор'),
       ('Сёгун'),
       ('Даймё'),
       ('Самурай'),
       ('Инкассатор'),
       ('Контролёр'),
       ('Главный оператор');

INSERT INTO replenishment_request_status_groups(name)
VALUES ('Активные'),
       ('Спорные'),
       ('Выполненные');

INSERT INTO bank_types(name)
VALUES ('Тинькофф'),
       ('Сбербанк');
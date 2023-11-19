CREATE TABLE IF NOT EXISTS stock_price (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36),
    company_id VARCHAR(10) NOT NULL,
    update_date DATETIME NOT NULL,
    price_date DATE NOT NULL,
    open INT UNSIGNED NOT NULL,
    close INT UNSIGNED NOT NULL,
    high INT UNSIGNED NOT NULL,
    low INT UNSIGNED NOT NULL,
    price_change INT NOT NULL,
    change_percent INT NOT NULL,
    volume BIGINT UNSIGNED NOT NULL,
    amount BIGINT UNSIGNED NOT NULL
);

CREATE TABLE IF NOT EXISTS future_price (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36),
    contract_id VARCHAR(4) NOT NULL,
    contract_date VARCHAR(7) NOT NULL,
    update_date DATETIME NOT NULL,
    price_date DATE NOT NULL,
    open INT UNSIGNED NOT NULL,
    close INT UNSIGNED NOT NULL,
    high INT UNSIGNED NOT NULL,
    low INT UNSIGNED NOT NULL,
    volume BIGINT UNSIGNED NOT NULL
);

CREATE TABLE IF NOT EXISTS option_oi (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36),
    contract_id VARCHAR(4) NOT NULL,
    contract_date VARCHAR(7) NOT NULL,
    update_date DATETIME NOT NULL,
    type CHAR(1) NOT NULL,
    strike_price SMALLINT NOT NULL,
    open INT UNSIGNED NOT NULL,
    close INT UNSIGNED NOT NULL,
    high INT UNSIGNED NOT NULL,
    low INT UNSIGNED NOT NULL,
    open_interest INT UNSIGNED NOT NULL,
    settle_price INT UNSIGNED
);
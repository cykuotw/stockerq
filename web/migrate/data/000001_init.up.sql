CREATE TABLE IF NOT EXISTS stock_price (
    id uuid PRIMARY KEY,
    company_id VARCHAR NOT NULL,
    update_date TIMESTAMP WITH TIME ZONE NOT NULL,
    price_date DATE NOT NULL,
    open INTEGER NOT NULL CHECK (open>0),
    close INTEGER NOT NULL CHECK (close>0),
    high INTEGER NOT NULL CHECK (high>0),
    low INTEGER NOT NULL CHECK (low>0),
    change INTEGER NOT NULL,
    change_percent INTEGER NOT NULL,
    volume BIGINT NOT NULL CHECK (volume>0),
    amount BIGINT NOT NULL CHECK (amount>0)
);

CREATE TABLE IF NOT EXISTS future_price (
    id uuid PRIMARY KEY,
    contract_id VARCHAR(4) NOT NULL,
    contract_date VARCHAR(7) NOT NULL,
    update_date TIMESTAMP WITH TIME ZONE NOT NULL,
    price_date DATE NOT NULL,
    open INTEGER NOT NULL CHECK (open>0),
    close INTEGER NOT NULL CHECK (close>0),
    high INTEGER NOT NULL CHECK (high>0),
    low INTEGER NOT NULL CHECK (low>0),
    volume BIGINT NOT NULL CHECK (volume>0)
);

CREATE TABLE IF NOT EXISTS option_oi (
    id uuid PRIMARY KEY,
    contract_id VARCHAR(4) NOT NULL,
    contract_date VARCHAR(7) NOT NULL,
    update_date TIMESTAMP WITH TIME ZONE NOT NULL,
    type CHAR(1) NOT NULL,
    strike_price SMALLINT NOT NULL,
    open INTEGER NOT NULL CHECK (open>0),
    close INTEGER NOT NULL CHECK (close>0),
    high INTEGER NOT NULL CHECK (high>0),
    low INTEGER NOT NULL CHECK (low>0),
    volume BIGINT NOT NULL CHECK (volume>0),
    open_interest INTEGER NOT NULL CHECK (open_interest>0),
    settle_price INTEGER
);
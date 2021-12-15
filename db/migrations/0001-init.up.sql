-- tokens table
create table tokens
(
    symbol     text primary key,
    name       text      not null,
    decimals   integer   not null,
    address    text,
    updated_at timestamp,
    created_at timestamp not null
);
create unique index idx_tokens_address on tokens (symbol);

-- markets table
create table markets
(
    id                   text primary key,
    base_token_symbol    text            not null,
    base_token_name      text            not null,
    base_token_decimals  integer         not null,
    base_token_address   text,
    quote_token_symbol   text            not null,
    quote_token_name     text            not null,
    quote_token_decimals integer         not null,
    quote_token_address  text,
    min_order_size       numeric(32, 18) not null,
    price_decimals       integer         not null,
    amount_decimals      integer         not null,
    maker_fee_rate       numeric(10, 5)  not null,
    taker_fee_rate       numeric(10, 5)  not null,
    is_published         boolean         not null default true,
    updated_at           timestamp,
    created_at           timestamp       not null
);

-- orders table
create table orders
(
    id               text            not null primary key,
    trader_address   text            not null,
    market_id        text            not null,
    side             text            not null,
    price            numeric(32, 18) not null,
    amount           numeric(32, 18) not null,
    status           text            not null,
    type             text            not null,
    available_amount numeric(32, 18) not null,
    confirmed_amount numeric(32, 18) not null,
    canceled_amount  numeric(32, 18) not null,
    pending_amount   numeric(32, 18) not null,
    signature        text            not null,
    flags            text            not null,
    updated_at       timestamp,
    created_at       timestamp       not null
);
create index idx_market_id_status on orders (market_id, status);
create index idx_market_trader_address on orders (trader_address, market_id, status, created_at);

-- trades table
create table trades
(
    id               SERIAL PRIMARY KEY,
    transaction_id   integer         not null,
    transaction_hash text,
    status           text            not null,
    market_id        text            not null,
    maker            text            not null,
    taker            text            not null,
    price            numeric(32, 18) not null,
    amount           numeric(32, 18) not null,
    taker_side       text            not null,
    maker_order_id   text            not null,
    taker_order_id   text            not null,
    executed_at      timestamp,
    updated_at       timestamp,
    created_at       timestamp
);
create index idx_trades_transaction_hash on trades (transaction_hash);
create index idx_trades_taker on trades (taker, market_id);
create index idx_trades_maker on trades (maker, market_id);
create index idx_market_id_status_executed_at on trades (market_id, status, executed_at);

-- transactions table
create table transactions
(
    id           SERIAL PRIMARY KEY,
    market_id    text not null,
    status       text not null,
    hash         text,
    block_number integer,
    gas_limit    integer,
    gas_used     integer,
    gas_price    numeric(32, 18),
    nonce        integer,
    data         text not null,
    executed_at  timestamp,
    updated_at   timestamp,
    created_at   timestamp
);
create index idx_transactions_created_at on transactions (created_at);
create unique index idx_transactions_hash on transactions (hash);

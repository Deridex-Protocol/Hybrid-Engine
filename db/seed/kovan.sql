insert into markets (id,
                     base_token_symbol, base_token_name, base_token_decimals,
                     quote_token_symbol, quote_token_name, quote_token_decimals,
                     min_order_size, price_decimals, amount_decimals,
                     maker_fee_rate, taker_fee_rate, is_published, created_at)
values ('TSLA-USDC',
        'TSLA', 'Tesla Inc', 6,
        'USDC', 'USDC', 6,
        0.0001, 18, 6,
        0.003, 0.003, true, NOW());

insert into tokens (symbol, name, decimals, created_at)
values ('TSLA', 'Tesla Inc', 6, NOW());

insert into tokens (symbol, name, decimals, created_at)
values ('USDC', 'USDC', 6, NOW());

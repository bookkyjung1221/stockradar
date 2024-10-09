-- +goose Up
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    payout NUMERIC(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    sale_amount NUMERIC(10, 2) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    shop_name VARCHAR(100) NOT NULL,
    shop_offset_hour INT NOT NULL
);

-- +goose Down
DROP TABLE transactions;

-- +goose Up
CREATE TABLE exchange_rates (
    currency VARCHAR(3) PRIMARY KEY,
    rate NUMERIC(10, 2) NOT NULL
);

INSERT INTO exchange_rates (currency, rate) VALUES
('AUD', 21.00),
('CNY', 4.50),
('THB', 1.00),
('USD', 29.00);

-- +goose Down
DROP TABLE exchange_rates;

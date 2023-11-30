
CREATE TABLE IF NOT EXISTS products (
    name VARCHAR(255),
    id UUID,
    price DOUBLE PRECISION,
    stock BIGINT,
    min_stock BIGINT,
    supplier_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
    );
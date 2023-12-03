
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

CREATE INDEX IF NOT EXISTS idx_supplier_id ON products(supplier_id);

--

CREATE TABLE IF NOT EXISTS sales (
    id UUID,
    product_name varchar(255),
    product_id UUID,
    quantity BIGINT,
    total DOUBLE PRECISION,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

create INDEX IF NOT EXISTS idx_product_id on sales(product_id);

--

CREATE TABLE IF NOT EXISTS users (
     fullname varchar(255),
     email varchar(255) UNIQUE,
     id UUID,
     password varchar(255),
     created_at TIMESTAMP,
     updated_at TIMESTAMP,
     PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_user_email on users(email);

-- complex query

SELECT
    products.id,
    products.name,
    products.stock,
    products.min_stock,
    products.supplier_id,
    products.price,
    AVG(sales.quantity) AS avg_montly_sales
FROM
    products
        JOIN
    sales ON products.id = sales.product_id
WHERE
        products.stock < products.min_stock
GROUP BY
    products.id, products.name, products.stock;
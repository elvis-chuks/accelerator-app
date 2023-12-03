
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

--

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
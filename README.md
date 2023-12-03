# Accelerator App
**Assesment B**

### Project Name: Advanced Inventory Management API (AIMA)
### Project Description
The Advanced Inventory Management API (AIMA) is designed to facilitate the management of product inventory through a RESTful API. 
The service offers CRUD operations on products and introduces an endpoint for a complex report, 
aiming to aggregate data for informed decision-making, such as calculating restocking needs based on historical sales data.

#### Language of choice = Golang

_**Postman Documentation**: [Here](https://documenter.getpostman.com/view/7190909/2s9YeK5AZL)_

### Configuration
* Setup Database
  * Create Tables
  when you run the app for the first time, it creates the tables and indexes it needs to run fine.
* Setup env vars
    ```bash
      $ export PORT=5001
      $ export DB_URL=$DB_URL
      $ export SIGNING_KEY=$SIGNING_KEY
    ```
#### Run project

```bash
$ make build
$ make run-container DB_URL=$DB_URL SIGNING_KEY=$SIGNING_KEY
```
the docker container should now be running on http://127.0.0.1:5001

[//]: # (##### Tests)

[//]: # ()
[//]: # (To test all packages in the project run)

[//]: # (```bash)

[//]: # ($ make test)

[//]: # (```)

[//]: # ()
[//]: # (To test a specific repository function run)

[//]: # (```bash)

[//]: # ($ make test-repo-fxn fxn=functionName)

[//]: # (```)

refer to Makefile for useage description
##### Complex Query
```sql
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
  products.id, products.name, products.stock, products.min_stock
HAVING
  AVG(sales.quantity) < products.min_stock;
```
In the query above, I retrieve records from the product table,
I join it to the sales table using the one-to-many product id relationship
with the sales table, the condition for the join is for products where the 
current stock is less than the minimum threshold, I find the average monthly sales of each product
and to organize and present this information, I group them starting from the id, name and stock.

This query is particularly helpful in facilitating informed decision-making based on product sales and inventory levels.

Please reach out to me for a db_url to test with (something that has already been seeded with test data).

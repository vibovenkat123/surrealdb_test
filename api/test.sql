-- NoSQL-esque modeling, schemaless
DELETE food;
CREATE food:pasta SET toppings=['basil'], rating=9;
CREATE food:pizza SET toppings=['cheese'], rating=10;

-- where clause
SELECT * FROM food WHERE rating > 9;
SELECT * FROM food WHERE toppings CONTAINS 'basil';
-- update array
UPDATE food SET toppings += ["pineapple"] WHERE id = "food:pasta"; -- :)

SELECT * FROM food;
-- SQL schema similar to postgres

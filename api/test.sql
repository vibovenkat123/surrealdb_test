DELETE food;
CREATE food:pasta SET toppings=['basil'], rating=9;
CREATE food:pizza SET toppings=['cheese'], rating=10;

-- where clause
SELECT * FROM food WHERE rating > 10;
SELECT * FROM food WHERE toppings CONTAINS 'basil';
-- update array
UPDATE food SET toppings += ["pineapple"] WHERE id = "food:pizza"; -- :)

SELECT * FROM food;

-- one to one
UPDATE food:pizza SET side = food:pasta;

SELECT side.rating, side.toppings FROM food:pizza;

-- one to many
CREATE drink:coke
SET name="Coca Cola", soda=true, price = 2, rating = 7;

CREATE drink:lemonade
SET name="Lemonade", soda=false, price = 1, rating = 9;

UPDATE food:pasta SET drinks=[drink:lemonade, drink:coke];

-- Many to Many
-- graph-esque
RELATE food:pizza->went_well_with->drink:lemonade;

SELECT ->went_well_with->drink FROM food;
SELECT <-went_well_with<-food FROM drink;

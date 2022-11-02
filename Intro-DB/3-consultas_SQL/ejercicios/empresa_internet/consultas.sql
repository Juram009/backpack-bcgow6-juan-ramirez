--Plantear 10 consultas SQL que se podrían realizar a la base de datos. Expresar las sentencias.

--Mostrar todos los clientes que tiene la compañia
SELECT dni, nombre, apellido, fecha_nacimiento, provincia, ciudad FROM empresa_internet.cliente

--Mostrar todos los clientes que tengan el plan de internet con id 3
SELECT nombre, apellido FROM empresa_internet.cliente WHERE id_plan=3

--Mostrar los planes que tengan al menos el 10% de descuento
SELECT * FROM empresa_internet.plan_internet WHERE descuento >=10

--Mostrar los clientes que tengan mas de 35 años
SELECT nombre, apellido FROM empresa_internet.cliente WHERE fecha_nacimiento < '1987-11-01'

--Mostrar los planes que tengan un costo mayor a $2000
SELECT * FROM empresa_internet.plan_internet WHERE precio >2000

--Mostrar los nombres de los clientes alfabeticamente
SELECT nombre, apellido FROM empresa_internet.cliente ORDER BY nombre ASC

--Mostrar el cliente mas joven
SELECT nombre, apellido FROM empresa_internet.cliente ORDER BY fecha_nacimiento DESC LIMIT 1

--Mostrar cuantos clientes tiene cada plan
SELECT id_plan, COUNT(id_plan) FROM empresa_internet.cliente GROUP BY(id_plan)

--Mostrar planes de internet sin descuento
SELECT * FROM empresa_internet.plan_internet WHERE descuento = 0

--Mostrar las ciudades en las que residen los clientes
SELECT ciudad FROM empresa_internet.cliente 
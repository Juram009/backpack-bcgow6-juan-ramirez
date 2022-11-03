--Agregar una película a la tabla movies.
INSERT INTO movies_db.movies (title, rating, awards, release_date, length, genre_id) VALUES("Pulp Fiction", 8.9, 9, '1994-10-14', 154, 4)

--Agregar un género a la tabla genres.
INSERT INTO movies_db.genres (name, ranking, active) VALUES ("Policiaca", 13, 1)

--Asociar a la película del punto 1. con el género creado en el punto 2.
UPDATE movies_db.movies SET genre_id = 13 WHERE id = 22

--Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
UPDATE movies_db.actors SET favorite_movie_id = 13 WHERE id = 3

--Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_db.movies_copy (id INTEGER, title VARCHAR(40), rating FLOAT, awards INTEGER, release_date DATE, length INTEGER, genre_id INTEGER);
INSERT INTO movies_db.movies_copy
SELECT id, title, rating, awards, release_date, length, genre_id FROM movies_db.movies

--Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM movies_db.movies_copy WHERE awards < 5;

--Obtener la lista de todos los géneros que tengan al menos una película.
SELECT name FROM movies_db.genres JOIN movies_db.movies WHERE movies.genre_id = genres.id

--Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT first_name as Nombre, last_name as Apellido FROM movies_db.actors JOIN movies_db.movies WHERE favorite_movie_id = movies.id AND awards > 3

--Crear un índice sobre el nombre en la tabla movies.
ALTER TABLE movies_db.movies ADD INDEX index_titulo (title ASC) VISIBLE;

--Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;

--En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
--Sí, el titulo se utiliza en algunas consultas, por lo que ahorraria tiempo y recursos al no requerir buscar en todos los registros que tengan las tablas

--¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
--En series crearía un indice para titulo, al igual que el de movies, esto por la misma razón, hay consultas que utilizan este dato

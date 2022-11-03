--Primer punto
--Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
CREATE TEMPORARY TABLE movies_db.twd (episode varchar(40), season varchar(40));
INSERT INTO twd
SELECT episodes.title, seasons.title  FROM movies_db.series JOIN movies_db.seasons JOIN movies_db.episodes WHERE series.id = seasons.serie_id AND seasons.id = episodes.season_id AND series.title = "The Walking Dead";

--Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.
SELECT episode FROM movies_db.twd WHERE season = "Primer Temporada"

--Segundo punto
--En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 
ALTER TABLE movies_db.movies ADD INDEX index_titulo (title ASC) VISIBLE;
SHOW INDEX FROM movies;

--Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.
--Creé un indice de la tabla movies en el campo de titulo para los casos en los que se necesite buscar por titulo de pelicula
--Mostrar el título y el nombre del género de todas las series.
SELECT title as Titulo, genres.name as Genero FROM series JOIN genres WHERE series.genre_id = genres.id

--Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT title as Titulo, actors.first_name as Nombre, actors.last_name as Apellido FROM episodes JOIN actor_episode JOIN actors WHERE actor_episode.episode_id = episodes.id AND actors.id = actor_episode.actor_id

--Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.title as Titulo, COUNT(series.id) as Temporadas FROM series JOIN seasons WHERE seasons.serie_id = series.id group by (series.id)

--Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT name as Genero, COUNT(genres.id) as Peliculas FROM genres JOIN movies WHERE genres.id = movies.genre_id GROUP BY(genres.id) HAVING Peliculas >= 3 

--Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT actors.first_name as Nombre, actors.last_name as Apellido FROM movies JOIN actor_movie JOIN actors WHERE actor_movie.movie_id = movies.id AND actor_movie.actor_id = actors.id AND title LIKE "La Guerra de las galaxias%" 

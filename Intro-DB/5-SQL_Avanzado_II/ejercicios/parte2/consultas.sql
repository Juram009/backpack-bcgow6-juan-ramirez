--Listar los datos de los autores.
SELECT * FROM biblioteca_db.autor

--Listar nombre y edad de los estudiantes
SELECT nombre, apellido, edad FROM biblioteca_db.estudiante

--¿Qué estudiantes pertenecen a la carrera informática?
SELECT nombre, apellido FROM biblioteca_db.estudiante WHERE carrera = "informatica"

--¿Qué autores son de nacionalidad francesa o italiana?
SELECT nombre FROM biblioteca_db.autor WHERE nacionalidad = "Francia" OR nacionalidad = "Italia"

--¿Qué libros no son del área de internet?
SELECT titulo FROM biblioteca_db.libro WHERE area != "internet"

--Listar los libros de la editorial Salamandra.
SELECT titulo FROM biblioteca_db.libro WHERE editorial = "Salamandra"

--Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT nombre, apellido, edad FROM biblioteca_db.estudiante WHERE edad > (SELECT AVG(edad) FROM biblioteca_db.estudiante);

--Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT nombre, apellido FROM biblioteca_db.estudiante WHERE apellido LIKE "G%"

--Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT nombre FROM biblioteca_db.libro JOIN biblioteca_db.libro_autor JOIN biblioteca_db.autor WHERE libro.id_libro = libro_autor.id_libro AND autor.id_autor = libro_autor.id_autor AND titulo="El universo: Guia de viaje"

--¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT titulo FROM biblioteca_db.libro JOIN biblioteca_db.prestamo JOIN biblioteca_db.estudiante WHERE libro.id_libro = prestamo.id_libro AND estudiante.id_lector = prestamo.id_lector AND nombre = "“Filippo" AND apellido = "Galli"

--Listar el nombre del estudiante de menor edad.
SELECT nombre, apellido FROM biblioteca_db.estudiante WHERE edad < 18

--Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT * FROM biblioteca_db.libro JOIN biblioteca_db.prestamo JOIN biblioteca_db.estudiante WHERE libro.id_libro = prestamo.id_libro AND estudiante.id_lector = prestamo.id_lector AND area = "Bases de datos"

--Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT titulo FROM biblioteca_db.libro JOIN biblioteca_db.libro_autor JOIN biblioteca_db.autor WHERE libro.id_libro = libro_autor.id_libro AND autor.id_autor = libro_autor.id_autor AND nombre = "J.K. Rowling"

--Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT * FROM biblioteca_db.libro JOIN biblioteca_db.prestamo WHERE libro.id_libro = prestamo.id_libro AND fecha_devolucion = "2021-07-16"


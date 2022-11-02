--Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT nombre, puesto, localidad FROM empleado JOIN departamento WHERE empleado.depto_nro = departamento.depto_nro

--Visualizar los departamentos con más de cinco empleados.
SELECT nombre_depto as Departamento, COUNT(depto_nro) as Empleados FROM departamento JOIN empleado WHERE empleado.depto_nro = departamento.depto_nro GROUP BY depto_nro HAVING Empleados >5

--Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
SELECT nombre, salario, nombre_depto as Departamento FROM empleado JOIN departamento WHERE empleado.depto_nro = departamento.depto_nro AND puesto = "Presidente"

--Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT nombre, puesto, salario FROM empleado JOIN departamento WHERE empleado.depto_nro = departamento.depto_nro AND nombre_depto = "Contabilidad" ORDER BY nombre ASC

--Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT nombre FROM empleado ORDER BY salario ASC LIMIT 1

--Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT nombre, puesto, salario FROM empleado JOIN departamento WHERE empleado.depto_nro = departamento.depto_nro AND nombre_depto = "Ventas" ORDER BY salario DESC LIMIT 1
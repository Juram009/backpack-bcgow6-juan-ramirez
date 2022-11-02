/*==============================================================*/
/* DBMS name:      Sybase SQL Anywhere 12                       */
/* Created on:     1/11/2022 6:37:16 p. m.                      */
/*==============================================================*/


DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet;
USE empresa_internet;

/*==============================================================*/
/* Table: PLAN_INTERNET                                         */
/*==============================================================*/
CREATE TABLE plan_internet 
(
   id_plan              INTEGER                        NOT NULL        AUTO_INCREMENT,
   velocidad            FLOAT                          NOT NULL,
   precio               FLOAT                          NOT NULL,
   descuento            FLOAT                          NOT NULL,
   CONSTRAINT  pk_plan_internet PRIMARY KEY (id_plan)
);

/*==============================================================*/
/* Index: PLAN_INTERNET_PK                                      */
/*==============================================================*/
CREATE UNIQUE INDEX plan_internet_pk ON plan_internet (
id_plan ASC
);

/*==============================================================*/
/* INSERTION OF INTERNET PLANS									*/
/*==============================================================*/
INSERT INTO plan_internet (id_plan, velocidad, precio, descuento) 
VALUES (1, 100, 1500 ,0), (2, 150, 1700, 0), (3, 200, 2000, 10), (4, 250, 2300, 10), (5, 300, 2500, 20);

/*==============================================================*/
/* Table: CLIENTE                                               */
/*==============================================================*/
CREATE TABLE cliente 
(
   dni                  INTEGER                        NOT NULL,
   id_plan              INTEGER                        NOT NULL,
   nombre               LONG VARCHAR                   NOT NULL,
   apellido             LONG VARCHAR                   NOT NULL,
   fecha_nacimiento     DATE                           NOT NULL,
   provincia            LONG VARCHAR                   NOT NULL,
   ciudad               LONG VARCHAR                   NOT NULL,
   CONSTRAINT pk_cliente PRIMARY KEY (dni)
);

/*==============================================================*/
/* Index: CLIENTE_PK                                            */
/*==============================================================*/
CREATE UNIQUE INDEX cliente_pk ON cliente (
dni ASC
);

/*==============================================================*/
/* Index: RELATIONSHIP_1_FK                                     */
/*==============================================================*/
CREATE INDEX relationship_1_fk ON cliente (
id_plan ASC
);

ALTER TABLE cliente
   ADD CONSTRAINT fk_cliente_relations_plan_int FOREIGN KEY (id_plan)
      REFERENCES plan_internet (id_plan)
      ON UPDATE RESTRICT
      ON DELETE RESTRICT;
      
/*==============================================================*/
/* INSERTION OF CLIENTS											*/
/*==============================================================*/
INSERT INTO cliente (dni, id_plan, nombre, apellido, fecha_nacimiento, provincia, ciudad) 
VALUES ('188151834', 2, 'Nicholas', 'Olson', '1983-04-07', 'Loreto', 'Carson City'), 
('397231862',3,'Omar','Wheeler','1989-03-26','Noord Brabant','Tarakan'), 
('591921758',4,'Emi','Gilbert','1973-12-30','Aberdeenshire','San Antonio'), 
('493217993',2,'Kadeem','King','2001-11-19','Lima','Tacurong'), 
('393650628',3,'Kylee','Andrews','1977-02-26','Western Visayas','Buguma'),
('668847981',3,'April','Conner','1991-04-03','West Sumatra','Rzeszów'),
('864753387',1,'Drake','Olsen','2003-03-16','Västra Götalands län','Provo'), 
('545827782',4,'Josephine','Townsend','1971-11-02','Carinthia','Maiduguri'), 
('685570774',5,'Cathleen','Fitzpatrick','1991-03-05','Special Capital Region of Jakarta','Davoli'), 
('632334058',2,'Kibo','Hickman','1985-09-24','Gyeonggi','Bad Kreuznach');


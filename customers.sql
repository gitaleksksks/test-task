-- Database: customersdb

-- DROP DATABASE customersdb;

CREATE DATABASE customersdb
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.customers

-- DROP TABLE public.customers;

CREATE TABLE customers
(
    id integer NOT NULL,
    firstname character varying(100),
    lastname character varying(100),
    birthdate date,
    gender text,
    email text,
    homeaddress character varying(200),
    CONSTRAINT customers_pkey PRIMARY KEY (id)
)

WITH (
 OIDS=FALSE
);

ALTER TABLE customers
    OWNER to postgres;
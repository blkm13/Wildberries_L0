CREATE DATABASE "L0"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


CREATE SCHEMA l0_model
    AUTHORIZATION postgres;


CREATE TABLE l0_model."Delivery "
(
    name character varying,
    phone character varying,
    zip character varying,
    city character varying,
    address character varying,
    region character varying,
    email character varying
);

ALTER TABLE l0_model."Delivery "
    OWNER to postgres;
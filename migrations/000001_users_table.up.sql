CREATE TYPE status AS ENUM ('admin', 'manager', 'customer');

CREATE TABLE users
(
    id              UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    first_name      varchar   NOT NULL,
    last_name       varchar   NOT NULL,
    email           varchar   NOT NULL  UNIQUE,
    age             INT NOT NULL,
    region          varchar NOT NULL,
    status status
);


INSERT INTO users (first_name, last_name, email, age, region, status)
VALUES ('Pupa', 'Pupovna', 'pupa@gmail.com', 23, 'BY', 'customer'),
       ('Lupa', 'Lupovna', 'lupa@gmail.com', 22,'BY', 'admin'),
       ('Gerald', 'Izrivii', 'miasnik@gmail.com', 132,'PL', 'manager'),
       ('Jonathan', 'Jostar', 'jojo@gmail.com', 27,'JP', 'manager'),
       ('Naruto', 'Ydzumaki', 'naruto@gmail.com', 17,'JP', 'manager'),
       ('Bruce', 'Wayne', 'batman1234@gmail.com', 54,'BY', 'customer');



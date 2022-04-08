CREATE TABLE users
(
    id              UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    first_name      varchar   NOT NULL,
    last_name       varchar   NOT NULL,
    email           varchar   NOT NULL  UNIQUE,
    age             INT NOT NULL
);


INSERT INTO users (first_name, last_name, email, age)
VALUES ('Pupa', 'Pupovna', 'pupa@gmail.com', 23),
       ('Lupa', 'Lupovna', 'lupa@gmail.com', 22),
       ('Gerald', 'Izrivii', 'miasnik@gmail.com', 132),
       ('Jonathan', 'Jostar', 'jojo@gmail.com', 27),
       ('Naruto', 'Ydzumaki', 'naruto@gmail.com', 17),
       ('Bruce', 'Wayne', 'batman1234@gmail.com', 54);



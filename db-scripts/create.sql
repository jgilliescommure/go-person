DROP TABLE IF EXISTS person;
CREATE TABLE person (
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    email VARCHAR(128),
    PRIMARY KEY (`id`)
);

INSERT INTO person (first_name, last_name, email)
VALUES 
('Fred', 'Murphy', 'f.m@gmail.com'),
('Nancy', 'Murphy', 'n.m@gmail.com'),
('George', 'Murphy', 'g.m@gmail.com'),
('Fred', 'Flintstore', 'ff@gmail.com'),
('Willma', 'Flintstore', 'wf@gmail.com');
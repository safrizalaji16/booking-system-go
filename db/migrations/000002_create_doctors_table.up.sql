CREATE TABLE IF NOT EXISTS doctors(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   schedule jsonb,
   speciality_id INT NOT NULL,
   CONSTRAINT fk_specialities FOREIGN KEY(speciality_id) REFERENCES specialities(id)
);
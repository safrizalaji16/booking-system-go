CREATE TABLE IF NOT EXISTS bookings(
  id serial PRIMARY KEY,
  doctor_id INT NOT NULL,
  patient_id INT NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  CONSTRAINT fk_doctors FOREIGN KEY(doctor_id) REFERENCES doctors(id),
  CONSTRAINT fk_patients FOREIGN KEY(patient_id) REFERENCES patients(id)
);
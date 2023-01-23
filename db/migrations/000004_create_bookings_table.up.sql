CREATE TABLE IF NOT EXISTS bookings(
  id serial PRIMARY KEY,
  DoctorId INT NOT NULL,
  PatientId INT NOT NULL,
  startDate DATE NOT NULL,
  endDate DATE NOT NULL,
  CONSTRAINT fk_doctors FOREIGN KEY(DoctorId) REFERENCES doctors(id),
  CONSTRAINT fk_patients FOREIGN KEY(PatientId) REFERENCES patients(id)
);
CREATE TABLE employee_clock_histories (
  id SERIAL PRIMARY KEY, 
  employee_id TEXT UNIQUE NOT NULL, 
  job_title TEXT NOT NULL, 
  update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_employee_clock_histories_employee_id ON employee_clock_histories(employee_id);
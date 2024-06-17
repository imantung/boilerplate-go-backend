CREATE TABLE employee_clock_histories (
  id SERIAL PRIMARY KEY, 
  employee_id TEXT UNIQUE NOT NULL, 
  clock_in_at TIMESTAMP, 
  clock_out_at TIMESTAMP, 
  work_duration TEXT, 
  work_duration_minutes INTEGER,
  deleted_at TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_employee_clock_histories_employee_id ON employee_clock_histories(employee_id);
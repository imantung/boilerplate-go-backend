CREATE TABLE employees (
  id SERIAL PRIMARY KEY, 
  employee_name TEXT UNIQUE NOT NULL, 
  job_title TEXT NOT NULL, 
  last_clock_in_at TIMESTAMP, 
  last_clock_out_at TIMESTAMP, 
  deleted_at TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
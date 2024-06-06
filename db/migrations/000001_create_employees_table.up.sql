CREATE TABLE employees (
  id SERIAL PRIMARY KEY, 
  employee_name TEXT UNIQUE NOT NULL, 
  job_title TEXT NOT NULL, 
  update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
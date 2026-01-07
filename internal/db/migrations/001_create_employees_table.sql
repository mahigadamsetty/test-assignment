CREATE TABLE IF NOT EXISTS employees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name  TEXT NOT NULL,
    job_title  TEXT NOT NULL,
    salary     INTEGER NOT NULL CHECK (salary > 0),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL, -- Assuming emails are unique
  username VARCHAR(50) UNIQUE NOT NULL, -- Assuming usernames are unique
  password VARCHAR(250) NOT NULL,
  created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (first_name, last_name, email, username, password) 
VALUES 
    ('Magda', 'Renac', 'mrenac0@businesswire.com', 'mrenac0', 'test_account');

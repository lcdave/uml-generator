CREATE TABLE IF NOT EXISTS request(
id SERIAL PRIMARY KEY,
service_id INT,
reaction_time FLOAT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT fk_service
      FOREIGN KEY(service_id) 
	  REFERENCES services(id)
);
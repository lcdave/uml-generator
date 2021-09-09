CREATE TABLE IF NOT EXISTS emails(
id SERIAL PRIMARY KEY,
service_id INT,
recipient VARCHAR(45),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT fk_service
      FOREIGN KEY(service_id) 
	  REFERENCES services(id)
);
CREATE TABLE IF NOT EXISTS contacts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL
);

INSERT INTO contacts (name, email, phone) 
SELECT 'Juan Perez', 'juan.perez@example.com', '123456789' 
WHERE NOT EXISTS (SELECT * FROM contacts WHERE email = 'juan.perez@example.com');

INSERT INTO contacts (name, email, phone) 
SELECT 'Ana Garcia', 'ana.garcia@example.com', '987654321' 
WHERE NOT EXISTS (SELECT * FROM contacts WHERE email = 'ana.garcia@example.com');

INSERT INTO contacts (name, email, phone) 
SELECT 'Carlos López', 'carlos.lopez@example.com', '456789123' 
WHERE NOT EXISTS (SELECT * FROM contacts WHERE email = 'carlos.lopez@example.com');

CREATE TABLE ips (
    id CHAR(36) PRIMARY KEY,
    ip_address VARCHAR(15) NOT NULL,
    note VARCHAR(64),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    last_update DATETIME DEFAULT NULL,
    status ENUM('active', 'blocked', 'suspended') NOT NULL
);

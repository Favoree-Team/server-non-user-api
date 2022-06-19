CREATE TABLE IF NOT EXISTS personal_datas (
    id VARCHAR(150) NOT NULL,
    ip_address VARCHAR(100) NOT NULL,
    device_access VARCHAR(100) NOT NULL,
    name VARCHAR(150),
    email VARCHAR(150),
    created_at DATETIME,
    updated_at DATETIME,
    Primary Key (id)
);

CREATE TABLE IF NOT EXISTS log_request_transactions (
    id VARCHAR(150) NOT NULL,
    personal_data_id VARCHAR(100) NOT NULL,
    sender_number VARCHAR(100) NOT NULL,
    sender_wallet VARCHAR(100) NOT NULL,
    receiver_name VARCHAR(100) NOT NULL,
    receiver_number VARCHAR(100) NOT NULL,
    receiver_wallet VARCHAR(100) NOT NULL,
    amount_transfer int NOT NULL,
    admin_fee int NOT NULL,
    amount_received int NOT NULL,
    status ENUM('pending', 'paid', 'success', 'failed', 'expired') NOT NULL,
    done boolean NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    Primary Key (id),
    FOREIGN KEY (personal_data_id) REFERENCES personal_datas(id)
);
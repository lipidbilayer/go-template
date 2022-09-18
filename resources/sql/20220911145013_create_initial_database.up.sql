CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL,
    type varchar(25) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name  VARCHAR(255),
    location_id INT,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

INSERT into users (username, password, name, role) 
     VALUES ('admin' ,crypt('password', gen_salt('bf')), 'admin', 'admin');

CREATE TABLE vehicles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plate_number VARCHAR(10) NOT NULL,
    variant VARCHAR(10) DEFAULT '',
    driver_name VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
); 



-- CREATE TABLE attendances (
--     id SERIAL PRIMARY KEY,
--     vehicle_id UUID NOT NULL,
--     driver_name VARCHAR(255) NULL,
--     status SMALLINT NOT NULL,
--     checkin_at TIMESTAMP,
--     checkout_at TIMESTAMP,
--     created_at TIMESTAMP DEFAULT NOW(),
--     updated_at TIMESTAMP DEFAULT NOW(),
--     deleted_at TIMESTAMP, 
-- )
-- CREATE INDEX attendance_vehicle_id_sq ON attendance(vehicle_id);
-- CREATE INDEX attendance_vehicle_id_and_status_sq ON attendance(vehicle_id, status);
CREATE TABLE IF NOT EXISTS games(
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false,
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR (50) NOT NULL,
    url VARCHAR(20) NOT NULL,
    platform VARCHAR (15) NOT NULL,
    description VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'NOT_REGISTERED'
);
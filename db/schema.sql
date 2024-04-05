-- Active: 1712343322333@@127.0.0.1@5432@todo
--
-- Struct for the table `todo`
--
CREATE TABLE IF NOT EXISTS todo(
    id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    complete VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
)
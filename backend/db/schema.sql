-- E-STORE Database Schema
-- Run this file to set up all tables: psql -d estore -f schema.sql

-- Users table: stores everyone who signs up
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,           -- auto-incrementing unique ID
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(150) UNIQUE NOT NULL, -- no two users can have same email
    password   TEXT NOT NULL,               -- will store hashed password (never plain text)
    created_at TIMESTAMP DEFAULT NOW()
);

-- Products table: everything available in the store
CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL,    -- e.g. 49.99
    stock       INT DEFAULT 0,              -- how many are available
    created_at  TIMESTAMP DEFAULT NOW()
);

-- Orders table: when a user places an order
CREATE TABLE IF NOT EXISTS orders (
    id         SERIAL PRIMARY KEY,
    user_id    INT REFERENCES users(id),    -- links to the user who placed it
    total      NUMERIC(10, 2) NOT NULL,
    status     VARCHAR(50) DEFAULT 'pending', -- pending, shipped, delivered
    created_at TIMESTAMP DEFAULT NOW()
);

-- Order items: each product inside an order
CREATE TABLE IF NOT EXISTS order_items (
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders(id),   -- links to which order
    product_id INT REFERENCES products(id), -- links to which product
    quantity   INT NOT NULL,
    price      NUMERIC(10, 2) NOT NULL      -- price at time of purchase
);

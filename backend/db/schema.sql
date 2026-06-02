-- E-STORE Database Schema
-- Run this file to set up all tables: psql -d estore -f schema.sql

-- Users table: stores everyone who signs up
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(150) UNIQUE NOT NULL,
    password   TEXT NOT NULL,                 -- stores hashed password (never plain text)
    role       VARCHAR(20) DEFAULT 'customer', -- customer or admin
    created_at TIMESTAMP DEFAULT NOW()
);

-- Products table: everything available in the store
CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    category    VARCHAR(100) NOT NULL,         -- e.g. Clothing, Footwear, Accessories
    price       NUMERIC(10, 2) NOT NULL,
    stock       INT DEFAULT 0,
    image_url   TEXT,                          -- link to product image
    created_at  TIMESTAMP DEFAULT NOW()
);

-- Orders table: when a user places an order
CREATE TABLE IF NOT EXISTS orders (
    id         SERIAL PRIMARY KEY,
    user_id    INT REFERENCES users(id),
    total      NUMERIC(10, 2) NOT NULL,
    status     VARCHAR(50) DEFAULT 'pending',  -- pending, shipped, delivered, cancelled
    created_at TIMESTAMP DEFAULT NOW()
);

-- Order items: each product inside an order
CREATE TABLE IF NOT EXISTS order_items (
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders(id),
    product_id INT REFERENCES products(id),
    quantity   INT NOT NULL,
    price      NUMERIC(10, 2) NOT NULL         -- price at time of purchase
);

-- Initial database schema for Elelden
-- Run this script to create the database structure

-- Create database
CREATE DATABASE IF NOT EXISTS elelden_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE elelden_db;

-- Users table
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone VARCHAR(20),
    avatar VARCHAR(500),
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    role ENUM('user', 'admin', 'moderator') DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    INDEX idx_email (email),
    INDEX idx_username (username),
    INDEX idx_deleted_at (deleted_at)
);

-- Categories table with type for different category systems
CREATE TABLE categories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    slug VARCHAR(200) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(100),
    type ENUM('emlak', 'arac', 'ikinci_el') NOT NULL,
    parent_id BIGINT UNSIGNED NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE SET NULL,
    INDEX idx_parent_id (parent_id),
    INDEX idx_type (type),
    INDEX idx_slug (slug),
    INDEX idx_deleted_at (deleted_at)
);

-- Listings table with category-specific JSON fields
CREATE TABLE listings (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    price DECIMAL(15,2) NOT NULL,
    currency VARCHAR(3) DEFAULT 'TRY',
    location VARCHAR(300),
    user_id BIGINT UNSIGNED NOT NULL,
    category_id BIGINT UNSIGNED NOT NULL,
    status ENUM('active', 'sold', 'inactive') DEFAULT 'active',
    view_count INT UNSIGNED DEFAULT 0,
    is_promoted BOOLEAN DEFAULT FALSE,
    emlak_details JSON,
    arac_details JSON,
    ikinci_el_details JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT,
    INDEX idx_user_id (user_id),
    INDEX idx_category_id (category_id),
    INDEX idx_status (status),
    INDEX idx_price (price),
    INDEX idx_location (location),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at),
    FULLTEXT idx_search (title, description)
);

-- Listing images table
CREATE TABLE listing_images (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    listing_id BIGINT UNSIGNED NOT NULL,
    url VARCHAR(1000) NOT NULL,
    alt VARCHAR(300),
    order_index INT UNSIGNED DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE,
    INDEX idx_listing_id (listing_id),
    INDEX idx_order (order_index),
    INDEX idx_deleted_at (deleted_at)
);

-- Favorites table
CREATE TABLE favorites (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    listing_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_listing (user_id, listing_id),
    INDEX idx_user_id (user_id),
    INDEX idx_listing_id (listing_id),
    INDEX idx_deleted_at (deleted_at)
);

-- Messages table
CREATE TABLE messages (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_id BIGINT UNSIGNED NOT NULL,
    receiver_id BIGINT UNSIGNED NOT NULL,
    listing_id BIGINT UNSIGNED NULL,
    content TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE,
    INDEX idx_sender_id (sender_id),
    INDEX idx_receiver_id (receiver_id),
    INDEX idx_listing_id (listing_id),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at)
);

-- Insert initial categories
-- EMLAK Categories
INSERT INTO categories (name, slug, description, icon, type, parent_id) VALUES
-- Main categories
('Emlak', 'emlak', 'Gayrimenkul ilanları', 'home', 'emlak', NULL),
('Araç', 'arac', 'Araç ilanları', 'car', 'arac', NULL),
('İkinci El', 'ikinci-el', 'İkinci el ürünler', 'shopping-bag', 'ikinci_el', NULL);

-- Emlak subcategories
INSERT INTO categories (name, slug, description, icon, type, parent_id) VALUES
('Konut', 'konut', 'Ev, daire, villa', 'home', 'emlak', 1),
('İş Yeri', 'is-yeri', 'Ofis, mağaza, fabrika', 'building', 'emlak', 1),
('Arsa', 'arsa', 'Konut, ticari, tarla arsası', 'map', 'emlak', 1),
('Bina', 'bina', 'Apartman, iş hanı', 'building-2', 'emlak', 1);

-- Araç subcategories  
INSERT INTO categories (name, slug, description, icon, type, parent_id) VALUES
('Otomobil', 'otomobil', 'Binek araçlar', 'car', 'arac', 2),
('Motosiklet', 'motosiklet', 'Motorsiklet ve scooter', 'bike', 'arac', 2),
('Ticari Araç', 'ticari-arac', 'Kamyonet, minibüs', 'truck', 'arac', 2),
('Tarım & İş Makinesi', 'tarim-is-makinesi', 'Traktör ve iş makineleri', 'tractor', 'arac', 2);

-- İkinci El subcategories
INSERT INTO categories (name, slug, description, icon, type, parent_id) VALUES
('Elektronik', 'elektronik', 'Telefon, bilgisayar, TV', 'smartphone', 'ikinci_el', 3),
('Giyim', 'giyim', 'Kadın, erkek, çocuk giyim', 'shirt', 'ikinci_el', 3),
('Ev & Yaşam', 'ev-yasam', 'Mobilya, dekorasyon', 'sofa', 'ikinci_el', 3),
('Spor & Outdoor', 'spor-outdoor', 'Spor malzemeleri', 'dumbbell', 'ikinci_el', 3);

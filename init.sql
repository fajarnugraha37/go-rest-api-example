
-- Create books table
CREATE TABLE books (
    id binary(16) DEFAULT UUID_TO_BIN(UUID()) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    book_status INT NOT NULL,
    book_attrs JSON NOT NULL
);
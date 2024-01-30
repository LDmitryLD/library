CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    second_name VARCHAR(255),
    rent_count INT
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    author_id INT,
    title VARCHAR(100) NOT NULL,
    status VARCHAR(100) NOT NULL,
    CONSTRAINT author_fk FOREIGN KEY(author_id) REFERENCES authors(id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    second_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS rented_books (
    user_id INT,
    book_id INT UNIQUE,
    borrow_date DATE,
    PRIMARY KEY (user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);
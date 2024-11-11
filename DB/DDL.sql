CREATE DATABASE db_library;

CREATE TABLE MST_USERS (
    id varchar(50) primary key,
    full_name varchar(255),
    email varchar(255) unique,
    password varchar(255),
    phone_number varchar(16) unique,
    address varchar(255),
    role varchar(10),
    created_at date,
    updated_at date
);

CREATE TABLE MST_CATEGORIES (
    id varchar(50) primary key,
    category_name varchar(255)
);

CREATE TABLE MST_AUTHOR_DETAILS (
    id varchar(50) primary key,
    author_name varchar(255),
    email varchar(255),
    phone_number varchar(16)
);

CREATE TABLE MST_PUBLISHER_DETAILS (
    id varchar(50) primary key,
    publisher_name varchar(255),
    email varchar(255),
    phone_number varchar(16)
);

CREATE TABLE MST_BOOKS (
    id varchar(50) primary key,
    title varchar(255),
    publication_year date,
    stock int,
    total_pages int,
	publisher_id varchar(50) REFERENCES MST_PUBLISHER_DETAILS(id),
    author_id varchar(50) REFERENCES MST_AUTHOR_DETAILS(id),
    category_id varchar(50) REFERENCES MST_CATEGORIES(id),
    created_at date,
    updated_at date,
);

CREATE TABLE TRN_BOOKS (
    id varchar(50) primary key,
	user_id varchar(50) REFERENCES MST_USERS(id),
    status varchar(15),
    penalty int,
    borrow_date date,
    return_date date,
    return_actual date,
    created_at date,
    updated_at date
);

CREATE TABLE TRN_BOOKS_DETAILS (
	id varchar(50) primary key,
	transaction_id varchar(50) REFERENCES TRN_BOOKS(id),
	books_id varchar(50) REFERENCES MST_BOOKS(id),
	qty int
);
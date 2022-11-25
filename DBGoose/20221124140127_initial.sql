-- +goose Up
create table genres(
    id int auto_increment primary key,
    name varchar(100) not null
);

create table books(
    id int primary key auto_increment,
    name varchar(100) not null,
    genre_id int not null,
    constraint fk_books_genre foreign key (genre_id) references genre(id)
);

create table authors(
    id int auto_increment primary key,
    name varchar(100) not null,
    email varchar(100) not null
);

create table books_authors(
    id int auto_increment primary key,
    author_id int not null,
    book_id int not null,
    constraint fk_author foreign key (author_id) references authors(id),
    constraint fk_book foreign key (book_id) references books(id)
);


-- +goose Down

drop table books_authors;
drop table books;
drop table authors;
drop table genre;



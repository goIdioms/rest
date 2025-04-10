CREATE TABLE users (
    id serial not null unique,
    name varchar(200) not null,
    username varchar(200) not null unique,
    password_hash varchar(200) not null
);

CREATE TABLE todo_lists (
    id serial not null unique,
    title varchar(200) not null,
    discription varchar(200) not null unique
);

CREATE TABLE user_lists (
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items (
    id serial not null unique,
    title varchar(200) not null,
    discription varchar(200) not null unique,
    done boolean not null default false
);

CREATE TABLE lists_item (
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

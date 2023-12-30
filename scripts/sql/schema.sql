-- to run: sqlite3 <datebase.db> < data/scripts/script.sql

drop table if exists service_categories; 
drop table if exists statuses;
drop table if exists services;
drop table if exists clearance;
drop table if exists users;

-- need the following tables
drop table if exists contact_requests;
drop table if exists groups;
drop table if exists addresses;


create table clearance (
    id integer primary key autoincrement,
    clearance_level varchar(60),
    c_description text
);

-- need a table for customers and administrative
create table users (
    id integer primary key autoincrement,
    email varchar(120) unique not null,
    username varchar(30) unique not null,
    password_hash text not null,
    clearance_level integer,
    foreign key (clearance_level) references clearance(id)
);

create table service_categories (
    id integer primary key autoincrement,
    category varchar(100),
    admin_information text,
    public_information text
);

create table statuses (
    id integer primary key autoincrement,
    status_title varchar(75),
    admin_information varchar(250)
);

create table services (
    id integer primary key autoincrement,
    category_id integer,
    status_id integer,
    service_name varchar(100),
    service_description text, 
    selling real,
    foreign key (category_id) references service_categries(id),
    foreign key (status_id) references status(id)
);

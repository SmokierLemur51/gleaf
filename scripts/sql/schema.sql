drop table if exists service_categories; 
drop table if exists statuses;
drop table if exists services;

create table service_categories (
    id integer primary key autoincrement,
    category varchar(100),
    description text
);

create table statuses (
    id integer primary key autoincrement,
    status varchar(75),
    description text
);

create table services (
    id integer primary key autoincrement,
    category_id integer,
    status_id integer,
    service varchar(100),
    description text, 
    selling real,
    foreign key (category_id) references service_categries(id),
    foreign key (status_id) references status(id)
);


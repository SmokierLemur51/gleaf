drop table if exists user;
drop table if exists service_category;
drop table if exists service;
drop table if exists address;
drop table if exists contact;
drop table if exists clean_group;
drop table if exists booking;

create table user (
    id integer primary key autoincrement,
    email text,
    username text,
    hash text,
    contact_id integer,
    address_id integer,
    foreign key (contact_id) references contact(id),
    foreign key (address_id) references address(id)
);


create table service_category (
  id integer primary key autoincrement,
  name text unique not null,
  description text  
);

create table service (
    id integer primary key autoincrement,
    category_id integer, 
    name text,
    description text,
    cost real
);

create table address (
    id integer primary key autoincrement,
    name text,
    street text,
    city text,
    state text,
    zip text
);

create table contact (
    id integer primary key autoincrement,
    user_id integer,
    name text,
    email text, 
    phone_number text,
    foreign key (user_id) references user(id)
);


create table clean_group (
    id integer primary key autoincrement,
    creator_id integer,
    invite_link text,
    name text,
    foreign key (creator_id) references user(id)
);

create table booking (
    id integer primary key autoincrement,
    service_id integer,
    address_id integer,
    contact_id integer,
    user_id integer,
    created_at date,
    requested_for date,
    completed_at date, 
    status_id integer,
    cancelled boolean,
    paid boolean,
    payment_id integer,
    group_booking boolean,
    clean_group_id integer,
    foreign key (service_id) references service(id),
    foreign key (address_id) references address(id),
    foreign key (contact_id) references contact(id),
    foreign key (user_id) references user(id),
    foreign key (status_id) references status(id),
    foreign key (payment_id) references payment(id),
    foreign key (clean_group_id) references clean_group(id)
);


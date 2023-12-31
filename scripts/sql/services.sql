-- this is for testing, not final

-- statuses
insert into statuses (status_title, admin_information) values 
('active', 'We are actively working to book these services.');

insert into statuses (status_title, admin_information) values 
('seasonal', 'We promote these at specific times of the year, outside of those times they can be requested but are not displayed.');

insert into statuses (status_title, admin_information) values 
('deactivated', 'Not currently offering these services.');

insert into statuses (status_title, admin_information) values 
('group-only', 'These services can only be booked by groups, not individuals.');

insert into statuses (status_title, admin_information) values 
('promotion', 'We are trying to promote these at a discount to drive sales.');


-- service categories
insert into service_categories (category, admin_information, public_information) values 
('Moving', 'Moving related services.', 'Our services to help make your move as painless as possible.');

insert into service_categories (category, admin_information, public_information) values 
('Residential', 'Customer house cleaning, no commercial offers. Wide umbrella to define general jobs.', 
"Whether its help catching up after a busy week, toddlers gone wild, a grocery pickup, or a party you don't have time to prepare for, we've got your back!");

insert into service_categories (category, admin_information, public_information) values 
('Residential Exterior', 'Non commercial exterior cleaning.', 
"Windows, siding, decks, gutters, leaf removal, driveways and garages. Almost anything except for junk removal.");

insert into service_categories (category, admin_information, public_information) values 
('Eco-Friendly', "'Green' cleaning solutions, its cheaper and healthier.", 
"Greenleaf Cleanings specialty! Ditch those harmful chemicals to protect your health and the environment!");

insert into service_categories (category, admin_information, public_information) values 
('Group Cleaning', 
"Our attempt to stand out from the crowd. Share booking links to friends if you have an account and recieve discounts based on the number of houses booked.", 
"Create a group and book with your friends, family or neighbors! We offer discounts based off of the size, distance, and history of booking with us.");

insert into service_categories (category, admin_information, public_information) values 
('Commercial', 'Office, warehouse, or moving company bookings. Etc...', 
"Commercial cleaning solutions, contact us to get a free estimate.");

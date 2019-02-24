insert into category(description) values ('Frutas');
insert into category(description) values ('Livros');
insert into category(description) values ('Carros');
insert into category(description) values ('Informática');

insert into products(name, price, quantity, amount, category) values
('Laranja', 1.80, 100, (1.8 * 100), 1);
insert into products(name, price, quantity, amount, category) values
('Banana', 1.10, 80, (1.8 * 80), 1);

insert into products(name, price, quantity, amount, category) values
('Harry Potter', 20, 10, (20 * 10), 2);
insert into products(name, price, quantity, amount, category) values
('Game of Thrones', 40, 20, (40 * 20), 2);

insert into products(name, price, quantity, amount, category) values
('Ferrari', 900000, 5, (900000 * 5), 3);
insert into products(name, price, quantity, amount, category) values
('Jaguar F-Type', 600000, 2, (600000 * 2), 3);

insert into products(name, price, quantity, amount, category) values
('Computador', 3000, 12, (3000 * 12), 4);
insert into products(name, price, quantity, amount, category) values
('Notebook', 2500, 8, (2500 * 8), 4);


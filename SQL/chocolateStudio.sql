---------------------------Строки для СОЗДАНИЯ ТАБЛИЦ--------------------------------
create table Photography (
  ID serial primary key,
  score int null,
  title varchar null,
  photography_text text null,
  price decimal null,
  price_text text null
);

create table PhotoBook (
  ID serial primary key,
  score int null,
  title varchar null,
  photoBook_text text null,
  price decimal null,
  price_text text null
);

create table Photo_image (
  ID serial,
  photo_id int references Photography(id) on delete cascade,
  image_url text UNIQUE null,
  is_image_main boolean
);

create table Book_image (
  ID serial,
  book_id int references PhotoBook(id) on delete cascade,
  image_url text UNIQUE null,
  is_image_main boolean
);






---------------------------Тестовые строки ЗАПРОСОВ (их выполнять не обязательно, они для backend'ера)--------------------------------


---------------------------Добавление фотостудий------------------------------------
insert into Photography(score, title, photography_text, price, price_text) values(1, 'Фотостудия для фотканья детей', 'блаблаблаблабла фоткаем детей', 34.00, 'Фоткаем детей');
insert into Photography(score, title, photography_text, price, price_text) values(2, 'Фотостудия для фотканья взрослых', 'блаблаблаблабла фоткаем взрослых', 142.00, 'Фоткаем взрослых');
insert into Photography(score, title, photography_text, price, price_text) values(3, 'Фотостудия для фотканья инвалидов', 'блаблаблаблабла фоткаем инвалидов', 14.00, 'Фоткаем инвалидов');
insert into Photography(score, title, photography_text, price, price_text) values(4, 'Фотостудия для фотканья студентов', 'блаблаблаблабла фоткаем студентов', 80.40, 'Фоткаем студентов');

---------------------------Добавление фотокниг------------------------------------
insert into PhotoBook(score, title, photobook_text, price, price_text) values(1, 'Фотокнига для фотканья детей', 'блаблаблаблабла книга детей', 34.00, 'Книга детей');
insert into PhotoBook(score, title, photobook_text, price, price_text) values(2, 'Фотокнига для фотканья взрослых', 'блаблаблаблабла книга взрослых', 142.00, 'Книга взрослых');
insert into PhotoBook(score, title, photobook_text, price, price_text) values(3, 'Фотокнига для фотканья инвалидов', 'блаблаблаблабла книга инвалидов', 14.00, 'Книга инвалидов');
insert into PhotoBook(score, title, photobook_text, price, price_text) values(4, 'Фотокнига для фотканья студентов', 'блаблаблаблабла книга студентов', 80.40, 'Книга студентов');

---------------------------Добавление картинок для фотостудии------------------------------------
insert into Photo_image(photo_id, image_url, is_image_main) values(1, '/media/photo1.jpg', true);
insert into Photo_image(photo_id, image_url, is_image_main) values(1, '/media/photo2.jpg', false);

insert into Photo_image(photo_id, image_url, is_image_main) values(2, '/media/photo3.jpg', true);
insert into Photo_image(photo_id, image_url, is_image_main) values(2, '/media/photo4.jpg', false);

insert into Photo_image(photo_id, image_url, is_image_main) values(3, '/media/photo5.jpg', true);
insert into Photo_image(photo_id, image_url, is_image_main) values(3, '/media/photo6.jpg', false);

insert into Photo_image(photo_id, image_url, is_image_main) values(4, '/media/photo7.jpg', true);
insert into Photo_image(photo_id, image_url, is_image_main) values(4, '/media/photo8.jpg', false);

---------------------------Добавление картинок для фотокниг------------------------------------
insert into Book_image(book_id, image_url, is_image_main) values(1, '/media/book1.jpg', true);
insert into Book_image(book_id, image_url, is_image_main) values(1, '/media/book2.jpg', false);

insert into Book_image(book_id, image_url, is_image_main) values(2, '/media/book3.jpg', true);
insert into Book_image(book_id, image_url, is_image_main) values(2, '/media/book4.jpg', false);

insert into Book_image(book_id, image_url, is_image_main) values(3, '/media/book5.jpg', true);
insert into Book_image(book_id, image_url, is_image_main) values(3, '/media/book6.jpg', false);

insert into Book_image(book_id, image_url, is_image_main) values(4, '/media/book7.jpg', true);
insert into Book_image(book_id, image_url, is_image_main) values(4, '/media/book8.jpg', false);


---------------------Тестовые запросы соединения таблиц и выборки данных--------------

-------------------------------------Запрос на Фотостудии----------------------------------


SELECT
    *
FROM Photography
WHERE id=1;



--------------------------Необходимо важные тестовые запросы--------------------------

select * from Photography;

----------------------------------------------Удаление таблицы-------------------------------------------------------
drop table Photography cascade;
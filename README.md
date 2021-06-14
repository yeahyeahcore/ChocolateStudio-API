## Для запуска сервера необходимо:

1. Запустить `Docker`<br />
2. Выполнить команду:

`docker run --name [имя контейнера] -e POSTGRES_PASSWORD=[пароль БД] -e POSTGRES_USER=[имя пользователя] -e POSTGRES_DB=[имя БД] -p 5432:5432 -d postgres`

3. SQL строки для создания таблиц находятся в папке `SQL`

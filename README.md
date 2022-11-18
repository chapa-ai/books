**Description**<br>
Just a small book service implemented on the GRPC approach<br>

<hr>


**Installation**<br>

```bash
git clone https://github.com/chapa-ai/books.git

cd build

make build

make client

```

<hr>

Тестовое задание на должность Golang-разработчик<br>
Цель задания: Проверить у кандидата его навыки и знания.<br>
Задание: Спроектировать базу данных, в которой содержится авторы
книг и сами книги.<br> Необходимо написать сервис который будет по
автору искать книги, а по книге искать её авторов.<br>
    Требования к сервису:<br>
1. Сервис должен принимать запрос по GRPC.<br>
2. Должна быть использована база данных MySQL<br>
3. Код сервиса должен быть хорошо откомментирован<br>
4. Код должен быть покрыт unit тестами<br>
5. В сервисе должен лежать Dockerfile, для запуска базы данных с
тестовыми данными<br>
6. Должна быть написана документация, как запустить сервис.
Плюсом будет если в документации будут указания на команды,
для запуска сервиса и его окружения, через Makefile<br>
7. код должен быть выложен на github<br>





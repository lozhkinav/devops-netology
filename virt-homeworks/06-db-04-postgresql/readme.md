# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

Подключитесь к БД PostgreSQL используя `psql`.

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
- подключения к БД
- вывода списка таблиц
- вывода описания содержимого таблиц
- выхода из psql

### Ответ:
```
Комманды : 
21:37:04 alex@upc(0):~$ docker pull postgres:13
21:34:58 alex@upc(0):~$ docker volume create vol_postgres
21:40:43 alex@upc(0):~$ docker run --rm --name pg-docker -e POSTGRES_PASSWORD=postgres -ti -p 5432:5432 -v vol_postgres:/var/lib/postgresql/data postgres:13
21:45:27 alex@upc(0):~$ docker exec -it pg-docker bash
root@8d48e9af5c27:/# psql -h localhost -p 5432 -U postgres -W
```
- вывода списка БД
```
postgres-# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)
```
- подключения к БД - \c[onnect] {[DBNAME|- USER|- HOST|- PORT|-] | conninfo}
```
postgres-# \c postgres
Password: 
вывода списка таблиц postgres-# \dt в таблицах пусто , исползовал оп параметр S - для системных объектов
postgres-# \dtS
                    List of relations
   Schema   |          Name           | Type  |  Owner   
------------+-------------------------+-------+----------
 pg_catalog | pg_aggregate            | table | postgres
 pg_catalog | pg_am                   | table | postgres
 pg_catalog | pg_amop                 | table | postgres
...
```
- вывода описания содержимого таблиц
```
postgres-# \d[S+] NAME
postgres-# \dS+ pg_index
                                      Table "pg_catalog.pg_index"
     Column     |     Type     | Collation | Nullable | Default | Storage  | Stats target | Description 
----------------+--------------+-----------+----------+---------+----------+--------------+-------------
 indexrelid     | oid          |           | not null |         | plain    |              | 
 indrelid       | oid          |           | not null |         | plain    |              | 
 indnatts       | smallint     |           | not null |         | plain    |              | 
 indnkeyatts    | smallint     |           | not null |         | plain    |              | 
 indisunique    | boolean      |           | not null |         | plain    |              | 
 indisprimary   | boolean      |           | not null |         | plain    |              | 
 indisexclusion | boolean      |           | not null |         | plain    |              | 
 indimmediate   | boolean      |           | not null |         | plain    |              | 
 indisclustered | boolean      |           | not null |         | plain    |              | 
 indisvalid     | boolean      |           | not null |         | plain    |              | 
 indcheckxmin   | boolean      |           | not null |         | plain    |              | 
 indisready     | boolean      |           | not null |         | plain    |              | 
 indislive      | boolean      |           | not null |         | plain    |              | 
.....
```
- выхода из psql
```
postgres-# \q
root@8d48e9af5c27:/# 
```

## Задача 2

Используя `psql` создайте БД `test_database`.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.

Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders` 
с наибольшим средним значением размера элементов в байтах.

**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

### Ответ:
```
postgres=# CREATE DATABASE test_database;
CREATE DATABASE

root@8d48e9af5c27:/var/lib/postgresql/data# psql -U postgres -f ./pg_backup.sql test_database

postgres=# \c test_database
Password: 
You are now connected to database "test_database" as user "postgres".
test_database=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner   
--------+--------+-------+----------
 public | orders | table | postgres
(1 row)

test_database=# ANALYZE VERBOSE public.orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE

test_database=# select avg_width from pg_stats where tablename='orders';
 avg_width 
-----------
         4
        16
         4
(3 rows)
```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

### Ответ:
- преобразовать существующую таблицу в партиционированную поэтому пересоздадим таблицу
```
test_database=# alter table orders rename to orders_simple;
ALTER TABLE
test_database=# create table orders (id integer, title varchar(80), price integer) partition by range(price);
CREATE TABLE
test_database=# create table orders_less499 partition of orders for values from (0) to (499);
CREATE TABLE
test_database=# create table orders_more499 partition of orders for values from (499) to (999999999);
CREATE TABLE
test_database=# insert into orders (id, title, price) select * from orders_simple;
INSERT 0 8
test_database=# 
```
- При изначальном проектировании таблиц можно было сделать ее секционированной, тогда не пришлось бы переименовывать исходную таблицу и переносить данные в новую.

## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

### Ответ:
```
root@8d48e9af5c27:/var/lib/postgresql/data# pg_dump -U postgres -d test_database >test_database_dump.sql
```
- Для уникальности можно добавить  первичный ключ.
```
ALTER TABLE orders

  ADD CONSTRAINT [ title_key ]

    PRIMARY KEY (ititle)

```

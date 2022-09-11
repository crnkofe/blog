# Go paging demo

This project demonstrates DB cursor paging in Golang. 

## Quickstart

```bash
make
make docker-up
# docker compose up will take a few moments to pull containers, run MySQL and run a migration
# afterwards just run the binary
./paging
```

## Prerequisites

* [Go 1.19](https://go.dev/dl/)
* make (or just run commands from Makefile manually)
* [Docker 2 (uses new `docker compose` instead of `docker-compose`)](https://github.com/docker/compose)
* (optional) [golangci-lint](https://golangci-lint.run/)

## Connect to MySQL in Docker

Execute the following and when asked for password input password.
```bash
docker exec -it paging-db-1 mysql -u root -p paging
```

Once in MySQL prompt selecting from `computer` table show yield some data.
```
mysql> SELECT * FROM computer LIMIT 5;
+----+----------------------------------------+
| id | name                                   |
+----+----------------------------------------+
|  1 | Commodore 64                           |
|  2 | Altair 8800                            |
|  3 | Apple I and also Replica 1             |
|  4 | Applix 1616                            |
|  5 | Compukit UK101                         |
```
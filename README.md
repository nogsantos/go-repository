# Poc Go

Pattern repository an example in Go.

## Setup

**dependencies**

```shell
go get
```

**Sqlite**

Manually create the database file and the table.

```shell
// Create the database file
touch repository.sqlite

// Access the database
sqlite3 repository.sqlite

// Inside the database, create the table
create table categories(id varchar(255), name varchar(255));

// Exit
Ctrl + d
```

**Run**

```shell
go run main.go
```

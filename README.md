# SQL DATABASES

1. [What is a DB](#what-is-a-database)
2. [Database management system](#database-management-sistem-dmbs)
3. [CRUD](#crud)
4. [Types of database](#types-of-database)
    - 4.1 [Relational Databases](#relational-databases-sql)
    - 4.2 [Non-relational Databases](#non-relational-databases-nosql)
5. [Database query](#database-query)
6. [Tables and keys](#tables-and-keys)

## What is a database?

* Database is any collection of relative information

| Examples      | Ways to store |
|---------------|---------------|
| Phone book    | On paper      |
| Shooping list | On computer   |
| TODO list     | In mind       |

## Database management sistem (DMBS)

* A special software program that helps users create and mantain a database
    - Makes it easy to manage a large amount of information
    - Handles security
    - Backups
    - Importing/exporting data
    - Concurrency
    - interacts with software application 

## C.R.U.D.
* There are 4 main operations with databases:
    - CREATE
    - READ
    - UPDATE
    - DELETE

## Types of database

| Relational Databases      | Non-relational databases |
|---------------|---------------|
| Organize data into tables    | Does not orgonize data in traditional tables |
| each table has columns and rows | Key-value stores |
| A unique key identifies each row| Documents, graphs |

### Relational Databases (SQL)

* Examples:

| Student table | Users tables |
|---------------|---------------|
| ID NAME Major    | Username Password email |
| 1 Jack biology   | jsmith22 wordpass ... |
| 2 Kate Sociology | catlover44 ilovecats111 ... |
| 3 Ivan Math      | ... ... ... |

* SQL databases: (mySQL, PostgreSQL, Oracle, etc)

* Structured Query Language:
    - Standardized language for interaction with RDBMS
    - Performs CRUD operations and other
    - Used to define tables and structures

### Non-relational Databases (noSQL)

* Examples:

* Documents: JSON, XML, etc.
* GRAPHS: Relational nodes.
* Key-value: strings, maps, etc.

* NRDBMS
    - Helps users create and maintain NRDB
        - (mongoDB, dynamoBD, firebase, Redis, etc.)

* Implementation specific:
    - No standard language for NRDBMS
    - Own language for performing CRUD and administrative operations

## Database query

- Queries are request made to DBMS for information

## Tables and keys

- Unique ID's are used to specify a row and called `primary key`
- `foreign key` refers to primary key in another table

![fkey](./images/foreign_key.png)
![branch-table](./images/branch-table.png)
* in the employee table we have branch_id field which is primary key for branch table

- `composite key` is a key that needs two attributes

## MySQL

* [download mysql macOS/windows](https://dev.mysql.com/downloads/mysql/)

* MySQL is a DBMS

## SQL

### Data Types

    - `INT`             // Whole number
    - `DECIMAL(M,N)`    // Decimal number - Exact value M - digits before decimal point, N - digits after decimal point
    - `VARCHAR(len)`    // String of text of lenght len
    - `BLOB`            // Binary large object(images, files)
    - `DATE`            // YYYY-MM-DD
    - `TIMESTAMP`       // YYYY-MM-DD HH:MM:SS - used for recording


### Commands

- Create a database: 
```sql
create database <dbname>;
```

- Create table:
```sql
CREATE TABLE <name>;
```

- Delete table: 
```sql
DROP TABLE <name>;
```

- Get table description
```sql
DESCRIBE <tablename>;
```

[!] **Every** sql query ends with a **semicolon** `;`

### DB Normalization
- Database normalization provides for reduction to normal form
- Normal form 1:
    - no duplicated rows
    - all atributes are atomic
    - no dublicated atributes with common mean
- Normal form 2:
    - Conditions of NF1 are met
    - Have a primary key
    - all non-key attributes depend on **whole** primary key
- Normal form 3:
    Conditions of NF2 are met
    - Non-key attrubutes are depend **only** on primary key and not depend on other atrubutes.

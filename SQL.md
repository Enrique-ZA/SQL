# SQL
## Introduction:
### - Statement: any SQL that takes an action.
```sql
SELECT * FROM table_name;
```
### - SQL is white-space independent.
### - Statements are made up of clauses.
```sql
SELECT *            -- clause
FROM table_name;    -- clause
```
### - Clauses are defined (or counted) by the number of keywords in a statement.
### - Statement is comprised of clauses which contain keywords and field names, 
###   predicates include a value or condition called an expression.
### - There are also operators which lets one check equality among other things.
### - Statements end with a semicolon.
### - Query: any statement that returns records.
### - Using SQL to work with data in existing tables is called 
###   Data Manipulation Language (DML).
### - DML:
###        - Edit data in a database.
###        - CRUD: Create, read, update or delete records.
### - DDL (data definition language):
###        - Edit the structure (schema) of a database.
###        - Add, change, or remove fields or tables.
## Asking for data from a database:
### Ask for data with ***SELECT***
### - ***SELECT*** keyword tells the database we want some information returned.
### - SELECT can also return data not stored in the database.
```sql
SELECT 'Hello, World';
```
```sql
SELECT first_name FROM people;
```
```sql
SELECT first_name, last_name FROM people;
```
### * represents all
```sql
SELECT * FROM people;
```
### Narrowing down a query using ***WHERE***: 
### - The WHERE keyword lets one add selection criteria to a statement.
```sql
SELECT * FROM people WHERE state_code='CA';
```
### - SQL may be case sensitive depending on how the database was setup.
```sql
SELECT first_name, last_name, shirt_or_hat
FROM people 
WHERE state_code='CA';
```
### - Order is important so SELECT then FROM then WHERE
```sql
SELECT first_name, last_name
FROM people 
WHERE state_code='CA' AND shirt_or_hat='hat'
```
### - AND is a logical operator
### - Logical conditions can be chained
```sql
SELECT first_name, last_name 
FROM people 
WHERE state_code='CA' AND shirt_or_hat='shirt' AND team='Angry Ants';
```
### - For not equal we add a ! - see _team_ below.
```sql
SELECT first_name, last_name 
FROM people 
WHERE state_code='CA' AND shirt_or_hat='shirt' AND team!='Angry Ants';
```
### - The keyword _IS_ or _IS NOT_ can also be used for equality.
```sql
SELECT first_name, last_name 
FROM people 
WHERE state_code='CA' AND shirt_or_hat='shirt' AND team IS 'Angry Ants';
```
### - The logical operator _<>_ can also be used for not
```sql
SELECT first_name, last_name 
FROM people 
WHERE state_code='CA' AND shirt_or_hat='shirt' AND team <> 'Angry Ants';
```
### - The logical operator _OR_ can also be used
```sql
SELECT shirt_or_hat, state_code, first_name, last_name 
FROM people 
WHERE (state_code='CA' OR state_code='CO') AND shirt_or_hat='shirt';
```
### - Parentheses are important for logicl operations 

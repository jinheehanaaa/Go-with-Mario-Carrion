# Transaction with pgx
- We will be using pgx instead of database sql
- You can still use Begin,Commit, Rollback

# What's new with pgx without stdlib?
- It defines 2 methods that wrap the actual transaction in 1 function.
- You don't have to deal explicitly with need of calling Commit, or Rollback.
- - conn can use BeginFunc & BeginTx, BeginTxFunc for transaction Begin

# Schema
```sql
CREATE TABLE users(
  name       VARCHAR  NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);
```

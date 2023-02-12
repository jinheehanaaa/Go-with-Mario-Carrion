# pgx with stdlib vs. pgx without stdlib

## stdlib
- Begin, Commit, Rollback

## pgx
- Access Begin, Commit, Rollback with Connect method
- Code is more cleaner

# SQL Schema
```sql
CREATE TABLE users(
  name       VARCHAR  NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);
```

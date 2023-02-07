# Transaction with stdlib
## SQL Only
- BEGIN (INSERT, UPDATE, INSERT)
- COMMIT
- ROLLBACK (If fail)

## SQL with Go
- Focus on inserUsers(){}
- - db.BeginTx
- - tx.Commit
- - tx.Rollback

# Learn
- If we don't use Transactions...
- we will still be inserting somethig into DB even if 1 of whole process is failing


# Schema
```sql
CREATE TABLE users(
  name       VARCHAR  NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);
```

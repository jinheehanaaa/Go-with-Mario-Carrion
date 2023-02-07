# pgx with stdlib vs. pgx without stdlib
## stdlib
- Begin
- Commit
- Rollback

## pgx
- We can access all Begin, Commit, Rollback in Connect Method
- Code is more cleaner

# SQL Schema
```sql
CREATE TABLE users(
  name       VARCHAR  NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);
```


# Resources
- [YouTube](https://www.youtube.com/watch?v=BLr2V6zB5k4&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=16)
- [Example Code](https://github.com/MarioCarrion/videos/tree/1fae8d5c97230c4c2d555dea6ade37b3251a5f2b/2021/11/26)
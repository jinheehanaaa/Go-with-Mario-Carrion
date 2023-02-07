# ???

# Step 1: Create Docker Container
## Run Docker with mysql
```docker
docker run \
  --rm \
  -d \
  -e MYSQL_PASSWORD=password \
  -e MYSQL_DATABASE=dbname \
  -p 3306:3306 \
  mysql:8.0.32
```
## Access DB with password & dbname
- docker logs -f b25ba7c0a82158f1837b2701c8c78d5e425609f70cc82064c59e7181d3b3de3f
- docker exec -it b25ba7c0a82158f1837b2701c8c78d5e425609f70cc82064c59e7181d3b3de3f mysql -u root -p -h localhost
- SHOW DATABASES;

# Step 2: Create Table

```sql
CREATE TABLE users(
  name       VARCHAR(10) NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```
- SELECT * FROM users;
- INSERT INTO users(name, birth_year) VALUES("mario", 1900);




# Resources
- [YouTube](https://www.youtube.com/watch?v=gT2ztW4j66A&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=15)
- [Example Code](https://github.com/MarioCarrion/videos/tree/0cb79c49f39e62b265e4e545a30169dd2e2d9d8f/2021/12/09)
- [go sql driver](https://github.com/go-sql-driver)
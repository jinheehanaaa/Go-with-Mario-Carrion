
# CMD
- <code>go get github.com/jackc/pgx</code>

# DOCKER
docker run \
-d \
-e POSTGRES_HOST_AUTH_METHOD=trust \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=dbname \
-p 5432:5432 \
postgres:15-alpine

# Docker CMD
Check Docker ID: <code>docker ps</code>
Access DB: <code>docker exec -it 913a5548e6f3 psql -U user -d dbname</code>

# Data Source Name Implementation

# Create DB Table
CREATE TABLE users(
name VARCHAR NOT NULL,
birth_year SMALLINT NULL DEFAULT 0);

# Insert value
INSERT INTO users(name, birth_year) VALUES('jinhee', 1900);

# Check
SELECT * FROM users;

# Access DB value using query method from Go

# 2. Insert into DB
- INSERT INTO users(name, birth_year) VALUES('jinhee', 1901)

# 3 - Placeholder instead of string manipulation
- INSERT INTO users(name, birth_year) VALUES($1, $2)

# Resources
- [YouTube](https://www.youtube.com/watch?v=2XCaKYH0Ydo&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=14)
- [Code Example](https://github.com/MarioCarrion/videos/tree/0e574a88b82ea18400ca525e01afd98d9f014e18/2021/11/17)
- [pgx package](https://github.com/jackc/pgx)
# Objective
- Insert values in csv into database
- Compare [stdlib] vs. [pgx's Copy Method] performance

# Docker
```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:13.5-alpine3.15
```

docker exec -it 7e1801637b677437f3341da9b96ec787da67d9cc2fc03161b448fca46856c2f9 psql -U user -d dbname

# Schema
```sql
CREATE TABLE users(
  first_name VARCHAR  NOT NULL,
  last_name  VARCHAR  NOT NULL,
  age        INT      NOT NULL
);
```

# Compare
- stdlib: 70ms for 50 Records
- pgx: 7 ms for 50 Records

# Conclusion
- pgx is much more faster
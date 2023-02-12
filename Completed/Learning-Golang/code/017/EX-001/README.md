# mysql's functions
- RegisterLocalFile
- RegisterReaderHandler

# 1.Register & Inserting data.csv into sql DB
- <code>RegisterLocalFile</code> for data.csv output
- <code>FIELDS TERMINATED BY ,ENCLOSED BY</code> for formatting

## 1-a: Docker
```
docker run \
  --rm \
  -d \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=dbname \
  -p 3306:3306 \
  mysql:8.0.27
```
- <code>docker -f logs 0d7f9f4b0e46dcb76f2a77ecd504a9efa12c7e7a0f79e765d80544a307e8e6d7</code>
- <code>docker exec -it 0d7f9f4b0e46dcb76f2a77ecd504a9efa12c7e7a0f79e765d80544a307e8e6d7 mysql -u root -p -h localhost</code>

## 1-b: Create SQL DB Table
- <code>USE dbname;</code>
```sql
CREATE TABLE users(
  first_name VARCHAR(10) NOT NULL,
  last_name  VARCHAR(10) NOT NULL
);
```
- SELECT * FROM users;
- Show local_infile: <code>SHOW GLOBAL VARIABLES LIKE 'local-infile'; </code>
- Enable local data: <code>SET GLOBAL local_infile=1;</code>

- Run our code: <code>go run main.go</code>

- See DB Table: <code>SELECT * FROM users;</code>
- Delete users: <code>DELETE FROM users;</code>

# 2. Modifying the data
## RegisterReaderHandler()
- - Open file => Read record from file => Modify record
- - Make sure you implement defer close()
## NewWriter()
- - We modifiy record with NewWriter (buffer as return value)
## Flush()
- - Save buffer into writer
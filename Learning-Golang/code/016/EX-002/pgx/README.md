# Objective
- Insert values in csv into database
- Compare [stdlib] vs. [pgx's Copy Method] performance

# pgx
## conn.CopyFrom function
- - Indicates how many rows were inserted into DB (count)
- - Receives CopyFromSource to read the data
## CopyFromSource structure
- This replaces BeginTxFunc by creating CopyFromSource struct
- 3 Methods (Next, Values, Err) is internally used by CopyFrom()

# Test
- 7 ms for 50 Records

# SQL
- SELECT Count(*) FROM users;

# Interface value
- A tuple of a value and a concrete type
- An Interface value is nil if:
- - both its value and concrete type are nil



# Interface{}
- Specifies zero methods
- May hold any values of any type
- Used by code that handles unknown types

## Examples of interface{}
- Functions in fmt
- Functions in encoding/json
- Types in database/sql

## Using interface{}
- Type Assertions
- Type Switches




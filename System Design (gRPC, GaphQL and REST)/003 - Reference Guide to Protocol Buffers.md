# Configuration
- pakcage name, syntax use, options
# Data
## messages
- declared using fields
- Type, Field, Number + Extra options

## Options
- Top level, Message level, Field level
- Specify "optional" for not requilred field
- - EX: <code>option deprecated = true;</code>
- - EX: <code>int64 birth_year = 3[deprecated = true];</code>
- - But type still exists
- - For backward compatibility, you can use reserved
- - EX: <code>reserverd 4;</code>

## enumerations
- Predefined list of values
- - EX: MaritalStatus with 0, 1, 2

# Tip
## messages
-  => not required field


## CMD
buf lint
buf generate
ruby reader/main.rb


# Resources
- [YouTube](https://www.youtube.com/watch?v=LNtmrqkooLo&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=3)

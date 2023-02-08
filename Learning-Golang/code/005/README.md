# Example of context package
- database/sql
- - database/sql.*DB.QueryContext
- - database/sql.*DB.ExecContext
- net/http
- - net/http.NewRequestWithContext
- OpenTelemetry

# What is in context
## Dealines
### WithDeadline: End at specific time
- rcv: parent context.Context
- rcv: d time.Time
- return: context.Context
- return: context.CancelFunc

### WithTimeout: End after some amount of time  (EX-001)
- rcv: parent context.Context
- rcv: timeout time.Duration
- return: context.Context
- return: context.CancelFunc

## Cancellation Signals
### WithCancel: Call to trigger the cancellation (EX-002)
- rcv: parent context.Context
- return: ctx context.Context
- return: cancel context.CancelFunc

## Request-scoped values (EX-003)
### WithValue: Pass value for different request
- rcv: parent context.Context
- rcv: key, val interface{}
- return: context.Context

# Best practices
- Define a ctx context.Context as first argument
- Make sure context.CancelFunc is called (otherwise, memory can leak!)
- Don't overuse context.*Context.WithValue

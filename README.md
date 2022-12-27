# Splitwise
Debt simplification algorithm written in Python, Go and Rust

# Getting Started
## Generate test data
```
./create-testdata.sh
```
## Golang
```
./go.sh
```
## Python
```
./python.sh
```
## Rust
```
./rust.sh
```


## Performance comparison

Thorough benchmarking has not been conducted, but sample algorithm execution times are printed to stdout by the main program.


### 20 debtors, 100 transactions

Python
```
read csv file in 3423.845 microseconds
completed simplify_debts execution in 428.073 microseconds, transactions 17
```

Golang
```
completed csv load execution in 90.245µs
completed SimplifyDebts execution in 70.173µs, settlement transactions 17
```

Rust
```
completed csv load execution in 75.81µs
completed simplify_debts execution in 16.79µs, settlement transactions 17
```

### 20 debtors, 1000 transactions

Python
```
read csv file in 5711.69 microseconds
completed simplify_debts execution in 1748.723 microseconds, transactions 18
```

Golang
```
completed csv load execution in 440.076µs
completed SimplifyDebts execution in 155.269µs, settlement transactions 18
```

Rust
```
completed csv load execution in 301.86µs
completed simplify_debts execution in 82.17µs, settlement transactions 18
```
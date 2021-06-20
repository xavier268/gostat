# gostat
Golang librairy create easy statistics on the console.

## Design objectives 

* minimal setting, reasonable defaults
* accept any type of scalar data : int, float ...
* all operations in bounded time
* bounded memory footprint should be bounded
* provide exact mean, variance, min, max, count
* provide reasonable estimates for mode, quartile, median
* display histogram of values on console, ithout 
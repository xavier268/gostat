# gostat
Golang librairy create easy statistics on the console.

## Design objectives 

* minimal setting, reasonable defaults
* accept any type of scalar data : int, float ...
* memory foot print and all operations should be lienar in time wrt the total number of data input
* provide exact mean, variance, min, max, count
* provide reasonable estimates for mode, quartile, median
* display histogram of values on console, ithout 
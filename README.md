# gostat
A Golang librairy to easily extract statistical information from a stream of scalar data.

## Design objectives 


* minimal setting, reasonable defaults                      ->  OK 
* accept any type of scalar data : int, float               ->  OK 
* all operations in bounded time                            ->  OK 
* bounded memory footprint                                  ->  OK 
* provide exact mean, variance, min, max, count             ->  OK
* provide reasonable estimates for mode, quartile, median   ->  TO DO (first, work on a "percentile" function)
* display histogram of values on console                    ->  TO DO 
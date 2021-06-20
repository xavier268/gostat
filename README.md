# gostat
A Golang librairy to easily extract statistical information from a stream of scalar data.

## Design objectives 


* minimal setting, reasonable defaults                      ->  OK 
* accept any type of scalar data : int, float               ->  OK 
* all operations in bounded time                            ->  OK 
* bounded memory footprint should be bounded                ->  OK 
* provide exact mean, variance, min, max, count             ->  TO DO 
* provide reasonable estimates for mode, quartile, median   ->  TO DO 
* display histogram of values on console                    ->  TO DO 
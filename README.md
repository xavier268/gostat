# gostat
A Golang librairy to easily extract statistical information from a stream of scalar data.

## Design objectives 

* minimal setting, reasonable defaults                      ->  OK 
* accept any type of scalar data : int, float               ->  OK 
* all operations in bounded time                            ->  OK 
* **bounded memory footprint**                              ->  OK 
* provide exact mean, variance, min, max, count             ->  OK
* display histogram of values on console                    ->  OK

## Typical usage 

    s := NewStat(200)       // create a stat object, precision is 200 buckets
    s.Add(2.3)
    s.Add(456.002)          // Add float
    s.Add(34)               // Add integer
    .../...                 // add further data points ...
    h := s.ToHisto(40)      // produce an histogram with 40 segments ( note that "segments" is "small" compared to "presision")
    fmt.Println(h)          // Display histogram on console
    .../...                 // add futher data and display a new histogram as needed

## Further tasks

* provide reasonable estimates for mode, quartile, median   (first, work on a "percentile" function)
* improve numerical stability when working with large number of (thin) buckets ?

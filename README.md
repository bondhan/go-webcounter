# Go-WebCounter

A web counter 'trying' to handle 'million' requests (on going)

## Design

Request <--> Buffered Redis Queue  --> GoRoutine Consumer --> PostreSQL Master-Slave

### Run

Run in a terminal:
```bash
docker-compose up --verbose
``` 

Run in another terminal:
```bash
go build  -o bin/web_counter main.go  && ./bin/web_counter
```

### Benchmark
System: i7 8650 4 cores 8 threads 32 GB RAM 512 GB SSD

10000 total requests with concurrent 100 request at a time:
```bash
ab -n 10000 -c 100 http://localhost:8080/web-counter/increment
```

Results: (** Requests per second:    1633.76 [#/sec] (mean) **)

```bash
bondhan@syuhada:~/workspace/go/go-webcounter$ ab -n 10000 -c 100 http://localhost:8080/web-counter/increment
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /web-counter/increment
Document Length:        1 bytes

Concurrency Level:      100
Time taken for tests:   6.121 seconds
Complete requests:      10000
Failed requests:        9991
   (Connect: 0, Receive: 0, Length: 9991, Exceptions: 0)
Total transferred:      1258894 bytes
HTML transferred:       38894 bytes
Requests per second:    1633.76 [#/sec] (mean)
Time per request:       61.208 [ms] (mean)
Time per request:       0.612 [ms] (mean, across all concurrent requests)
Transfer rate:          200.85 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      0      12
Processing:     1   60  26.5     57     165
Waiting:        1   60  26.5     56     164
Total:          1   61  26.3     57     165
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     57
  66%     73
  75%     81
  80%     84
  90%     98
  95%    106
  98%    118
  99%    126
 100%    165 (longest request)

```

In database 10,001 counter is inserted without failure (+1 because starts from 0)
```bash
select * from m_visitor mv order by id desc;
select count(*) from m_visitor mv;
```
# Go-WebCounter

A try web counter handling million requests (on going)

## Design

Request <--> Buffered Redis Queue  --> 4 GoRoutine Consumer --> PostreSQL Master-Slave

### Run

Run in a terminal:
```bash
docker-compose up --verbose
``` 

Run in a another terminal:
```bash
go build  -o bin/web_counter main.go  && ./bin/web_counter
```

### Benchmark
System: i7 8650 4 cores 8 threads 32 GB RAM 512 GB SSD

10000 total requests with concurrent 1000 request at a time:
```bash
ab -n 10000 -c 1000 http://localhost:8080/web-counter/increment
```

Results: (**Requests per second:    841.67 [#/sec] (mean)**)

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
Time taken for tests:   11.881 seconds
Complete requests:      10000
Failed requests:        9991
   (Connect: 0, Receive: 0, Length: 9991, Exceptions: 0)
Total transferred:      1258894 bytes
HTML transferred:       38894 bytes
Requests per second:    841.67 [#/sec] (mean)
Time per request:       118.811 [ms] (mean)
Time per request:       1.188 [ms] (mean, across all concurrent requests)
Transfer rate:          103.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.0      0      14
Processing:     1  117 104.3     68     586
Waiting:        1  116 104.3     68     586
Total:          1  117 104.1     69     587

Percentage of the requests served within a certain time (ms)
  50%     69
  66%    105
  75%    142
  80%    164
  90%    262
  95%    363
  98%    477
  99%    505
 100%    587 (longest request)
```

In database 10,000 counter is written without failure.
```bash
select * from m_visitor mv order by id desc;
select count(*) from m_visitor mv;
```
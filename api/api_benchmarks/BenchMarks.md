## API benchmarks:
Following are the sample benchmarks of the qlikapp apis

1) A qlikapp hello message GET api at `/hello`:
```
go-wrk -d 5 http://localhost:8080/v1/hello -H "accept: application/json" 
Running 5s test @ http://localhost:8080/v1/hello
  10 goroutine(s) running concurrently
50552 requests in 4.960705256s, 5.35MB read
Requests/sec:		10190.49
Transfer/sec:		1.08MB
Avg Req Time:		981.307µs
Fastest Request:	169.498µs
Slowest Request:	36.991819ms
Number of Errors:	0

```

2) Fetch a list of messages in response to a valid GET request at `/messages`: 
```
go-wrk -d 5 http://localhost:8080/v1/messages -H "accept: application/json"
Running 5s test @ http://localhost:8080/v1/messages
  10 goroutine(s) running concurrently
12873 requests in 4.987729042s, 3.03MB read
Requests/sec:		2580.93
Transfer/sec:		622.55KB
Avg Req Time:		3.874566ms
Fastest Request:	617.106µs
Slowest Request:	63.72256ms
Number of Errors:	0

```

3) Create a new message in response to a valid POST request at `/messages`:
```
go-wrk -d 5 http://localhost:8080/v1/messages -H "accept: application/json" -M POST -body {\r\n \"message\" : \"this is a benchmark message\" \r\n}
Running 5s test @ http://localhost:8080/v1/messages
  10 goroutine(s) running concurrently
12556 requests in 4.987905204s, 2.96MB read
Requests/sec:		2517.29
Transfer/sec:		607.20KB
Avg Req Time:		3.972527ms
Fastest Request:	603.038µs
Slowest Request:	75.022784ms
Number of Errors:	0

```

4) Check if a given message is a palindrome in response to a valid GET request at `/messages/{id}`:
```
go-wrk -d 5 http://localhost:8080/v1/messages/palindromeChk/3 -H "accept: application/json"
Running 5s test @ http://localhost:8080/v1/messages/palindromeChk/3
  10 goroutine(s) running concurrently
12518 requests in 4.986341712s, 1.74MB read
Requests/sec:		2510.46
Transfer/sec:		357.94KB
Avg Req Time:		3.983337ms
Fastest Request:	600.105µs
Slowest Request:	25.109579ms
Number of Errors:	0

```
5) Fetch a message in response to a valid GET request at `/messages/{id}`:
```
go-wrk -d 5 http://localhost:8080/v1/messages/1 -H "accept: application/json"
Running 5s test @ http://localhost:8080/v1/messages/1
  10 goroutine(s) running concurrently
12587 requests in 4.988493067s, 1.56MB read
Requests/sec:		2523.21
Transfer/sec:		320.33KB
Avg Req Time:		3.96321ms
Fastest Request:	663.795µs
Slowest Request:	21.160227ms
Number of Errors:	0

```

6) Update a message in response to a valid PUT request at `/messages/{id}`:
```
go-wrk -d 5 http://localhost:8080/v1/messages/1 -H "accept: application/json" -M PUT -body {\r\n \"message\" : \"this is a benchmark message\" \r\n}
Running 5s test @ http://localhost:8080/v1/messages/1
  10 goroutine(s) running concurrently
12441 requests in 4.989894877s, 1.52MB read
Requests/sec:		2493.24
Transfer/sec:		311.65KB
Avg Req Time:		4.010847ms
Fastest Request:	610.51µs
Slowest Request:	28.747914ms
Number of Errors:	0

```
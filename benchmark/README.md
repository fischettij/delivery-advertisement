# Benchmark results


### With Postman
The results of the tests can be viewed in the web browser.

In both cases the setup is the same:
- Virtual: users 100
- Duration: 5 minutes
- Ramp-up 1: minute

## Result with postgres DB
- P90: 247 ms
- P99: 1.038 ms
- Error rate: 0%

[Report](postman-postgres-performance-report.html)

## Result with in memory storage
- P95: 7ms
- P99: 13ms
- Error rate: 0%

[Report](postman-memorystorage-performance-report.html)

## Result with in memory with many entries
Benchmark result with 5 million of entries in storage
- P95: 20ms
- P99: 186ms
- Error rate: 0%

[Report](postman-memorystorage-5millionentries-report.html)

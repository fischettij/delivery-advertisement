# Delivery Advertisement application

### Clone

```bash
git clone git@github.com:fischettij/delivery-advertisement.git
cd delivery-advertisement
```

#### Run with docker
Configure the following env vars in docker-compose.yml

File Download: 
* `CSV_RESOURCE_URL=https://url-to-file.com/file.csv`: Url for file downloading.
* `FILE_POLLING_INTERVAL_MINUTES=10`: Time interval in minutes for how often the file is searched


```bash
sudo docker-compose up
```

### Interface
The application will enable the following endpoints.
```
GET /delivery-services
  Query parameters:
    - lat: Mandatori. Is the latitude of the user location. Ej: lat=50.053419699999999&
    - lon: Mandatori. Is the longitude of the user location. Ej: lon=8.6705214000000002

  Response:
    - Status: 200
    - Content Type: application/json
    - Body: 
      {
        ids: [] // string slice of deliveries establishments  
      }
```
``` curl --location 'http://localhost:8080/delivery-services?lat=50.053419699999999&lon=8.6705214000000002'```

### Future work and ideas
- Document Api error responses
- Instead of shutting down the app when the file is not available, or if it's not possible to update the information, retries or another recover strategy for the downloader should be implemented.
- In memory storage:
  - Try another data organization, like a matrix, and compare response times between implementations.
  - To avoid information downtime, a second temporal database in Memory could be implemented, in order to load the new entries before discarding the old ones.
- I would like to explore if SQLIte in memory could perform better for this solution. My hypothesis is that it wouldn't, but it would be best to run a POC.

### Discarded proof of concept
- Postgres database: The idea was to test PostGIS to use arithmetic operations and test faster filtering. It didn't work, so I tried to pre-filter the largest number of values in the database and then do the distance calculations in the app to obtain the final result. The benchmark for this solution was not efficient as expected, the database turned out to be a bottleneck. That code is not available in this package storage.

### Benchmark

Can be executed with Artillery or Postman.

I chose postman because I found it easier to generate the reports and read the data. Artillery offers more flexibility to test but, for good data visualization, it requires integration with external services.

#### Artillery
Requires installation of npm or npx installed

How to run the benchmark
```
cd delivery-advertisement
docker-compose up -d
cd benchmark
npm install
npx artillery run --output report.json load-test.yml
npx artillery report report.json
open file report.json.html
```

#### Postman

It is a good alternative to do a quick test of the application. It has the limitation that only 100 requests per second can be configured.

- Import collection from ./benchmark/Delivery.postman_collection.json in postman
- Run with performance test option. [Postman docs](https://learning.postman.com/docs/collections/performance-testing/performance-test-metrics/)

### Result
You can check the benchmark results in [./benchmark/README.md](./benchmark/README.md) 

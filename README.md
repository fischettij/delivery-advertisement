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
* `FilePollingIntervalMinutes=10`: Time interval in minutes for how often the file is searched 

Database: The following values has a default but can be changed. The change has to be made in `advertisement` and `db` containers
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

```bash
sudo docker-compose up
```

### Interface
The application will make available the following endpoints.
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

#### Ideas and TODO
- Document Api errors response
- Select between implemented databases by config
- Add retries or another recover meth for downloader. It is not good to kill the app if the file is not available or I cannot update the information. Do something smarter
- Database avoid downtime. Two alternatives:
  - Use two databases. While database A is in use, load information in a database B. When information was loaded, witch queries to database B.
  - Use a second temporal database (SQLite or in Memory), load data in temp database, switch selects to the temp database while main database is loading the new data. When main datase finish switch selects and destroy temp database
- Postgres storage:
  - Add tests for file postgres.go.
- In memory storage:
  - Try another data organization like a matrix and compare response times between implementations.
  - In list (actual), insert values sorted by latitude using quick sort or similar algorithm.
  - Search with divide and conquer strategy analyzing elements in rage of latitude -0,25 and latitude +0,25.
- Will SQLIte in memory be performant for this solution? Maybe a poc is not difficult to do. 

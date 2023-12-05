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

Database: The following values has a default but can be changed. The change has to be made in `advertisement` and `db` containers
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

```bash
sudo docker-compose up
```

#### Ideas and TODO
- Select between implemented databases by config
- Database avoid downtime. Two alternatives:
  - Use two databases. While database A is in use, load information in a database B. When information was loaded, witch queries to database B.
  - Use a second temporal database (SQLite or in Memory), load data in temp database, switch selects to the temp database while main database is loading the new data. When main datase finish switch selects and destroy temp database
- Postgres storage:
  - Add tests for file postgres.go.
  - Use two databases to avoid having times with the database not available. While database A is used, the information is loaded into database B. When the information is finished loading, a switch is made to base B for new queries.
- In memory storage:
  - Try another data organization like a matrix and compare response times between implementations.
  - In list (actual), insert values sorted by latitude using quick sort or similar algorithm.
  - Search with divide and conquer strategy analyzing elements in rage of latitude -0,25 and latitude +0,25.
- Will SQLIte in memory be performant for this solution? Maybe a poc is not difficult to do.

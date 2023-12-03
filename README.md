# Delivery Advertisement application

### Clone

```bash
git clone git@github.com:fischettij/delivery-advertisement.git
cd delivery-advertisement
npm install
cp .env.example .env
```

Modificar los valores en el archivo .env` según corresponda.

#### Run with docker
Configure the following env vars in docker-compose.yml
* `CSV_RESOURCE_URL=https://url-to-file.com/file.csv`: Url for file downloading.

```bash
sudo docker-compose up
```

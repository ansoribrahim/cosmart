# cosmart

how to run :

1. make sure host in .env file same as you ip, or you can use localhost
2. make sure you have docker installed
3. run docker-compose up


list of api :

1. get list of book : http://localhost:8080/books?subject=history
2. schedule booking : http://localhost:8080/schedule-pickup

    body
```
{
    "title": "halo",
    "author": "jude",
    "editionNumber": 423,
    "time": "2022-10-23"
}
```

3. listing booking schedule : http://localhost:8080/schedule-pickup?time=2022-10-23



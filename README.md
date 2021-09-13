# goReactApp

goReactApp is a simple platform that displays users from different countries and show how many users are there in each country. The data is arleady seeded from a csv file.

## Tech

goReactApp uses:

- [Go] - For the backend service
- [React] - For the frontend service
- [PostgresSQL] - As the database

## Installation
goReactApp requires [Docker](https://www.docker.com/) and [Docker compose](https://docs.docker.com/compose/install/) to run.

Clone the repo and then open a terminal in the location of the cloned repo and run the next command
```sh
docker-compose up -d --build
```
Wait for it to finish before you can start.

Verify everything is working by navigating to your browser and vist 

```sh
localhost:3000
```

## License

MIT

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   
   [Go]: <https://golang.org/>
   [React]: <https://reactjs.org/>
   [PostgresSQL]: <https://www.postgresql.org/>

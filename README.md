# URL Shortner
A simple URL shortener service built with Go and PostgreSQL as database.
It provides REST APIs to shorten long URLs and redirect shortened URLs back to their original destinations. And also get The Original url. 
It uses the client to call CreatePostgresRow , GetPostgresRow functions from the previous project ( In-Memory-Database). [Visit](https://github.com/imsumedhaa/In-memory-database "In memory database") 


Here I used docker to pull the image of postgres from the Docker Hub. 

## Docker Command:

    docker run --name my-postgres-db \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=SecretPassword \
    -e POSTGRES_DB=mydb \
    -p 5431:5432 \
    -v pgdata:/var/lib/postgresql/data \
    -d postgres
pgdata for persistent volume.

## First start postgres before running the application:
**Step 1:** Start the container(if not already running):
```
docker start my-postgres-db
```
**Step 2:** Connect to the PostgreSQL database:

    psql -h localhost -p 5431 -U admin -d mydb

It will prompt for the password.

**Step 3:** Run the application:
```
go run main.go
```

## curl commands:
```
curl -X POST http://localhost:8080/create     -H "Content-Type: application/json"     -d '{"OriginalURL":"https://www.youtube.com/"}'

curl -X DELETE http://localhost:8080/delete     -H "Content-Type: application/json"     -d '{"ShortURL":"http://localhost:8080/dba51bcc"}'

curl "http://localhost:8080/get?short=dba51bcc" 
   
curl -X GET http://localhost:8080/dba51bcc     # redirect 
```

<img src="https://img.freepik.com/free-vector/cute-girl-hacker-operating-laptop-cartoon-vector-icon-illustration-people-technology-isolated-flat_138676-9487.jpg?semt=ais_hybrid&w=740&q=80" alt="lco mascot" width="160" align="left"/>
 

<br clear="left"/>

>Keep learning, keep building, keep growing.



    

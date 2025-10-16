# URL Shortner
A simple URL shortener service built with Go and PostgreSQL as database.
It provides REST APIs to shorten long URLs and redirect shortened URLs back to their original destinations. And also get The Original url. 
It uses the client to call CreatePostgresRow , GetPostgresRow functions from the previous project [In-Memory-Database](https://github.com/imsumedhaa/In-memory-database) .


Here I used docker to pull the image of postgres from the Docker Hub. 

## 🐳 Docker Command:

    docker run --name my-postgres-db \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=SecretPassword \
    -e POSTGRES_DB=mydb \
    -p 5431:5432 \
    -v pgdata:/var/lib/postgresql/data \
    -d postgres
pgdata for persistent volume.

## ▶️ First start postgres before running the application:
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

## ⌨️ curl commands:
```
curl -X POST http://localhost:8080/create     -H "Content-Type: application/json"     -d '{"OriginalURL":"https://www.youtube.com/"}'

curl -X DELETE http://localhost:8080/delete     -H "Content-Type: application/json"     -d '{"ShortURL":"http://localhost:8080/dba51bcc"}'

curl "http://localhost:8080/get?short=dba51bcc" 
   
curl -X GET http://localhost:8080/dba51bcc     # redirect 
```
## 🐳 Dockerize the project:
Write the dockerfile and compose for the project then check everything is working fine or not using the command "docker compose up".

**Step 1:** Build the image:
```
docker build -t docker_username/project_name:v1.0.0 .
```
**Step 2:** Push the image into docker hub:
```
docker push docker_username/project_name:v1.0.0
```
## ☸️ Deploying the Application to Kubernetes

Once your Docker image is pushed to Docker Hub, you can deploy it to your Kubernetes cluster using a **Deployment** and a **Service**.

```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## 🔄 Accessing Your Application via Port Forwarding

You can use the following command to access applications or services running inside your Kubernetes cluster **as if they were running locally** on your machine:

```bash
kubectl port-forward service/api-service 8081:8500
```

## 🌐 Port Details

>**8081 → Local Machine**  
The port you access in your browser.  

Example:

`curl -X POST http://localhost:8081/create  -H "Content-Type: application/json"   -d '{"OriginalURL":"https://www.youtube.com/"}'`

>**8500 → Kubernetes Service**  
The internal port exposed by the Kubernetes Service inside the cluster.


>**8080 → Pod / Container**  
The port where your actual application is running inside the container.


So, any request you send to localhost:8081 will actually hit your Kubernetes service on port 8500.

## 🖥️ Node port:
NodePort is a type of Kubernetes Service that exposes your application so it can be accessed from outside the cluster.
It opens a specific port on every node in the cluster.

Any request to <NodeIP>:<NodePort> is forwarded to the service port, and then to the pod’s container port.
**Command:**
```
minikube service <service-name> --url
```
It will you a port in the range 30000–32767. Once you have the port, you can make a request to your service from your local machine. 

**For example:**
```
curl -X POST http://localhost:<nodePort>/create \
     -H "Content-Type: application/json" \
     -d '{"OriginalURL":"https://www.google.com/"}'
```


<img src="https://img.freepik.com/free-vector/cute-girl-hacker-operating-laptop-cartoon-vector-icon-illustration-people-technology-isolated-flat_138676-9487.jpg?semt=ais_hybrid&w=740&q=80" alt="lco mascot" width="160" align="left"/>
 

<br clear="left"/>

>Keep learning, keep building, keep growing.



    

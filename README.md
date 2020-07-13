# Memequotes Front
Frontend for Memequotes Application. Check the [backend](https://github.com/airabinovich/memequotes_back)

# Structure
This Project has two parts: The frontend and the api

## Api
The api is under `/api` and it's written in Go.
It works as the frontend server.  
To run this server just do
```sh
cd api
go run main.go
```

## Frontend
The frontend is under `/web` and is written in React.
To compile it use the script `/compile.sh`
That allows to hot-swap the frontend executing with the Go server.  
For this project to run correctly, you should have the backend running

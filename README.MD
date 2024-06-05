This is a Backend REST API Project using Golang

## Getting Started

Requirement

```bash
1. mysql database
2. Goland version > 1.20
```

First, create your mysql database. in this project, we create database using _tinderleaked_ as a database name.

Second, create _.env_ or edit the .env.example file on folder config with your own value

Run migrate for database using this command

```bash
migrate -database "mysql://{your_mysql_username}:{your_mysql_password}@tcp({your_mysql_host}:{your_mysql_port})/tinderleaked" -path db/migrations up
```

Run the development server:

```bash
go run main.go
```

Open [http://localhost:3000](http://localhost:{SERVICE_PORT}) with your browser to see the result.

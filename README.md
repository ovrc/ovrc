# ovrc

ovrc aims to deliver various monitoring tools for developers:
* HTTP monitoring/health checks.
* Server monitoring via agent. 
* DNS monitoring.
* SSL monitoring.
* ... more, maybe?

## Requirements:
* Go (currently developed under 1.13)
* NPM
* PostgreSQL
* Flyway (migrations)

## Development

All programs are located inside `backend/cmd`:
* **api**: HTTP api used by the frontend.
* **httpmonitor**: The program that runs the HTTP monitors.
```
cd backend/cmd/api
go build
./api
```

### Frontend:
```
npm run dev
```

### Database setup:
Postgres 11+:
```
CREATE USER ovrc;
CREATE DATABASE ovrc;
ALTER USER ovrc WITH ENCRYPTED PASSWORD 'ovrc';
GRANT ALL PRIVILEGES ON DATABASE ovrc TO ovrc;
```

### Migrations
Create a `flyway.conf` file at the root of the project with the following details:
```
flyway.url=jdbc:postgresql://127.0.0.1:5432/ovrc
flyway.user=ovrc
flyway.password=ovrc
flyway.locations=filesystem:./
```



###Running api dev server:
```
go build
./ovrc
```

### Running frontend server:
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

flyway.conf (migrations)
```
flyway.url=jdbc:postgresql://127.0.0.1:5432/ovrc
flyway.user=ovrc
flyway.password=ovrc
flyway.locations=filesystem:./
```



Running api dev server:
```
ovrc uvicorn api.main:app --reload --port 8002 --ssl-keyfile dev/certs/ovrc.test+4-key.pem  --ssl-certfile dev/certs/ovrc.test+4.pem
```

Running frontend server:
```
npm run dev
```
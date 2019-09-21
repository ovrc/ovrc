FROM node:alpine AS frontend
WORKDIR /app
COPY package*.json ./
COPY frontend ./frontend
RUN npm install && npm run build

FROM golang:latest AS binaries
ENV GO111MODULE=on
WORKDIR /app
COPY api /app
# TODO: Remove dev certificates for production.
COPY dev_certs /dev_certs
COPY --from=frontend /app/dist ./dist
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ovrc

FROM alpine:latest
WORKDIR /app
COPY --from=binaries /app .
# TODO: Remove dev certificates for production.
COPY --from=binaries /dev_certs /dev_certs
EXPOSE 8002
#CMD ./ovrc
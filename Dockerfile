FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOOS linux
ENV GOARCH amd64

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -tags=jsoniter -ldflags="-s -w" -o /app/VacancyService cmd/vacancy_service/main.go

# Второй этап: запуск
FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/VacancyService /app/VacancyService

CMD ["./VacancyService"]
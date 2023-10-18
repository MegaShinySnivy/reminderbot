FROM golang:1.21.3
LABEL org.opencontainers.image.source="https://github.com/MegaShinySnivy/reminderbot"
WORKDIR /reminderbot
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY bot/bot.go ./bot/bot.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /reminderbot-bin
CMD ["/reminderbot-bin"]
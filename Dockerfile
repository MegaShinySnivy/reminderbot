FROM golang:1.22.6
LABEL org.opencontainers.image.source="https://github.com/MegaShinySnivy/reminderbot"
WORKDIR /reminderbot
COPY . .
RUN go build -o /reminderbot-bin
CMD ["/reminderbot-bin"]
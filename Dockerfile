FROM golang:1.22-alpine AS backend
COPY . /go/src/github.com/beadko/mywinebook
WORKDIR /go/src/github.com/beadko/mywinebook
RUN apk add --no-cache build-base
RUN go mod download
RUN CGO_ENABLED=1 go build -o mywinebook .

FROM node:18-alpine AS frontend-build
COPY frontend /opt/my-wine-book/
WORKDIR /opt/my-wine-book/
RUN npm install
RUN npm run build

FROM alpine
WORKDIR /opt/mywinebook/
COPY --from=backend /go/src/github.com/beadko/mywinebook/mywinebook bin/
COPY --from=frontend-build /opt/my-wine-book/dist static
RUN chmod +x bin/mywinebook
RUN mkdir -p data/images
RUN ["/opt/mywinebook/bin/mywinebook", "init"]
CMD ["/opt/mywinebook/bin/mywinebook", "server"]
EXPOSE 80
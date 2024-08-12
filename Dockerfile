FROM golang:1.22-alpine AS backend
COPY . /go/src/github.com/beadko/mywinebook
WORKDIR /go/src/github.com/beadko/mywinebook
RUN apk add --no-cache build-base
RUN go mod download
RUN CGO_ENABLED=1 go build -o mywinebook .

#FROM node:18-alpine AS frontend-build
#WORKDIR /opt/my-wine-book/
#COPY frontend2/ /opt/my-wine-book
#RUN npm install
#RUN npm run build

FROM alpine
WORKDIR /opt/mywinebook/
COPY --from=backend /go/src/github.com/beadko/mywinebook/mywinebook bin/
#COPY --from=frontend-build /opt/my-wine-book/dist static
COPY static static
RUN ["bin/mywinebook", "init"]
CMD ["bin/mywinebook", "server"]
EXPOSE 80
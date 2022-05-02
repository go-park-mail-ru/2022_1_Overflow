## We specify the base image we need for our
## go application
FROM golang:1.12 AS build
## We create an /app directory within our
## image that will hold our application source
## files
RUN apt update && apt install ca-certificates libgnutls30 -y

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o auth ./cmd/auth/run_auth.go
RUN go build -o folder_manager ./cmd/folder_manager/run_folder_manager.go
RUN go build -o mailbox ./cmd/mailbox/run_mailbox.go
RUN go build -o profile ./cmd/profile/run_profile.go
RUN go build -o repository ./cmd/repository/run_repository.go
RUN go build -o main /cmd/app/main.go

FROM alpine:latest AS runtime
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app
COPY --from=build /app ./
## Our start command which kicks off
## our newly created binary executable
CMD (bash -C '/app/repository > repository.log' &) && \
        (bash -C '/app/profile > profile.log' &) && \
        (bash -C '/app/mailbox > mailbox.log' &) && \
        (bash -C '/app/folder_manager > folder_manager.log' &) && \
        (bash -C '/app/auth > auth.log' &) && \
        /app/main
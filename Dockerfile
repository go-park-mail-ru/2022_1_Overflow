## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /app directory within our
## image that will hold our application source
## files
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
## Our start command which kicks off
## our newly created binary executable
CMD [   "(nohup /app/repository > repository.log)", "&&",
        "(nohup /app/profile > profile.log)", "&&",
        "(nohup /app/mailbox > mailbox.log)", "&&",
        "(nohup /app/folder_manager > folder_manager.log)", "&&",
        "(nohup /app/auth > auth.log)", "&&",
        "/app/main"]
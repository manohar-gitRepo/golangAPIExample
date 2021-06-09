FROM golang

#RUN go mod init github.com/golangbot/mysqltutorial

# RUN go get -u github.com/go-sql-driver/mysql

# RUN go get github.com/manohar-gitRepo/golangAPIExample
COPY . /golangAPIExample
WORKDIR /golangAPIExample
#copy files
# COPY . /golangAPIExample

#gooo build
RUN go build /golangAPIExample/api.go

EXPOSE 8000

# run the binary
CMD ["/golangAPIExample/api"]

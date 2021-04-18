FROM golang

LABEL maintainer="Nami Ops <namiops@gmail.com>"
 
RUN mkdir /app
COPY . /app
WORKDIR /app
 
# Golang build and output binary server file
RUN go build -o server .
 
# Run the server executable
CMD [ "/app/server" ]

EXPOSE 5050

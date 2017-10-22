# requires statically linked go binary to be compiled
# to ./magalix/simple-service before docker build
FROM golang:latest
# executable folder
RUN mkdir /magalix
COPY simple-service /magalix
# run micro-service after boot up
ENTRYPOINT ["/magalix/simple-service"]
#expose port 8080
EXPOSE 8080

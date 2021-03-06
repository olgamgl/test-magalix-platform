# requires statically linked go binary to be compiled
# to ./magalix/simple-service before docker build
FROM golang:latest
# executable folder
RUN mkdir /olgamgl
COPY simple-service /olgamgl
# run micro-service after boot up
ENTRYPOINT ["/olgamgl/simple-service"]
#expose port 8080
EXPOSE 8080

FROM scratch
MAINTAINER Patrick Easters <patrick@easte.rs>
COPY ./example-web-service /
EXPOSE 8080
ENTRYPOINT ["/example-web-service"]

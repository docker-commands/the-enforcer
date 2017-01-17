FROM alpine
RUN chmod +x /the-enforcer
ENDPOINT ["/the-enforcer"]
CMD [-h]


FROM busybox
COPY the-enforcer.sh /the-enforcer
RUN chmod +x /the-enforcer
ENDPOINT ["/the-enforcer"]
CMD [-h]

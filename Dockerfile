FROM scratch

ENV SERVICE_PORT 8081

EXPOSE $SERVICE_PORT

COPY receipt-print /

CMD ["/receipt-print"]
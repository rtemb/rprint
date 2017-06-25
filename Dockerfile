FROM alpine:3.5

ENV SERVICE_PORT 8080

EXPOSE $SERVICE_PORT

COPY rprint /app/

COPY fonts/ /app/fonts/
RUN ls -la /app/fonts/*

COPY images/ /app/images/
RUN ls -la /app/images/*

COPY receiptCustom/ /app/receiptCustom/
RUN ls -la /app/receiptCustom/*

COPY receiptSchema /app/receiptSchema/
RUN ls -la /app/receiptSchema/*

WORKDIR app

CMD ["./rprint"]
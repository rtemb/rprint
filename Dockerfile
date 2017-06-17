FROM alpine:3.5

ENV SERVICE_PORT 8081

EXPOSE $SERVICE_PORT

COPY receipt-print /app/

COPY fonts/ /app/fonts/
RUN ls -la /app/fonts/*

COPY images/ /app/images/
RUN ls -la /app/images/*

COPY receiptCustom/ /app/receiptCustom/
RUN ls -la /app/receiptCustom/*

COPY receiptSchema /app/receiptSchema/
RUN ls -la /app/receiptSchema/*

WORKDIR app

CMD ["./receipt-print"]
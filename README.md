Simple app for printing receipts and store it in file system. 
Provides the link to printed receipt.

# To create receipt:
## Request:
curl -X POST -d '{"id":1,"name":"apple","price":10.00, "bill":"000913"}' http://localhost:8081/create

## Response: 
{"link":"http://localhost:8081/pdf/1496491926257726883"

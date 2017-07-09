Simple app for printing receipts and store it in file system. 
Provides the link to printed receipt.

### Travis CI Status: [![Build Status](https://travis-ci.org/rtemb/rprint.svg?branch=master)](https://travis-ci.org/rtemb/rprint)

To launch the app need to export PORT environment variable: 

```
export PORT=8081
./rprint
```
Or you can launch app from latest docker container:
```
docker run -p 8081:8080 --rm -it rtemb/rprint:latest
```

Try to execute following curl requests:

## To create receipt:

### Request example. Schema receipt:
curl -X POST -d '{"Schema": "default", "ReceiptS": {"MPlaceName": "Exmaple header", "MPlaceAddress": "www.example.com",  "MPlaceINN": "00000111111239990", "OperationType": "Sell", "Items": [{"Name": "Raincoat", "Quantity": 1.000, "Price": 100.0}, {"Name": "Black Hat", "Quantity": 1.000, "Price": 33.0}, {"Name": "Gloves", "Quantity": 1.000, "Price": 15.0}], "TaxPercent": "18%", "Total": 148.0, "FiscalNumber": "000000000011198", "Date": "2017-06-11 23:21:11"}}' http://localhost:8081/v1/create

### Request example. Coustom receipt (alpha):
curl -X POST -d '{"PageConfig": {"Orientation": "P",  "Format": "A4",  "FontStyle": "I"}, "Instructions": [{"Type": "text", "Value": "www.example.com", "LineConfig": {"FontSize": 16.0, "Width": 0, "Height": 7.0, "NewLine": 1, "Align": "C"}}, {"Type": "text", "Value": "1. Apple", "LineConfig": {"FontSize": 16.0, "Width": 5.0, "Height": 7.0, "NewLine": 0, "Align": "L"}}, {"Type": "text", "Value": "1 pc", "LineConfig": {"FontSize": 16.0, "Width": 150.0, "Height": 7.0, "NewLine": 0, "Align": "C"}}, { "Type": "text", "Value": "5.00 $", "LineConfig": {"FontSize": 16.0, "Width": 0, "Height": 7.0, "NewLine": 1, "Align": "R"}}, {"Type": "text", "Value": "2. Chocolate", "LineConfig": {"FontSize": 16.0, "Width": 5.0, "Height": 7.0, "NewLine": 0, "Align": "L"}}, {"Type": "text", "Value": "2 pc", "LineConfig": {"FontSize": 16.0, "Width": 150.0, "Height": 7.0, "NewLine": 0, "Align": "C"}}, {"Type": "text", "Value": "10.30 $", "LineConfig": {"FontSize": 16.0, "Width": 0, "Height": 7.0, "NewLine": 1, "Align": "R"}}, {"Type": "nl", "Value": "3", "LineConfig": {"FontSize": 0, "Width": 0, "Height": 0, "NewLine": 0, "Align": ""}}, {"Type": "text", "Value": "TOTAL: ", "LineConfig": {"FontSize": 16.0, "Width": 0, "Height": 7.0, "NewLine": 0, "Align": "L"}}, {"Type": "text",   "Value": "15.30 $", "LineConfig": {"FontSize": 16.0, "Width": 0, "Height": 7.0, "NewLine": 1, "Align": "R"}} ]}' http://localhost:8081/v1/createcustom

## Response example: 
{"link":"http://localhost:8081/v1/pdf/1496491926257726883"}
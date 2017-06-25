Simple app for printing receipts and store it in file system. 
Provides the link to printed receipt.

### Travis CI Status: [![Build Status](https://travis-ci.org/rtemb/rprint.svg?branch=master)](https://travis-ci.org/rtemb/rprint)

## To create receipt:

### Request example. Schema receipt:
curl -X POST -d '{ "Schema": "default", "ReceiptS": {  "MPlaceName": "Exmaple header",  "MPlaceAddress": "www.example.com",  "MPlaceINN": "00000111111239990",  "OperatinType": "Sell",  "Items": [   {    "Name": "Raincoat",    "Quantity": 1.000,    "Price": 100.0   },    {    "Name": "Black Hat",    "Quantity": 1.000,    "Price": 33.0   },   {    "Name": "Gloves",    "Quantity": 1.000,    "Price": 15.0   }  ],  "TaxPercent": "18%",  "Total": 148.0,  "FiscalNumber": "000000000011198",  "Date": "2017-06-11 23:21:11" }}' http://localhost:8081/create

### Request example. Coustom receipt (not implemented yet):
curl -X POST -d '{ "FileSetting": {  "Format":"A4",  "ZeroX":10,  "ZeroY":10 }, "ReceiptN": {  "MPlaceName": {   "Value": "www.example.com",   "LineSetting": {    "Fsize": 16,    "Fstyle": "I",    "PosX": 10,    "PosY": 14,    "Align": "L"   }  },  "MPlaceAddress": {   "Value": "Russia, Smolensk region, Smolesk Krasninskpe highway 195",   "LineSetting": {    "Fsize": 16,    "Fstyle": "I",    "PosX": 10,    "PosY": 18,    "Align": "L"   }  },  "MPlaceINN": {   "Value": "0000001111229990",   "LineSetting": {    "Fsize": 16,    "Fstyle": "I",    "PosX": 10,    "PosY": 22,    "Align": "L"   }  },  "OperationType": {   "Value": "Продажа",   "LineSetting": {    "Fsize": 16,    "Fstyle": "I",    "PosX": 10,    "PosY": 26,    "Align": "L"   }  } }}' http://localhost:8081/createcustom

## Response example: 
{"link":"http://localhost:8081/pdf/1496491926257726883"
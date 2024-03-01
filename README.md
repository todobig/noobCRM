# noobCRM [golang + mongoDB]
a basic crm - add / retrieve / update / delete
----------------

###clone and compile the main.go 
### go build -o noobCRM main.go
### ./noobCRM


###pre requisites
mongoDB 



-----------------------------------------------------
API Documentation
Base URL
The base URL for all endpoints is https://vpsIP:8080.

Endpoints
1. Add Lead
URL: /leads
Method: POST
Description: Add a new lead to the system.
Request Body: JSON object representing the lead information.

{
    "first_name": "John",
    "last_name": "Doe",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001",
    "phone": "1234567890",
    "email": "john@example.com"
}
Response: 201 Created if successful, along with the added lead object.


Example:
```
curl -X POST https://vpsIP:8080/leads \
     -H "Content-Type: application/json" \
     -d '{
            "first_name": "John",
            "last_name": "Doe",
            "address": "123 Main St",
            "city": "New York",
            "state": "NY",
            "zip": "10001",
            "phone": "1234567890",
            "email": "john@example.com"
         }'
```
--------------------------------------------------------------------------------------

2. Get Lead
URL: /leads?phone=<phone_number>
Method: GET
Description: Retrieve lead information by phone number.
Query Parameters:
phone: The phone number of the lead to retrieve.
Response: 200 OK if successful, along with the lead object.
Example:
```
curl -X GET https://vpsIP:8080/leads?phone=1234567890

{"first_name":"John","last_name":"Doe","address":"123 Main St","city":"New York","state":"NY","zip":"10001","phone":"1234567890","email":"john@example.com"}
```
-------------------------------------------------------------------------------------

3. Update Lead
URL: /leads?phone=<phone_number>
Method: PUT
Description: Update lead information by phone number.
Query Parameters:
phone: The phone number of the lead to update.
Request Body: JSON object representing the updated lead information.

{
    "first_name": "John",
    "last_name": "Doe",
    "address": "456 Elm St",
    "city": "New York",
    "state": "NY",
    "zip": "10001",
    "email": "john@example.com"
}
Response: 200 OK if successful, along with the updated lead object.
Example:
```
curl -X PUT https://vpsIP:8080/leads?phone=1234567890 \
     -H "Content-Type: application/json" \
     -d '{
            "first_name": "John",
            "last_name": "Doe",
            "address": "456 Elm St",
            "city": "New York",
            "state": "NY",
            "zip": "10001",
            "email": "john@example.com"
         }'

```
--------------------------------------------------------------------------------------

4. Delete Lead
URL: /leads?phone=<phone_number>
Method: DELETE
Description: Delete lead information by phone number.
Query Parameters:
phone: The phone number of the lead to delete.
Response: 200 OK if successful, along with a success message.
Example:
```
curl -X DELETE https://vpsIP:8080/leads?phone=1234567890

{"message":"Lead deleted successfully"}
```

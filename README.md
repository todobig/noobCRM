# noobCRM
a basic crm
----------------

mkdir 1billion
cd 1billion
clone and extract 

-----------------------------------------------------
API Documentation
Base URL
The base URL for all endpoints is http://vpsIP:8080.

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

curl -X POST http://vpsIP:8080/leads \
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



2. Get Lead
URL: /leads?phone=<phone_number>
Method: GET
Description: Retrieve lead information by phone number.
Query Parameters:
phone: The phone number of the lead to retrieve.
Response: 200 OK if successful, along with the lead object.
Example:
curl -X GET http://vpsIP:8080/leads?phone=1234567890

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

curl -X PUT http://vpsIP:8080/leads?phone=1234567890 \
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


--------------------------------------------------------------------------------------

4. Delete Lead
URL: /leads?phone=<phone_number>
Method: DELETE
Description: Delete lead information by phone number.
Query Parameters:
phone: The phone number of the lead to delete.
Response: 200 OK if successful, along with a success message.
Example:
bash
Copy code
curl -X DELETE http://vpsIP:8080/leads?phone=1234567890

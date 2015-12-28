# cmpe273-Assignment2

Configurations required to execute project
1)	In server/services.go file
-	Go to method connectionWithDB
func connectionWithDB() (sess*mgo.Session, dbName string) {
-	Replace uri value with mongolab configured database
uri="mongodb://admin:admin@ds043694.mongolab.com:43694/addressbook"
-	Replace database(dbName) name with your configured database name like dbName = "addressbook"


----------------------
Command Line Arguments

1) Start server using the following command, server will start listening on port 8080
    Go run startserver.go

2) Start postman or any other REST client 
1.	Create New Location - POST  /locations 
Provide JSON in format as below: 
{
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}
Expected output, in below format:
{
  "Id": "562b17c52ff8612e86000001",
  "Name": "John Smith",
  "Address": "123 Main St",
  "City": "San Francisco",
  "State": "CA",
  "Zip": "94113",
  "Coordinate": {
    "Lat": 37.7917618,
    "Lng": -122.3943405
  }
}

2.	Get a Location - GET   /locations/{location_id}
Expected output, in below format:

{
  "Id": "562b17c52ff8612e86000001",
  "Name": "John Smith",
  "Address": "123 Main St",
  "City": "San Francisco",
  "State": "CA",
  "Zip": "94113",
  "Coordinate": {
    "Lat": 37.7917618,
    "Lng": -122.3943405
  }
}

3.	Update a Location - PUT /locations/{location_id}
{
   "address" : "4th srt",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}

Expected output, in below format:
{
  "Id": "562b17c52ff8612e86000001",
  "Name": "John Smith",
  "Address": "4th srt",
  "City": "San Francisco",
  "State": "CA",
  "Zip": "94113",
  "Coordinate": {
    "Lat": 37.7819032,
    "Lng": -122.5025342
  }
}

4.	Delete a Location - DELETE /locations/{location_id}
Expected output:
•	HTTP Response Code: 200







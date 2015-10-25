package server

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2/bson"
)

type LocationRequest struct {
    Name string
    Address string
    City string
    State string
    Zip string
}

type Coordinates struct {
  Lat float64
  Lng float64
}

type LocationResponse struct {
  Id bson.ObjectId
  Name string
  Address string
  City string
  State string
  Zip string
  Coordinate Coordinates
}

type LocationCoordinates struct {
  Results[]struct {
    Geometry struct {
      Location struct {
        Lat float64
        Lng float64
      }
    }
  }
}

func updateLocationHandler(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  var locationRequest LocationRequest
  var locationCoordinates LocationCoordinates
  idValue := convertToObjectId(p.ByName("location_id"))
  jsonDataFromHttp, err := ioutil.ReadAll(request.Body)
  if err != nil {
          panic(err)
  }
  err = json.Unmarshal([]byte(jsonDataFromHttp), &locationRequest)

  url:= MakeUrl(locationRequest)
  DataFromGoogelApi(url,&locationCoordinates)
  lat := locationCoordinates.Results[0].Geometry.Location.Lat
  lng := locationCoordinates.Results[0].Geometry.Location.Lng
  response :=  UpdateInDB(&locationRequest,lat,lng,idValue)
  jsonResponse, _ := json.Marshal(response)
  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(201) // Status code for success
  fmt.Fprintf(rw, "%s", jsonResponse)

}

func getLocationHandler(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
 idValue := convertToObjectId(p.ByName("location_id"))
 response := getFromDB(idValue)
 jsonResponse, _ := json.Marshal(response)
 rw.Header().Set("Content-Type", "application/json")
 rw.WriteHeader(200) // Status code for success
 fmt.Fprintf(rw, "%s", jsonResponse)
}

func createLocationHandler(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  var locationRequest LocationRequest
  var locationCoordinates LocationCoordinates
  jsonDataFromHttp, err := ioutil.ReadAll(request.Body)
  if err != nil {
          panic(err)
  }
  err = json.Unmarshal([]byte(jsonDataFromHttp), &locationRequest)

    fmt.Println(locationRequest.Name)
    url:= MakeUrl(locationRequest)
    DataFromGoogelApi(url,&locationCoordinates)
    lat := locationCoordinates.Results[0].Geometry.Location.Lat
    lng := locationCoordinates.Results[0].Geometry.Location.Lng
    response := SaveInDB(&locationRequest,lat,lng)
    jsonResponse, _ := json.Marshal(response)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201) // Status code for success
    fmt.Fprintf(rw, "%s", jsonResponse)
}

func deleteLocationHandler(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  idValue := convertToObjectId(p.ByName("location_id"))
  statusCode := delete(idValue)
  rw.WriteHeader(statusCode) // Status code for success
  fmt.Fprint(rw)
}

func Server() {
  fmt.Println("starting server...........")

  mux := httprouter.New()
  mux.GET("/locations/:location_id", getLocationHandler)
  mux.POST("/locations", createLocationHandler)
  mux.PUT("/locations/:location_id", updateLocationHandler)
  mux.DELETE("/locations/:location_id", deleteLocationHandler)

  server := http.Server{
          Addr:        "0.0.0.0:8080",
          Handler: mux,
  }
  server.ListenAndServe()

}

# JPPP

### Requirements:
```
You will need docker installed on your machine to run the application locally.
https://docs.docker.com/desktop/install/mac-install/ (mac link)

I've utilzied goose for handling db migrations.
https://github.com/pressly/goose
brew isntall goose.

Copy the .env.sample file to .env and replace the values with your db configuration.

Once the you have docker running you can spin up the web server and db by running:
`make run`

Once the services are running you can validate the status by curl'ing:
`curl -vvv -X GET 'http://localhost:8000/v1/status'`

Once the webserver has been validated you can run migrations by running:
`make migrate DB_USER={{your user}} DB_PASS={{your pass}} DB_NAME={{your db name}}`
```

### Routes
POST	/v1/cages<br>
POST	/v1/dinosaurs<br>
PATCH	/v1/cages/:id<br>
PATCH	/v1/dinosaurs/:id<br>
PATCH	/v1/cages/:id/dinosaurs/:id<br>
DELETE	/v1/cages/:id/dinosaurs/:id<br>
GET	    /v1/cages<br>
GET	    /v1/cages/:id/dinosaurs<br>
GET	    /v1/dinosaurs<br>
GET	    /v1/cage/:id<br>
GET	    /v1/dinosaur/:id<br>
GET	    /v1/dinoaurs/species<br>

### MODELS
```
Cage
{
    "id": "uuid",
    "type": "string EMUM", (HERBIVOR, CARNIVORE)
    "capacity": int,
    "currentCapacity": int,
    "status": "string ENUM", (ACTIVE, DOWN)
    "createdAt": int,
    "updatedAt": int
}

Dinosaur
{
    "id": "uuid",
    "cage_id": "uuid nullable",
    "name": "string",
    "species": "string ENUM", (Spinosaurus, Megalosaurus, Brachiosaurus, Stegosaurus, Ankylosaurus, Triceratops, Tyrannosaurus, Velociraptor)
    "diet": "string ENUM", (HERBIVOR, CARNIVORE)
    "createdAt": int,
    "updatedAt": int
}

API Error
{
  "error": {
    "code": "string ENUM", (BAD_REQUEST, INTERNAL_SERVER_ERROR, NOT_FOUND)
    "message": "string",
    "status_code": int,
    "details": {
      "string": "string"
    }
  }
}
```

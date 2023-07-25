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
POST	/v1/cages\n
POST	/v1/dinosaurs\n
PATCH	/v1/cages/:id\n
PATCH	/v1/dinosaurs/:id\n
PATCH	/v1/cages/:id/dinosaurs/:id\n
DELETE	/v1/cages/:id/dinosaurs/:id\n
GET	    /v1/cages\n
GET	    /v1/cages/:id/dinosaurs\n
GET	    /v1/dinosaurs\n
GET	    /v1/cage/:id\n
GET	    /v1/dinosaur/:id\n
GET	    /v1/dinoaurs/species\n

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
```

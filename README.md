This project was built utilizing Golang on the server, and ReactJS on the client.

The Client follows a Model-View-Presenter architecture. The Presenter interacts with the data source and tells the View
what to do. The View largely plays dumb, simply relaying events to the Presenter and doing whatever the Presenter tells 
it to do. I left Presenters out of the ResultViews as the role of those views was to simply display the props.

The Server's architecture was inspired by a CLEAN architecture. The controllers receive events/requests from the client,
and pass them off to the `usecase` package's service classes, whose responsibility is to perform the business logic. The controller
receives the results from the service in the form of an entity model, and then maps this entity model into a formatted 
API response model. If this project were to require a database, I would have created a separate `access` package which would 
define the repository interfaces. These repositories would be responsible for interacting with the database.

The startup script assumes the user has Docker installed on their computer
https://www.docker.com/

To build the application for the first time, run `docker-compose up`

Upon making code changes, run `docker-compose build && docker-compose up`

When the script is complete, visit `localhost:3000`
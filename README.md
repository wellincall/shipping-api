# Shipment api

This is a project implemented using Golang and Gin framework.

# Why Gin framework?

Gin framework abstracts how requests are handled and response are served by a webserver. Using it allowed me to focus only in implementing the behaviour expected in the endpoints.

# Storage

For the time being, adding a database to store the data given to the application would introduce to more dependencies/complexity. So for the sake of time, a csv file was used for this purpose.

# How it works?

This project exposes two urls:

```
GET /shipments -> Lists all shipments registered
POST /shipments -> Create a shipents with a price generated according to given pricing rules
```

# Testing?

Due to timing issues, there were no tests added to systematically check for the correct behaviour of the system.

# How to build it?

Make sure you have Golang installed in your machine.

Then run:

```
go build .
```

It is going to create an executable for the project.

Execute it with the following command:

`./shipping-api`

Now you have the application running on `localhost:8080`

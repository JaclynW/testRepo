# testRepo
This is used for a testing the functionality of Serverless Orchestrator.

In this repo is a very basic example of a Product Service microservice using Go and the Gin web framework for routing and handling HTTP requests. This will provide endpoints to fetch all products, fetch a specific product by ID, and add a new product.

## To Run It
To run this, you would need to install Go and Gin (go get -u github.com/gin-gonic/gin), then you can simply run go run main.go inside the product-service directory, and the service will start on port 8080.

The product service utilizes the Gin web framework, which is set up to listen and serve on port 8080. This means that the service will be waiting for HTTP requests (like GET, POST, etc.) on that port.

When you run the service, you can access it through a web browser, cURL, or any other tool or software that can send HTTP requests by navigating to http://localhost:8080 or sending requests to that URL.

For instance, with the provided code, if you navigate to http://localhost:8080/products, it should trigger the handler associated with the /products endpoint and return the list of products.

Remember, the port number (in this case, 8080) is just a choice. You can configure the Gin server to listen on a different port if needed. However, 8080 is a common default for many web services.

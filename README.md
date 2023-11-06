## Go CLI Application
This is a Go Command Line Interface (CLI) application that performs specific actions with URLs. You can use this application to fetch metadata from URLs and save the output in a specified directory.

### How to Build
Before running the application, you need to build a Docker image for it. Follow these steps:

Open your terminal.


Build the Docker image:
```
docker build --file build/Dockerfile -t fetch . --no-cache
```

### How to Run and Fetch Metadata
After building the Docker image, you can run the application to process URLs and save the output in the current directory. You can pass an array of URLs as arguments. Follow these steps:

Open your terminal.

Run the Docker container:
```
docker run -v "$(pwd)":/app/output fetch https://google.com https://twitter.com
```
fetch is the name of Docker image.
https://google.com and https://twitter.com are examples of URLs. You can pass an array of strings as needed.
The application will fetch metadata from the provided URLs and save the output in the output folder of your current directory.

### How to Print Metadata
You can also print the metadata from a single URL. But you need to run  Follow these steps:

Open your terminal.

Run the Docker container with a single URL for metadata:
```
docker run -v "$(pwd)":/app/output fetch --metadata https://twitter.com
```
https://twitter.com is an example of a URL. You can provide a single string argument to fetch metadata for that URL.
The application will fetch metadata from the provided URL and save the output in the current directory.

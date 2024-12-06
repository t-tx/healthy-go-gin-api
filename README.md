## Overview
Healthy is a Go-based application designed to manage and track health-related data. This project includes various features such as user management, user's event tracking, user's diary entries, and public articles

## Problem
Managing and tracking health-related data can be cumbersome and error-prone when done manually. Users need a reliable and efficient way to record their health activities, monitor progress, and access their data securely.

## Solution
Healthy provides a comprehensive solution for managing health-related data. It offers features such as user authentication, event tracking, diary entries, and data visualization. The application ensures data integrity and security, making it easy for users to keep track of their health activities.

## Verification Procedure
To verify the functionality of the Healthy application, follow these steps:

1. **Run the Application**:
   - Use the following command to start the application:
     ```sh
     make run
     ```
   - The application will start on the port specified in the [config.yml](./config.yml) file.

2. **API Endpoints**:
   - Use tools like Postman or cURL to interact with the API endpoints. 
   - Postman collection: postman/Healthy.postman_collection.json
   - Verify user authentication, event tracking, and diary entry functionalities by making appropriate API requests.

## Build the Docker Image
   - To build the Docker image, run:
     ```sh
     make build
     ```
   - Output docker image name: go-healthy-api

## Development Container
This project includes a development container configuration for Visual Studio Code.

- **Image**: `golang:1.20`
- **VS Code Extensions**: `golang.go`

To use the development container, open the project in Visual Studio Code and follow the prompts to reopen in the container.

By following these steps, you can verify that the Healthy application is working as intended and effectively managing health-related data.

## Project structure
   ```text
   .
   |-- documents                 -> Contains documents.
   |-- scripts                   -> Database migration scripts
   |-- internal                  
   |   |-- api
   |   |   `-- v1                -> The API route path is registered.
   |   |-- cmd                   -> Commands that are available.
   |   |-- config                -> Format of the configuration.
   |   |-- database              -> Schema and migrations for the database.
   |   |   |-- models            -> Models representing data.
   |   |   `-- repositories      -> Functions to interact with the database.
   |   |-- handler               -> Parsing http request, mapping business implementation, set response.
   |   |   |-- public            -> Business logic for publicly accessible logic.
   |   |   `-- user              -> Business logic for user logic.
   |   `-- pkg
   |       |-- defined           -> Constants.
   |       |-- enc               -> Encryption functionality.
   |       `-- utils             -> Common utility package.
   `-- test
   ```

## Documentation
- db table required: [database.go](internal/database/database.go)
- response json data of the required api, and their connections in a diagram: [postman.json](documents/postman.json) and [API.md](documents/API.md)

# Todo App

A todo app built using Turso and Go.
[![Netlify Status](https://api.netlify.com/api/v1/badges/0a743edc-8c45-49a3-ba9d-922c15554c3c/deploy-status)](https://app.netlify.com/sites/go-react-todo/deploys)

## Demo

You can see a demo of this project at the [**demo**](https://go-react-todo.netlify.app/)

## Technologies Used

- Frontend: React, Chakra UI,Vite
- Backend: Turso, Go

## Environment Variables

The following environment variables are required for the project to function properly:

- `BACKEND_PORT`: The port number on which the backend server will run.
- `DB_URL`: The URL of the database used by the project.
- `DB_TOKEN`: The token required to authenticate with the database.
- `GO_VERSION`: The version of Go used for building the backend.
- `IS_LOCAL`: A boolean value indicating whether the project is running locally or not.

Make sure to set these environment variables before running the project.

## Running Locally

To run the project locally, follow these steps:

### Backend

1. Create a new file named `.env` in the `backend/api` directory based on `.env.local.example`.
2. Run the following command to start the backend server:
   ```shell
   go run main.go
   ```
3. The backend server will start and listen on the port specified by the `BACKEND_PORT` environment variable.

### Frontend

1. Create a new file named `.env` in the `client` directory based on `.env.local.example`.

2. Install the project dependencies by running the following command in the `client` directory:
   ```shell
   npm install
   ```
3. Start the development server by running the following command in the `client` directory:
   ```shell
   npm run dev
   ```
4. The frontend will start and be accessible at url mentioned in the console.
 
## Deploying to Netlify

To deploy the project to Netlify, follow these steps:

1. Set the environment variables found in backend/api/.env.netlify.example within your Netlify project settings.

2. Set the environment variables from client/.env.netlify.example in your Netlify project settings, adjusting them according to your site's naming conventions.

3. Deploy the project to Netlify.

## API Documentation

For API documentation, please refer to the [Postman documentation](https://www.postman.com/hady-space/workspace/golang-for-node-devs/overview).

## License

This project is licensed under the [MIT License](LICENSE).

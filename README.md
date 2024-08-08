
```markdown
# Todo App

A todo app built using Turso and Go.

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

1. Set the environment variables mentioned above in a `.env` file in the `backend/api` directory.
2. Run the following command to start the backend server:
   ```shell
   go run main.go
   ```
3. The backend server will start and listen on the port specified by the `BACKEND_PORT` environment variable.

### Frontend

1. Install the project dependencies by running the following command in the `client` directory:
   ```shell
   npm install
   ```
2. Start the development server by running the following command in the `client` directory:
   ```shell
   npm run dev
   ```
3. The frontend will start and be accessible at url mentioned in the console.

## Deploying to Netlify

To deploy the project to Netlify, follow these steps:

1. Set the environment variables mentioned above in the Netlify project settings.
2. Configure the build command to build the backend server:
   ```shell
   go build -o main
   ```
3. Set the start command to run the backend server:
   ```shell
   ./main
   ```
4. Deploy the project to Netlify.

## API Documentation

For API documentation, please refer to the [Postman documentation](https://www.postman.com/hady-space/workspace/golang-for-node-devs/overview).

## License

This project is licensed under the [MIT License](LICENSE).
```

Feel free to customize this section to fit your project's specific needs. Let me know if you need any further assistance!

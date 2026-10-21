## Prerequisites

Before running this project, make sure you have:

- Go 1.22+
- MySQL server running locally
- A database named `sharing_vision`

## Environment Configuration

Create a `.env` file in the backend folder with the following values:

```env
DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_NAME=sharing_vision
DB_USER=root
DB_PASSWORD=
SERVER_PORT=8080
```

You can also copy from `.env.example`:

```bash
copy .env.example .env
```

## Run Locally

1. Start your MySQL service.
2. Make sure database `sharing_vision` exists or the app will create it automatically.
3. Open terminal in the backend folder and run:

```bash
go mod tidy
go run ./cmd/server
```

If you want to run with environment variables directly:

```bash
set DB_CONNECTION=mysql
set DB_HOST=localhost
set DB_PORT=3306
set DB_NAME=sharing_vision
set DB_USER=root
set DB_PASSWORD=
set SERVER_PORT=8080
go run ./cmd/server
```

For macOS/Linux shell:

```bash
export DB_CONNECTION=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_NAME=sharing_vision
export DB_USER=root
export DB_PASSWORD=""
export SERVER_PORT=8080
go run ./cmd/server
```

## API Endpoints

- `GET /health`
- `POST /article`
- `GET /article?limit=10&offset=0`
- `GET /article/:limit/:offset`
- `GET /article/:id`
- `PUT /article/:id`
- `PATCH /article/:id`
- `DELETE /article/:id`

## Example Request Body

```json
{
  "title": "This is a valid article title with at least 20 characters",
  "content": "This is a valid article content. It must be longer than 200 characters to satisfy validation requirements for the backend use case. This is example content that will pass validation criteria.",
  "category": "Technology",
  "status": "publish"
}
```

## Notes

- `DB_CONNECTION` should be `mysql`.
- `DB_HOST` can be `localhost` or `127.0.0.1`.
- `DB_PORT` is usually `3306`.
- `DB_NAME` is `sharing_vision`.
- `DB_USER` is typically `root` for local development.
- `DB_PASSWORD` should be empty if your local MySQL config uses no password.

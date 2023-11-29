# Client App

## API Endpoints

| HTTP Method | Endpoint                             | Description                         |
| ----------- |--------------------------------------| ----------------------------------- |
| POST        | `{{base_url}}/auth/register`         | Auth - Register                     |
| POST        | `{{base_url}}/auth/login`            | Auth - Login                        |
| POST        | `{{base_url}}/auth/logout`           | Auth - Logout                       |
| DELETE      | `{{base_url}}/user/:id`              | User - Delete user by id            |
| GET         | `{{base_url}}/user/:id`              | User - Update user data by id       |
| PUT         | `{{base_url}}/user/:id`              | User - Update user by id            |
| GET         | `{{base_url}}/user/:id`              | User - Get user by Id               |
| POST        | `{{base_url}}/booking`               | Booking - Create new booking        |
| GET         | `{{base_url}}/booking/:booking_id`   | Booking - Get Booking by id         |
| PUT         | `{{base_url}}/booking/:booking_id`   | Booking - Update booking by id      |
| DELETE      | `{{base_url}}/booking/:booking_id`   | Booking - Delete booking by id      |
| GET         | `{{base_url}}/user/:customer_id/bookings` | Booking - Get bookings from cust id |
| GET         | `(Not specified)`                      | Booking - Cancel Booking            |
| GET         | `(Not specified)`                      | Booking - Get All Bookings          |

## Notes
- Replace `{{base_url}}` with the actual base URL of the service.
- Some endpoints are not fully specified in this table. Please refer to the API documentation for complete details.

## API Docs
https://crimson-meadow-438973.postman.co/workspace/PAT~5e4b20a9-a21e-48b8-8eef-baeb56a29ad7/collection/30701742-ef76d9bf-81c4-4994-8d61-64c0b34c3457?action=share&creator=30701742&active-environment=30701742-3c17942c-d556-4a3c-b175-2402ac791441

## How to Use
1. Clone or fork this repository
```sh
https://github.com/NicholasLiem/IF4031_M1_Client_App
```
2. Initialize .env file using the template given (.env.example and docker.env.example)
```sh
touch .env
touch docker.env
```
3. Run docker compose and build
```sh
docker-compose up --build
```

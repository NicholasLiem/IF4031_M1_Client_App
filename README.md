# Go-REST-API-template
Building a RESTful API template repo using GoLang and Gorilla for learning purposes complete with continuous integration workflow

## How to Use
1. Clone or fork this repository
```sh
https://github.com/NicholasLiem/Go-REST-API-template.git
```
2. Initialize .env file using the template given (.env.example)
```sh
touch .env
```
3. Prepare database connection details and update .env
```text
You can use cloud database providers such as Supabase
```
4. To add routes, see: /user/routes.go
5. Run the server
```sh
go run .
```
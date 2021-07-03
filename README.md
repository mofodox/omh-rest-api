# omh-rest-api

This is a project for the Backend Engineer Exam at OhMyHome apprenticeship. 

The requirement is to create a restful api for a set of problem which is to let users post property they want to sell or rent out. The target audience are for users in Singapore, Malaysia and Philippines. 

I have created a set of endpoints where the user can do a simple CRUD operation which is: 

Property endpoints: 

- GET All Properties
- GET Single Property
- POST Single Property
- PUT Single Property
- DELETE Single Property

Country endpoints:

- GET All Countries
- GET Single Country
- POST Single Country
- PUT Single Country
- DELETE Single Property

The tech stack for this project are:

- GORM for MySQL ORM
- GoFiber for Routing
- godotenv for env values
- MySQL for DBMS

There’s a relationship between Property and Country where Property belongs to Country.

Property Model:

- gorm.Model
- Address
- Bedroom
- Bathroom
- Price
- Sqm
- Mode
- HomeType
- CountryID
- Country

Country model:

- gorm.Model
- Name

How to run the REST API server:

- If you’re using git, you can clone the project at [https://github.com/mofodox/ omh-rest-api](https://github.com/mofodox/omh-rest-api).
- Before you run the project make sure you create a .env file at the root project folder.
- In the .env file, you will need these values:

    ```
    DBUser=<mysql_username>
    DBPassword=<mysql_password>
    DBHost=<mysql_host>
    DBPort=<mysql_port>
    DBName=<mysql_name>
    ```

- Once you have cloned the project, you can run the app by typing this on the terminal (make sure you are inside of the project directory): `go run main.go`

Room for improvement/added:

- Dockerized the app
- CI/CD
- User authentication
- UI
- Data sanitization
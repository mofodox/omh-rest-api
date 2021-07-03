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

### Screenshots:

![Image1](https://user-images.githubusercontent.com/1651333/124359483-dc3a3600-dc57-11eb-9ceb-15dad360a3ec.png)

GET All properties (no data yet)

![Image2](https://user-images.githubusercontent.com/1651333/124359493-e6f4cb00-dc57-11eb-8724-ee9e148b6bd8.png)

POST Single Property with CountryID (Mode: Sale)

![Image3](https://user-images.githubusercontent.com/1651333/124359644-6c787b00-dc58-11eb-8ce4-28635eb3ef6b.png)

POST Single Property with CountryID (Mode: Rental)

![Image4](https://user-images.githubusercontent.com/1651333/124359679-992c9280-dc58-11eb-885b-69fd45ec6abb.png)

GET All Properties

![Image6](https://user-images.githubusercontent.com/1651333/124359708-ae092600-dc58-11eb-8bef-9cb0b542772c.png)

GET Single Property

![Image7](https://user-images.githubusercontent.com/1651333/124359714-b6616100-dc58-11eb-97d2-832178a5a47b.png)

GET All Countries

![Image8](https://user-images.githubusercontent.com/1651333/124359719-c1b48c80-dc58-11eb-954b-8d9b23f357a6.png)

PUT Update Property

How to run the REST API application:

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
- Before adding the Property to the DB, we should add country by using the endpoint:

```bash
# Create Country endpoint
# <country_name_here should be changed to either Singapore, Malaysia or Philippines
curl --location --request POST 'http://localhost:3000/api/v1/countries' --header 'Content-Type: application/json'  --data-raw '{ "Name": "<country_name_here>" }'
```

Room for improvement/added:

- Dockerized the app
- CI/CD
- User authentication
- UI
- Data sanitization

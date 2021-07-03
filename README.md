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

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9d83d7a3-04c9-4c0e-b1da-c0c286b69c5a/Screenshot_2021-07-03_at_11.16.49_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9d83d7a3-04c9-4c0e-b1da-c0c286b69c5a/Screenshot_2021-07-03_at_11.16.49_PM.png)

GET All properties (no data yet)

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/58b55120-3362-4a8d-bc88-31124d49e37c/Screenshot_2021-07-03_at_11.18.52_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/58b55120-3362-4a8d-bc88-31124d49e37c/Screenshot_2021-07-03_at_11.18.52_PM.png)

POST Single Property with CountryID (Mode: Sale)

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/69ff47f5-ce4a-47f7-8f2a-c691371a441c/Screenshot_2021-07-03_at_11.20.32_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/69ff47f5-ce4a-47f7-8f2a-c691371a441c/Screenshot_2021-07-03_at_11.20.32_PM.png)

POST Single Property with CountryID (Mode: Rental)

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6633d2c3-17ea-468c-a32c-b9c36204a245/Screenshot_2021-07-03_at_11.21.25_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6633d2c3-17ea-468c-a32c-b9c36204a245/Screenshot_2021-07-03_at_11.21.25_PM.png)

GET All Properties

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7d0f1621-578d-44b4-ab7a-db6f00772480/Screenshot_2021-07-03_at_11.29.54_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7d0f1621-578d-44b4-ab7a-db6f00772480/Screenshot_2021-07-03_at_11.29.54_PM.png)

GET Single Property

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6fe20bf7-e63a-4e00-b8f6-2e61d304f1cf/Screenshot_2021-07-03_at_11.22.20_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6fe20bf7-e63a-4e00-b8f6-2e61d304f1cf/Screenshot_2021-07-03_at_11.22.20_PM.png)

GET All Countries

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/bb58e405-c13b-4df6-867a-4539b61376d0/Screenshot_2021-07-03_at_11.31.46_PM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/bb58e405-c13b-4df6-867a-4539b61376d0/Screenshot_2021-07-03_at_11.31.46_PM.png)

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
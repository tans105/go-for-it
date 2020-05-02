# go-for-it
Simple web-application with Sign-up, Sign In implementation using session, cookie based authentication built on Go(Golang). 
### Setting it up
- Setup golang on local (https://golang.org/doc/install)
- Once setup, clone the repo
- Go to clone directory
- Install 3rd party dependencies
  - `go get golang.org/x/crypto/bcrypt`
  - `go get github.com/jinzhu/gorm/dialects/postgres`
  - `go get github.com/jinzhu/gorm`
  - `go get github.com/satori/go.uuid`
- Run `go build`. This will generate an executable file with name `go-for-it`.
- Run the executable 
    - Linux/Mac `./go-for-it` 
    - Windows `cmd /K "go-for-it.exe"`

### Pre-requisite
- PostgreSQL setup on the machine (https://www.postgresql.org/download/)

###Features 
[v1.0](https://github.com/tans105/go-for-it/releases/tag/v1.0 "Named link title")

- Sign up page ( with basic server-side validations )
- Sign in page  ( with basic server-side validations )
- Logout implementation
- Barebone home page ( to be worked on the next release )
- Cookie & session-based authentication
- Database integration using `Gorm ORM` with Postgres ( can be used with any other vendor )
- Storage of user and session details in the database and its retrieval at time of login
- Password encryption and verification using `bcrypt`

Screenshots

![image](https://user-images.githubusercontent.com/8297056/80859029-45e93980-8c7b-11ea-965e-d07ce709a2b7.png)

![image](https://user-images.githubusercontent.com/8297056/80859033-4b468400-8c7b-11ea-8c96-082ff4ed1e81.png)

![image](https://user-images.githubusercontent.com/8297056/80859045-57cadc80-8c7b-11ea-8e22-fc2989d4dd74.png)

# go-for-it
Demo web-application with Sign-up, Sign In implementation using session, cookie based authentication on Go(Golang).

# Setting it up
- Setup golang on local (https://golang.org/doc/install)
- Once setup, clone the repo
- Go to clone directory and run
 - go get golang.org/x/crypto/bcrypt
 - go get github.com/jinzhu/gorm/dialects/postgres
 - go get github.com/jinzhu/gorm
 - go get github.com/satori/go.uuid
- run `go build`. This will generate an executable.

#Pre-requisite
- PostgreSQL setup on the machine

#Features (https://github.com/tans105/go-for-it/releases/tag/v1.0)
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

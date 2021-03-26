
![home.jpg](https://github.com/gitaleksksks/test-task/blob/main/screenshots/home.jpg)


![add.jpg](https://github.com/gitaleksksks/test-task/blob/main/screenshots/add.jpg)


# How to Use

* Install PostgreSQL on your system. You can skip this step if already installed.
* Clone this repository.
* Execute the `customers.sql` file into your PostgreSQL client. This will import sample database and tables that will be used for this example.
* Modify `main.go` file, line 26. You must configure the PostgreSQL database connection. Change your password used in PostgreSQL.
* Run `go build` in command line in this directory to build the executable file.
* Run `./test-task.exe` in command line in this directory for connecting to database and don't close command line.
* Modify `insert.go` file, line 14. You must configure the PostgreSQL database connection. Change your password used in PostgreSQL.
* Run `go run insert.go` in command line in directory `insert` for insert customers to database.
* Open your web browser and navigate to `http://localhost:8080/`.

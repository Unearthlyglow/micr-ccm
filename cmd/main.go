package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//(This essentially defines a new command-line flag with the name addr, a default value of ":8080" and some short help text explaining what the flag controls.)
	addr := flag.String("addr", ":8080", "HTTP network address")
	// Define a new command-line flag for the MySQL DSN string.
	// dsn := flag.String("dsn", "web:pass@/musicregistry?parseTime=true", "MySQL data source")
	flag.Parse()
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So in this code, that means the addr variable
	// is actually a pointer, and we need to dereference it (i.e. prefix it with
	// the * symbol) before using it. Note that we're using the log.Printf()
	// function to interpolate the address with the log message.

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	mux := http.NewServeMux()

	logger.Info("Well hello there, starting server now!", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)

}

// db, err := openDB(*dsn)
// if err != nil {
// 	logger.Error(err.Error())
// 	os.Exit(1)
// }

// We also defer a call to db.Close(), so that the connection pool is closed before the main() function exits.
// defer db.Close()

// app := &application{
// 	logger: logger,
// }

// 	logger.Info("starting server", "addr", *addr)
// 	log.Printf("starting server on %s", *addr)
// 	err = http.ListenAndServe(*addr, mux)
// 	log.Fatal(err)
// 	os.Exit(1)
// }

// // The openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN.
// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		db.Close()
// 		return nil, err
// 	}

// 	return db, nil
// }

// --------------------------------------------
// OLD WORKING CODE TO GET SOMETHING from GO into the web application, use for reference.
// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func databases(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
// 	w.WriteHeader(http.StatusOK)
// 	test := []string{}
// 	test = append(test, "Hello Now Testing Pictures")
// 	test = append(test, "World")

// 	// json.NewEncoder(w).Encode(test)
// 	err := json.NewEncoder(w).Encode(test)
// 	if err != nil {
// 		// Handle the error, for example, log it or send an error response.
// 		log.Println("Error encoding JSON:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// }

// func main() {

// 	//  mime.AddExtensionType(".js", "application/javascript")

// 	http.Handle("/test", http.HandlerFunc(databases))
// 	http.Handle("/", http.FileServer(http.Dir("static")))
// 	srv := &http.Server{
// 		Addr:         ":8080",
// 		Handler:      nil,              // Set your router/handler here
// 		ReadTimeout:  10 * time.Second, // Set a reasonable read timeout
// 		WriteTimeout: 10 * time.Second, // Set a reasonable write timeout
// 	}

// 	log.Fatal(srv.ListenAndServe())
// }

//Notes for SQL Snippet
// A LOT to go over regarding logging: file:///Users/apple/Downloads/lets-go-professional-package/html/03.02-structured-logging.html
//ALl the notes here come from here: file:///Users/apple/Downloads/lets-go-professional-package/html/04.01-setting-up-mysql.html
//Creating a new DB with UTF8 Encoding: UTF-8 is a character encoding system. It lets you represent characters as ASCII text, while still allowing for international characters, such as Chinese characters
//mysql> CREATE DATABASE musicregistry CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
//Switching to new DB:
//mysql> USE musicregistry;
//(Can I update the fields of a table already created?)
//Creating a Table:
// CREATE TABLE musicregistry (
// 	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
// 	title VARCHAR(100) NOT NULL,
// 	serial TEXT NOT NULL,
// 	created DATETIME NOT NULL
// );
//Creating an index:
//Indexes are special lookup tables that need to be used by the database search engine to speed up data retrieval. An index is simply a reference to data in a table. A database index is similar to the index in the back of a journal. It cannot be viewed by the users and just used to speed up the database access.
//(-- Add an index on the created column.)
//mysql> CREATE INDEX idx_musicregistry_created ON musicregistry(created);
//Adding some Data
// INSERT INTO musicregistry (title, serial, created) VALUES (
// 	'UserOne',
// 	'123456',
// 	UTC_TIMESTAMP()
// );
//From a security standpoint it is better idea to NOT connect to MYSQL's server from the root user from a web application. Time to create a new user.
// CREATE USER 'web'@'localhost';
// GRANT SELECT, INSERT, UPDATE, DELETE ON musicregistry.* TO 'web'@'localhost';
// ALTER USER 'web'@'localhost' IDENTIFIED BY '<PUT IN A UNIQUE PASSWORD>';
// ------
//Installing a DB Driver
//terminal command: go get github.com/go-sql-driver/mysql@v1

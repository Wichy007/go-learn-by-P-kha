// // —------------ http simple —-----------------------
// package main

// import (
//     "fmt"
//     "io"
//     "log"
//     "net/http"
// )

// func FuncIndex(w http.ResponseWriter, r *http.Request) {

//     fmt.Println(r.URL.Query().Get("name"))
//     fmt.Println(r.Method)
//     if r.Method == "GET" {
//    	 io.WriteString(w, "hello, world, this is get method")
//     } else {
//    	 io.WriteString(w, "hello, world, this is other method")
//     }

// }
// func main() {
//     http.HandleFunc("/", FuncIndex)
//     fmt.Println("start server completed")
//     error := http.ListenAndServe(":8080", nil)
//     if error != nil {
//    	 log.Fatal(error)
//     }

// }

// — Go echo
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres"
	DB_HOST     = "host.docker.internal"
)

type Member struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// func init(){

// }

var db *sql.DB

func getData(c echo.Context) error {
	// name := c.QueryParam("name")
	rows, err := db.Query("SELECT id, username, email FROM members")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
	fmt.Println(*db)
	fmt.Println(&db)

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Username, &member.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan members"})
			return err
		}
		members = append(members, member)
	}

	c.JSON(http.StatusOK, members)
	return err

}
func back8() {
	// Back5()
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", getData)
	// e.POST("/", HandleIndex)
	e.Logger.Fatal(e.Start(":1323"))
}

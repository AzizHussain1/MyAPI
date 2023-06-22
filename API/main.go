package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type transaction struct {
	LenderName   string `json:"lender"`
	ReceiverName string `json:"receiver"`
	Date         string `json:"date"`
	Amt          int    `json:"amt"`
}

var transactions = []transaction{}

// --------------------------API--------------------------
func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware())

	router.POST("/transactions", postTransactions)
	router.Run("localhost:9090")
}

// Middleware to enable CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// --------------------------POST REQ--------------------------
func postTransactions(context *gin.Context) {
	var newTxn transaction

	if err := context.BindJSON(&newTxn); err != nil {
		return
	}

	transactions = append(transactions, newTxn)
	context.IndentedJSON(http.StatusCreated, newTxn)

	ExecQuery(newTxn.LenderName, newTxn.ReceiverName, newTxn.Date, newTxn.Amt)
}

// --------------------------BACK-END--------------------------
func ExecQuery(p_lender string, p_receiver string, p_date string, p_amt int) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "hussain"
		dbname   = "transactiondb"
	)

	psql_conn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql_conn)
	CheckError(err)

	defer db.Close()

	query := `INSERT INTO transaction_content(lendername,receivername,date,totalamt)values($1,$2,$3,$4)`
	_, e := db.Exec(query, p_lender, p_receiver, p_date, p_amt)
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//--------------------------BACK-END--------------------------

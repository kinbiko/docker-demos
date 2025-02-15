package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Todo List</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 40px auto; }
        ul { list-style-type: none; padding: 0; }
        li { padding: 8px; border-bottom: 1px solid #ddd; }
        button { margin-left: 10px; }
    </style>
</head>
<body>
    <h1>Todo List</h1>
    <form id="todo-form">
        <input type="text" id="title" placeholder="New Todo" required>
        <button type="submit">Add</button>
    </form>
    <ul id="todo-list">
        {{ range .Todos }}
        <li>
            {{ .Title }} 
            <button onclick="deleteTodo({{ .ID }})">Delete</button>
        </li>
        {{ end }}
    </ul>

    <script>
        document.getElementById("todo-form").addEventListener("submit", async function(e) {
            e.preventDefault();
            const title = document.getElementById("title").value;
            await fetch("/add", { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify({ title }) });
            location.reload();
        });

        async function deleteTodo(id) {
            await fetch("/delete/" + id, { method: "DELETE" });
            location.reload();
        }
    </script>
</body>
</html>
`

func initDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func getTodos(c *gin.Context) {
	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}
	defer rows.Close()

	var todos []struct {
		ID    int
		Title string
	}

	for rows.Next() {
		var todo struct {
			ID    int
			Title string
		}
		if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
			c.String(http.StatusInternalServerError, "Error scanning todos")
			return
		}
		todos = append(todos, todo)
	}

	tmplParsed := template.Must(template.New("webpage").Parse(tmpl))
	tmplParsed.Execute(c.Writer, gin.H{"Todos": todos})
}

func addTodo(c *gin.Context) {
	var todo struct {
		Title string `json:"title"`
	}
	if err := c.BindJSON(&todo); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.Exec("INSERT INTO todos (title) VALUES (?)", todo.Title)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error adding todo")
		return
	}
	c.String(http.StatusOK, "Todo added")
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting todo")
		return
	}
	c.String(http.StatusOK, "Todo deleted")
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/", getTodos)
	r.POST("/add", addTodo)
	r.DELETE("/delete/:id", deleteTodo)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

package Task_3

import (
	"github.com/jmoiron/sqlx"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	ID         int    `DB:"id"`
	Name       string `DB:"name"`
	Department string `DB:"department"`
	Salary     int    `DB:"salary"`
}

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateTable() error {
	schema := `DROP TABLE IF EXISTS employees;
	CREATE TABLE employees (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    department TEXT NOT NULL,
	    salary INTEGER
	);`
	_, err := r.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) FindEmployeesByDepartment(department string) ([]Employee, error) {
	var employees []Employee
	err := r.db.Select(&employees, "SELECT id,name,department,salary FROM employees WHERE department = ?", department)
	return employees, err
}

func (r *EmployeeRepository) FindHighestPaidEmployee() (*Employee, error) {
	var employee Employee
	err := r.db.Get(&employee, "SELECT id,name,department,salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) InsertEmployee(e *Employee) error {
	_, err := r.db.Exec("INSERT INTO employees (name,department,salary) VALUES (?,?,?)", e.Name, e.Department, e.Salary)
	return err
}

type Book struct {
	ID     int     `DB:"id"`
	Title  string  `DB:"title"`
	Author string  `DB:"author"`
	Price  float64 `DB:"price"`
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateTable() error {
	schema := `DROP TABLE IF EXISTS books;
	CREATE TABLE books (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title  TEXT,
		author TEXT,
		price  REAL
	);`
	_, err := r.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) FindBooksByPrice(minPrice float64) ([]Book, error) {
	var books []Book
	err := r.db.Select(&books, "SELECT id,title,author,price FROM books WHERE price > ?", minPrice)
	return books, err
}

func (r *BookRepository) InsertBook(b *Book) error {
	_, err := r.db.Exec("INSERT INTO books (title,author,price) VALUES (?,?,?)", b.Title, b.Author, b.Price)
	return err
}

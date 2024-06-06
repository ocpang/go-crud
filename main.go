package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go-book")
	if err != nil {
		panic(err)
	}
}

type Author struct {
	ID   int
	Name string
}

type Book struct {
	ID       int
	Title    string
	AuthorID int
	Stock    int
	Price    float64
}

// CreateAuthor = menambahkan author baru kedalam database
func CreateAuthor(name string) (int64, error) {
	result, err := db.Exec("INSERT INTO author (name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// ReadAuthors = mengambil semua author dari database
func ReadAuthors() ([]Author, error) {
	rows, err := db.Query("SELECT id, name FROM author")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// UpdateAuthor = mengupdate author yang sudah ada di database
func UpdateAuthor(id int, name string) error {
	_, err := db.Exec("UPDATE author SET name = ? WHERE id = ?", name, id)
	return err
}

// DeleteAuthor = menghapus author dari database
func DeleteAuthor(id int) error {
	_, err := db.Exec("DELETE FROM author WHERE id = ?", id)
	return err
}

// CreateBook = menambahkan buku baru kedalam database
func CreateBook(title string, authorID, stock int, price float64) (int64, error) {
	result, err := db.Exec("INSERT INTO book (title, author_id, stock, price) VALUES (?, ?, ?, ?)", title, authorID, stock, price)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// ReadBooks = mengambil semua buku dari database
func ReadBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author_id, stock, price FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Stock, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// UpdateBook = mengupdate buku yang sudah ada di database
func UpdateBook(id int, title string, authorID, stock int, price float64) error {
	_, err := db.Exec("UPDATE book SET title = ?, author_id = ?, stock = ?, price = ? WHERE id = ?", title, authorID, stock, price, id)
	return err
}

// DeleteBook = menghapus buku dari database
func DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM book WHERE id = ?", id)
	return err
}

func main() {
	authorID, err := CreateAuthor("Author Example")
	if err != nil {
		fmt.Println("Error creating author:", err)
		return
	}
	fmt.Println("Created author with ID:", authorID)

	bookID, err := CreateBook("Book Example", int(authorID), 10, 19.99)
	if err != nil {
		fmt.Println("Error creating book:", err)
		return
	}
	fmt.Println("Created book with ID:", bookID)

	authors, err := ReadAuthors()
	if err != nil {
		fmt.Println("Error reading authors:", err)
		return
	}
	fmt.Println("Authors:", authors)

	books, err := ReadBooks()
	if err != nil {
		fmt.Println("Error reading books:", err)
		return
	}
	fmt.Println("Books:", books)

	err = UpdateAuthor(int(authorID), "Updated Author Example")
	if err != nil {
		fmt.Println("Error updating author:", err)
	}
	fmt.Println("Updated author with ID:", authorID)

	err = UpdateBook(int(bookID), "Updated Book Example", int(authorID), 15, 25.99)
	if err != nil {
		fmt.Println("Error updating book:", err)
	}
	fmt.Println("Updated book with ID:", bookID)

	err = DeleteBook(int(bookID))
	if err != nil {
		fmt.Println("Error deleting book:", err)
	}
	fmt.Println("Deleted book with ID:", bookID)

	err = DeleteAuthor(int(authorID))
	if err != nil {
		fmt.Println("Error deleting author:", err)
	}
	fmt.Println("Deleted author with ID:", authorID)

}

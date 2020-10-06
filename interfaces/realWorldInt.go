package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)
// Notifier
type User struct {
	Name 	string
	Email 	string
	Notifier UserNotifier
}
type EmailNotifier struct {

}
type UserNotifier interface {
	SendMessage(user *User, message string) error
}
type SmsNotifier struct {
	
}
// Notifier
func (notifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending email to %s with content %s\n", user.Name, message)
	return err
}
func (notifier SmsNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending SMS to %s with content %s\n", user.Name, message)
	return err
}
func (user *User) notify(message string) error {
	return user.Notifier.SendMessage(user, message)
}
// Book
type Book struct {
	Title string
	Author string
}
// Customer
type Customer struct {
	Name string
	Age int
}
// Customer
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(js)
	return err
}
// Book
func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
// Book
type Count int
// Book
func (c Count) String() string  {
	return strconv.Itoa(int(c))
}
func WriteLog(s fmt.Stringer)  {
	log.Println(s.String())
}

func main() {
	// Notifier
	user := User{
		"Alex",
		"alexhankey1998@hotmail.co.uk",
		EmailNotifier{},
	}
	user2 := User{
		Name:     "Callum",
		Email:    "callum@hotmail.co.uk",
		Notifier: SmsNotifier{},
	}
	fmt.Printf("Welcome %s\n", user.Name)

	user.notify("Welcome Email user!")
	user2.notify("Welcome SMS user!")

	// Book
	book := Book{"Alice in wonderland", "Lewis Carrol"}
	WriteLog(book)

	count := Count(3)
	WriteLog(count)

	// Customer
	c := &Customer{
		Name: "Alex",
		Age: 21,
	}
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
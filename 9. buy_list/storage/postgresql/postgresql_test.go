package postgresql

import (
	"buy_list/product"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var (
	nickname       string = " "
	name           string = "Жопа"
	user_id        int64  = 424242424
	chat_id        int64  = 2424224242
	created_at            = time.Now()
	finished_at, _        = time.Parse("02-01-2006", "31-12-2025")
)

func TestCreateUser(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.CreateUser(nickname, name, user_id, chat_id)
}
func TestAddIn(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	p := product.Product{User_id: -666, Chat_id: -666}
	db.AddIn(&p)
}
func TestGetStatus(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.GetStatus(user_id, chat_id)
}
func TestGetList(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.GetList(user_id, chat_id, 3)
}
func TestStatusTime(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.StatusTime(user_id, chat_id, created_at, finished_at)
}
func TestSetStatus(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.SetStatus(1, user_id, chat_id)
}
func TestSetFridge(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.SetFridge(user_id, chat_id, created_at, finished_at, name)
}
func TestSetTrash(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.SetTrash(user_id, chat_id, name)
}
func TestSetUsed(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.SetUsed(user_id, chat_id, created_at, finished_at, name)
}
func TestUpdateTimer(t *testing.T) {
	db := Connect(os.Getenv("DB_STRING"))
	db.UpdateTimer(user_id, chat_id)
}

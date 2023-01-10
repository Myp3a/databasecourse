package commands

import (
	"buy_list/storage/postgresql"
	"strconv"
	"strings"
)

func CommandHandler(user_msg string, user_id int64, chat_id int64, user_nickname string, user_name string, db *postgresql.Connection) string {
	msg := "Неизвестная команда."
	switch user_msg {
	case "/start":
	case "/reg":
		res := db.CreateUser(user_nickname, user_name, user_id, chat_id)
		if res {
			msg = "Теперь ты можешь начать работу с ботом! Загляни в команды!"
		} else {
			msg = "Ошибка регистрации."
		}

		return msg
	case "/add":
		msg = "Введите название продукта, вес и дату напоминания через пробел"
		db.SetStatus(1, user_id, chat_id)
		return msg
	case "/buy":
		msg = "введите название продукта и срок хранения через пробел"
		db.SetStatus(2, user_id, chat_id)
		return msg
	case "/open":
		msg = "Введите название продукта, который вы открыли и новый срок хранения"
		db.SetStatus(3, user_id, chat_id)
		return msg
	case "/finish":
		msg = "Введите название продукта, который вы приготовили или выбросили"
		db.SetStatus(4, user_id, chat_id)
		return msg
	case "/list":
		msg = "Введите цифру сортировки списка продуктов: \n1.По алфавиту\n2.По истечению срока годности\n3.Список всех продуктов"
		db.SetStatus(5, user_id, chat_id)
		return msg
	case "/listused":
		msg = "Список ранее использованных продуктов"
		products, _ := db.GetList(user_id, chat_id, 4)
		for i := 0; i < len(products); i++ {
			msg += "\n" + strconv.Itoa(i+1) + ". " + products[i].Name
			if products[i].Weight != 0 {
				msg += ", " + strconv.FormatFloat(products[i].Weight, 'f', 0, 64) + "гр."
			}
			rtime := products[i].Rest_time.String()
			if products[i].Rest_time > 0 {
				rtime = strings.ReplaceAll(rtime, "h", " часов, ")
				rtime = strings.ReplaceAll(rtime, "m", " минут, ")
				rtime = strings.ReplaceAll(rtime, "s", " секунд ")
				msg += ", испортится через: " + rtime
			} else {
				msg += ", срок годности вышел"
			}
		}
		db.SetStatus(0, user_id, chat_id)
		return msg
	case "/statustime":
		msg = "Введите, за какой промежуток промежуток времени вы хотите просмотреть статистику\nПример: 30-12-2022 31-12-2022"
		db.SetStatus(6, user_id, chat_id)
		return msg
	}
	return msg
}

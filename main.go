package main

import (
	"github.com/zenoncoode/golang/internal/entity"
	//_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

func main() {
	/*db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	uc.Execute(input) */

	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}

	car := Car{ //Declara e atribui variavel a car
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Fiat" //alterando o valor do atributo
	println((car.Model))
	car.Start()
	car.ChangeColor("Blue")
	println(car.Color)

}

func (c Car) Start() { //metodo
	println(c.Model + " has been started")
}

func (c Car) ChangeColor(color string) {
	c.Color = color                  //duplicando na memoria o valor de c.color
	println("New color: " + c.Color) // quando o metodo termina de ser executado o valor novo é deletado da memoria
}

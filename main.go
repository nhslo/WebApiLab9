package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
	Category string
}

// Стоимость всех товаров на складе с учетом их количества.
func totalValue(products []Product) float64 {
	total := 0.0

	for _, product := range products {
		productValue := product.Price * float64(product.Quantity)
		total = total + productValue
	}

	return total
}

// Средняя цена одной товарной позиции.
func averagePrice(products []Product) float64 {
	total := 0.0

	for _, product := range products {
		total = total + product.Price
	}

	return total / float64(len(products))
}

// Поиск товара с самой высокой ценой.
func mostExpensive(products []Product) Product {
	expensiveProduct := products[0]

	for _, product := range products {
		if product.Price > expensiveProduct.Price {
			expensiveProduct = product
		}
	}

	return expensiveProduct
}

func main() {
	products := []Product{
		{ID: 1, Name: "Ноутбук", Price: 350000, Quantity: 4, Category: "Электроника"},
		{ID: 2, Name: "Мышь", Price: 5000, Quantity: 15, Category: "Аксессуары"},
		{ID: 3, Name: "Монитор", Price: 85000, Quantity: 3, Category: "Электроника"},
		{ID: 4, Name: "Клавиатура", Price: 12000, Quantity: 8, Category: "Аксессуары"},
		{ID: 5, Name: "Наушники", Price: 25000, Quantity: 2, Category: "Аксессуары"},
	}

	fmt.Println("Список товаров:")
	for _, product := range products {
		fmt.Println(product.ID, product.Name, "—", product.Price, "тг,", product.Quantity, "шт.,", product.Category)
	}

	warehouseValue := totalValue(products)
	meanPrice := averagePrice(products)
	expensiveProduct := mostExpensive(products)

	fmt.Printf("\nОбщая стоимость склада: %.0f тг\n", warehouseValue)
	fmt.Println("Средняя цена товара:", meanPrice, "тг")
	fmt.Println("Самый дорогой товар:", expensiveProduct.Name, "—", expensiveProduct.Price, "тг")

	fmt.Println("\nТовары с низким остатком:")
	for _, product := range products {
		if product.Quantity < 5 {
			fmt.Println(product.Name, "—", product.Quantity, "шт.")
		}
	}

	// Считаем, сколько товарных позиций относится к каждой категории.
	categories := map[string]int{}
	for _, product := range products {
		categories[product.Category] = categories[product.Category] + 1
	}

	fmt.Println("\nКоличество позиций по категориям:")
	for category, count := range categories {
		fmt.Println(category, ":", count)
	}

	fmt.Println("\nСкидка 10%:")
	for _, product := range products {
		if product.Price > 50000 {
			discount := product.Price * 10 / 100
			newPrice := product.Price - discount
			fmt.Println(product.Name, "—", newPrice, "тг")
		}
	}
}

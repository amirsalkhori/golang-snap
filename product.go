package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func manageProduct() {
	fmt.Println("task2 =====================================")
	inventory := make(map[int]Product)
	product1 := Product{ID: 1, Name: "Tea", Quantity: 10, Price: 5.0}
	product2 := Product{ID: 2, Name: "Tea", Quantity: 10, Price: 5.0}
	product3 := Product{ID: 3, Name: "Tea", Quantity: 10, Price: 5.0}
	product4 := Product{ID: 4, Name: "Tea", Quantity: 10, Price: 5.0}
	product5 := Product{ID: 5, Name: "Tea", Quantity: 10, Price: 5.0}

	AddProduct(inventory, product1)
	AddProduct(inventory, product2)
	AddProduct(inventory, product3)
	AddProduct(inventory, product4)
	AddProduct(inventory, product5)

	product6 := Product{ID: 6, Name: "Tea", Quantity: 10, Price: 5.0}
	err := AddProduct(inventory, product6)
	fmt.Printf("add Product %d ,error %v\n\n", product6.ID, err)
	err = AddProduct(inventory, product3)
	fmt.Printf("add Product %d ,error %v\n\n", product3.ID, err)

	err = RemoveProduct(inventory, product6.ID)
	fmt.Printf("remove Product %d ,error %v\n\n", product6.ID, err)

	quantity, exists := CheckStock(inventory, product1.ID)
	if exists {
		fmt.Printf("quantity of product %d is %d\n\n", product1.ID, quantity)
	} else {
		fmt.Printf("not exist product %d\n", product1.ID)
	}

	totalPrice := CalculateTotalValue(inventory)
	fmt.Printf("total product price is = %v\n", totalPrice)
}

// task 2 ======================================================
func AddProduct(inventory map[int]Product, product Product) error {
	if _, exists := inventory[product.ID]; exists {
		return fmt.Errorf("Product with ID %d already exists", product.ID)
	}
	inventory[product.ID] = product
	return nil
}

func RemoveProduct(inventory map[int]Product, productID int) error {
	if _, exists := inventory[productID]; !exists {
		return fmt.Errorf("Product with ID %d does not exists", productID)
	}
	delete(inventory, productID)
	return nil
}

func CheckStock(inventory map[int]Product, productID int) (int, bool) {
	if product, exists := inventory[productID]; exists {
		return product.Quantity, true
	}
	return 0, false
}

func CalculateTotalValue(inventory map[int]Product) float64 {
	var total float64
	for _, product := range inventory {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

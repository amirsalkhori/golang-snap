package main

import "fmt"

type OrderItem struct {
	ProductID int
	Quantity  int
	Price     float64
}

type Order struct {
	ID       int
	Items    []OrderItem
	Discount float64
	Total    float64
	Status   string
}

type DiscountRule struct {
	MinAmount       float64
	DiscountPercent float64
	Description     string
}

func manageOrders() {
	fmt.Println("task3 =====================================")
	// Define discount rules
	rules := []DiscountRule{
		{MinAmount: 100.0, DiscountPercent: 5.0, Description: "5% off for orders over $100"},
		{MinAmount: 500.0, DiscountPercent: 10.0, Description: "10% off for orders over $500"},
		{MinAmount: 1000.0, DiscountPercent: 15.0, Description: "15% off for orders over $1000"},
	}

	order1 := Order{
		ID: 1,
		Items: []OrderItem{
			{ProductID: 1, Quantity: 2, Price: 50.0},
			{ProductID: 2, Quantity: 1, Price: 30.0},
		},
		Status: "pending",
	}

	order2 := Order{
		ID: 2,
		Items: []OrderItem{
			{ProductID: 3, Quantity: 10, Price: 100.0},
		},
		Status: "pending",
	}
	// Create orders
	orders := []Order{order1, order2}

	subtotal := CalculateSubtotal(order1)
	fmt.Printf("subtotal order subtotal is = %v\n", subtotal)

	order1 = ApplyDiscountRules(order1, rules)
	fmt.Printf(" applyDiscountRules = %v\n", order1)

	processedOrders := ProcessOrders(orders, rules)
	fmt.Printf(" order = %v\n", processedOrders)

	completedOrders := FilterOrdersByStatus(processedOrders, "completed")
	fmt.Printf(" completedOrders = %v\n", completedOrders)

	stats := CalculateOrderStatistics(processedOrders)
	fmt.Printf(" status = %v\n", stats)
}

// task 1 ======================================================
func SumNumbers(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// task 3 ======================================================
func CalculateSubtotal(order Order) float64 {
	var total float64
	for _, orderItem := range order.Items {
		total += float64(orderItem.Quantity) * orderItem.Price
	}
	return total
}

func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
	subtotal := CalculateSubtotal(order)

	var maxDiscountAmount float64
	for _, rule := range rules {
		if subtotal >= rule.MinAmount {
			discountAmount := subtotal * rule.DiscountPercent / 100
			if discountAmount > maxDiscountAmount {
				maxDiscountAmount = discountAmount
			}
		}
	}

	order.Discount = maxDiscountAmount
	order.Total = subtotal

	return order
}

func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
	//Chech nil safe pointer for the both orders and rules
	
	for i := range orders {
		order := &orders[i]

		//Use const
		order.Status = "processing"
		*order = ApplyDiscountRules(*order, rules)
		//Use const
		order.Status = "completed"
	}
	return orders
}

//Instead of use status as a string, use const
func FilterOrdersByStatus(orders []Order, status string) []Order {
	//Chech nil safe pointer for the the orders
	var filteredOrders []Order
	for _, order := range orders {
		if order.Status == status {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func CalculateOrderStatistics(orders []Order) map[string]interface{} {
		//Chech nil safe pointer for the the orders
	orderStatus := make(map[string]interface{})

	totalOrders := len(orders)
	completedOrders := 0
	totalRevenue := 0.0
	totalDiscount := 0.0
	for _, order := range orders {
			//Use const
		if order.Status == "completed" {
			completedOrders++
			totalRevenue += order.Total - order.Discount
			totalDiscount += order.Discount
		}
	}

	averageOrderValue := 0.0
	if totalOrders > 0 {
		averageOrderValue = totalRevenue / float64(totalOrders)
	}

	orderStatus["total_orders"] = totalOrders
	orderStatus["total_revenue"] = totalRevenue
	orderStatus["average_order_value"] = averageOrderValue
	orderStatus["completed_orders"] = completedOrders
	orderStatus["total_discount"] = totalDiscount

	return orderStatus
}

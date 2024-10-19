package main

import (
	"fmt"
	"os"
)

const (
	FrontRowTicketPrice = 10
	BackRowTicketPrice  = 8
)

var (
	ticketSold, totalSeats, income, rows, seats int
)

// printHeader prints the cinema layout header
func printHeader(seats int) {

	fmt.Println("\nCinema:")
	fmt.Print(" ") // Leading space for alignment with row numbers

	for i := 1; i <= seats; i++ {
		fmt.Print(" ", i) // Gives the initial row
	}
}

// printLayout gives the seating layout
func printLayout(rows, seats int, layout [][]string) {
	printHeader(seats) // Print seat header
	for i := 0; i < rows; i++ {
		fmt.Print("\n", i+1) // Print row number
		for j := 0; j < seats; j++ {
			fmt.Print(" ", layout[i][j]) // Print each seat with a space
		}
	}
	fmt.Println()
}

// createLayout creates the seating layout
func createLayout() (int, int, [][]string) {
	//var rows, seats int
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)

	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	totalSeats = rows * seats

	layout := make([][]string, rows)
	for i := 0; i < rows; i++ {
		layout[i] = make([]string, seats)
		for j := 0; j < seats; j++ {
			layout[i][j] = "S" // Initialize each seat with "S"
		}
	}
	return rows, seats, layout
}

// showMenu displays the menu to the users
func showMenu() int {
	var option int
	fmt.Println(`
1. Show the seats
2. Buy a ticket
3. Statistics
0. Exit`)
	fmt.Scan(&option)
	return option
}

// getSeat returns the row and seat number based on user's inputs
func getSeat() (int, int) {
	var rowNumber, seatNumber int

	fmt.Println("\nEnter a row number:")
	fmt.Scan(&rowNumber)

	fmt.Println("Enter a seat number in that row:")
	fmt.Scan(&seatNumber)

	return rowNumber, seatNumber
}

// buyTicket purchases the ticket based on chosen seat
func buyTicket(rows int, seats int, layout [][]string) {
	for {
		rowNumber, seatNumber := getSeat()

		if (rowNumber > rows || seatNumber > seats) || (rowNumber <= 0 || seatNumber <= 0) {
			fmt.Println("\nWrong input!")
			continue
		}

		seatSelection := layout[rowNumber-1][seatNumber-1]

		if seatSelection == "B" {
			fmt.Println("\nThat ticket has already been purchased!")
			continue
		}

		price := ticketPrice(totalSeats, rows, rowNumber)
		fmt.Printf("Ticket price: $%d\n", price)

		layout[rowNumber-1][seatNumber-1] = "B"
		ticketSold += 1
		income += price
		break
	}
}

// ticketPrice takes the total seats and calculates the ticket price based on seat location
func ticketPrice(totalSeats, rows, rowNumber int) int {
	frontRow := rows / 2 // To identify where the seat falls into, front or back
	if totalSeats > 60 && rowNumber > frontRow {
		return BackRowTicketPrice
	}
	return FrontRowTicketPrice
}

func totalSeatsIncome() int {
	if totalSeats > 60 {
		frontRows := rows / 2
		backRows := rows - frontRows
		total := (frontRows * seats * FrontRowTicketPrice) + (backRows * seats * BackRowTicketPrice)
		return total
	}
	return rows * seats * 10
}

// printStats print the details of tickets sold, income, etc.,
func printStats() {

	fmt.Printf("\nNumber of Purchased tickets: %d", ticketSold)
	fmt.Printf("\nPercentage: %.2f%%", float32(ticketSold)/float32(totalSeats)*float32(100))
	fmt.Printf("\nCurrent income: $%d", income)
	fmt.Printf("\nTotal income: $%d\n", totalSeatsIncome())
}

func main() {
	//var rows, seats int
	rows, seats, layout := createLayout()
	for {
		option := showMenu()
		switch option {
		case 1:
			printLayout(rows, seats, layout)
		case 2:
			buyTicket(rows, seats, layout)
		case 3:
			printStats()
		case 0:
			os.Exit(0)
		}
	}
}

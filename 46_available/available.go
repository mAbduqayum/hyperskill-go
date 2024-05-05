package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Order struct {
	index   int
	arrival int
}

type Truck struct {
	start, end, capacity, index int
}

func assignOrderToTruck(orders []Order, trucks []Truck) []int {
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].arrival < orders[j].arrival
	})

	sort.Slice(trucks, func(i, j int) bool {
		return trucks[i].start < trucks[j].start
	})

	type TruckOrder struct {
		index int
		truck Truck
	}

	var result []int
	for _, order := range orders {
		assigned := false

		var resTrucks []TruckOrder

		viewed := make(map[int]struct{})
		dublicate := false

		for j, truck := range trucks {
			if truck.capacity > 0 && order.arrival >= truck.start && order.arrival <= truck.end {
				resTrucks = append(resTrucks, TruckOrder{index: j, truck: truck})

				if _, ok := viewed[truck.start]; ok {
					dublicate = true
				}

				viewed[truck.start] = struct{}{}

				//result = append(result, truck.index+1)
				//trucks[j].capacity--
				//
				//// if capacity is 0, remove the truck from the list
				//if trucks[j].capacity == 0 {
				// trucks = append(trucks[:j], trucks[j+1:]...)
				//}

				assigned = true
			}
		}

		if !assigned {
			result = append(result, -1)
			continue
		}

		if len(resTrucks) == 0 {
			result = append(result, -1)
			continue
		}

		fmt.Println(resTrucks)
		fmt.Println(dublicate)

		if dublicate == false {
			t := resTrucks[0]

			fmt.Println(result)
			result = append(result, t.truck.index+1)
			fmt.Println(result)
			trucks[t.index].capacity--

			continue
		}

		sort.Slice(resTrucks, func(i, j int) bool {
			return resTrucks[i].truck.index < resTrucks[j].truck.index
		})

		t := resTrucks[0]

		result = append(result, t.truck.index+1)
		trucks[t.index].capacity--
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		arrivals := strings.Split(scanner.Text(), " ")
		orders := make([]Order, n)
		for j, a := range arrivals {
			arrival, _ := strconv.Atoi(a)
			orders[j] = Order{index: j, arrival: arrival}
		}

		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())

		trucks := make([]Truck, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			vals := strings.Split(scanner.Text(), " ")
			trucks[j].start, _ = strconv.Atoi(vals[0])
			trucks[j].end, _ = strconv.Atoi(vals[1])
			trucks[j].capacity, _ = strconv.Atoi(vals[2])
			trucks[j].index = j
		}

		result := assignOrderToTruck(orders, trucks)
		originalOrder := make([]int, n)
		for i, order := range orders {
			originalOrder[order.index] = result[i]
		}
		fmt.Println(strings.Trim(fmt.Sprint(originalOrder), "[]"))
	}
}

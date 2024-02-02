// [Seat{Label: "1A", Weight: 1.0}, Seat{Label: "1B", Weight: 1.4}, Seat{Label: "1C", Weight: 3 ]
// [Seat{Label: "2A", Weight: 0.5}, Seat{Label: "2B", Weight: 1.4}, Seat{Label: "2C", Weight: 3 ]
// [Seat{Label: "3A", Weight: 0.5}, Seat{Label: "3B", Weight: 1.4}, Seat{Label: "3C", Weight: 3 ]
// [Seat{Label: "4A", Weight: 0.5}, Seat{Label: "4B", Weight: 1.4}, Seat{Label: "4C", Weight: 3 ]
// [Seat{Label: "5A", Weight: 0.5}, Seat{Label: "5B", Weight: 1.4}, Seat{Label: "5C", Weight: 3 ]
// [Seat{Label: "6A", Weight: 0.5}, Seat{Label: "6B", Weight: 1.4}, Seat{Label: "6C", Weight: 3 ]
// [Seat{Label: "7A", Weight: 0.5}, Seat{Label: "7B", Weight: 1.4}, Seat{Label: "7C", Weight: 3 ]

// Balance: we have balance when weight is distributed across the plane with good propertions symetrically.

// [0, 1, 0] [0, 1, 0] [0, 1, 0]
// [1, 1, 0] [0, 1, 0] [1, 1, 1]
// [0, 1, 1] [0, 1, 0] [1, 1, 1]
// [0, 1, 0] [0, 1, 0] [0, 1, 0]
// Key observation: when seating a passenger we strive to balance the other side, by symetrically pairing one more passenger
// 1: we should try to fit all passengers in the middle row, starting from the center, in that case plane weight distribution is achieved
// 2: if we can't fit all passengers in the middle, we start spreading them around the middle

package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/impzero/plane/calc"
)

func main() {
	plane := NewPlane("Boeing 747", 10, 6)
	// plane.ManualAssign(3, 4, 2)
	// plane.ManualAssign(29, 4, 4)
	// plane.ManualAssign(4, 3, 7)
	// plane.ManualAssign(12, 1, 13)
	// plane.ManualAssign(22, 6, 23)
	// plane.ManualAssign(2, 1, 99)
	// plane.ManualAssign(30, 6, 9)
	// plane.ManualAssign(1, 1, 33)

	plane.ManualAssign(1, 1, 5)
	plane.ManualAssign(1, 6, 5)
	// plane.ManualAssign(10, 1, 5)
	// plane.ManualAssign(10, 6, 5)
	// plane.ManualAssign(10, 5, 30)

	plane.Print()
	spew.Dump(plane.CalculateBalanceVector())
	spew.Dump(plane.IsBalanced())
}

type Plane struct {
	Name  string
	Seats [][]Seat

	yBalanceTolerance float64
	xBalanceTolerance float64
}

type Seat struct {
	Label     string
	Weight    float64
	Available bool
}

func NewSeat(row, column int) Seat {
	return Seat{
		Label:     strings.ToUpper(fmt.Sprintf("%s%s", strconv.Itoa(row), string(rune('a'+column)))),
		Weight:    0,
		Available: true,
	}
}

func NewPlane(name string, rows, columns int) Plane {
	plane := Plane{
		Name:  name,
		Seats: [][]Seat{},
	}
	for i := 0; i < rows; i++ {
		plane.Seats = append(plane.Seats, []Seat{})
		for j := 0; j < columns; j++ {
			plane.Seats[i] = append(plane.Seats[i], NewSeat(i+1, j))
		}
	}

	plane.yBalanceTolerance = float64(rows / 2)
	plane.xBalanceTolerance = float64(columns / 2)
	return plane
}

// IsBalanced checks for balance on the plane using the Lever's Law (f1*l1 = f2*l2)

func (p *Plane) CalculateBalanceVector() [2]float64 {
	length := (len(p.Seats) / 2)
	width := (len(p.Seats[0]) / 2)

	var vectors [][2]float64 = make([][2]float64, 0)
	for i, row := range p.Seats {
		for j, seat := range row {
			yDirection := -1.0
			if i < length {
				yDirection = 1
			}

			// if plane has even number of rows there are TWO rows, considered center
			if length%2 == 1 {
				if i == length || i+1 == length {
					yDirection = 0
				}
			}

			xDirection := 1.0
			if j < width {
				xDirection = -1
			}

			// if plane has even number of seats per row there are TWO seats, considered center
			if width%2 == 1 {
				if j == width || j+1 == width {
					xDirection = 0
				}
			}

			seatOffsetX := calc.GetCenterOffset(len(p.Seats[i]), j)
			seatOffsetY := calc.GetCenterOffset(len(p.Seats), i)

			xCoordinate := xDirection * (seat.Weight * float64(seatOffsetX))
			yCoordinate := yDirection * (seat.Weight * float64(seatOffsetY))

			vectors = append(vectors, [2]float64{xCoordinate, yCoordinate})
		}
	}

	return calc.GetVector(vectors)
}

func (p *Plane) IsBalanced() bool {
	point := p.CalculateBalanceVector()
	x := point[0]
	y := point[0]

	width := float64(len(p.Seats[0])) / 2.0
	length := float64(len(p.Seats)) / 2.0

	if math.Abs(x/width) > p.xBalanceTolerance {
		return false
	}
	if math.Abs(y/length) > p.yBalanceTolerance {
		return false
	}
	return true
}

func (p *Plane) ManualAssign(rowVal, columnVal int, weight float64) (Seat, error) {
	for y, row := range p.Seats {
		for x, seat := range row {
			if y == (rowVal-1) && x == columnVal-1 {
				if seat.Available {
					p.Seats[y][x].Weight = weight
					p.Seats[y][x].Available = false
					return p.Seats[y][x], nil
				} else {
					return Seat{}, errors.New("ManualAssign: cannot pick this seat")
				}
			}
		}
	}
	return Seat{}, errors.New("ManualAssign: seat not found")
}

func (p *Plane) AutoAssign(weight float64) Seat {
	// if we have a 2 balls one with weight of 1kg and 2 meters away from the center
	// and one with 2kg of weight but 1 meter away from the center
	// we achieve balance
	// formula: 2kg * 1m = 1kg * 2m
	// if 2kg * 1m and 4kg -> distance = 2kg * 1m = 4kg * x
	// 2kgm = 4kg*x // divide by 4kg
	// 0.5m = x -> x = 0.5m
	// planeRows := len(p.Seats)
	// planeColumns := len(p.Seats[0])

	// centerSeat := p.Seats[planeRows/2][planeColumns/2]
	return Seat{}
}

func (p *Plane) mirror(s Seat) (Seat, error) {
	for y, row := range p.Seats {
		for x, seat := range row {
			if seat.Label == s.Label {
				mirroredRow := len(p.Seats) - y
				mirroredColumn := len(p.Seats[0]) - x
				return p.Seats[mirroredRow][mirroredColumn], nil
			}
		}
	}
	return Seat{}, errors.New("mirror: no symetrical seat found")
}

func (p *Plane) Print() {
	for _, row := range p.Seats {
		fmt.Println()
		for _, seat := range row {
			indicator := "🟥"
			if seat.Available {
				indicator = "🟩"
			}
			if len(seat.Label) < 3 {
				fmt.Printf(" 0%s %s ", seat.Label, indicator)
			} else {
				fmt.Printf(" %s %s ", seat.Label, indicator)
			}
		}
		fmt.Println()
	}
}

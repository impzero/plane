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

	plane.ManualAssign(1, 1, 5)
	plane.ManualAssign(1, 6, 5)

	plane.Print()
	spew.Dump(plane.CalculateBalanceVector())
	spew.Dump(plane.Balanced())
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

func (p *Plane) CalculateBalanceVector() [2]float64 {
	length := (len(p.Seats) / 2)
	width := (len(p.Seats[0]) / 2)

	var vectors [][2]float64 = make([][2]float64, 0)
	for i, row := range p.Seats {
		for j, seat := range row {
			xDirection := 1.0
			yDirection := -1.0
			if i < length {
				yDirection = 1
			}
			if length%2 == 1 {
				if i == length || i+1 == length {
					yDirection = 0
				}
			}
			if j < width {
				xDirection = -1
			}
			if width%2 == 1 {
				if j == width || j+1 == width {
					xDirection = 0
				}
			}
			seatOffsetX := calc.CenterOffset(len(p.Seats[i]), j)
			seatOffsetY := calc.CenterOffset(len(p.Seats), i)
			xCoordinate := xDirection * (seat.Weight * float64(seatOffsetX))
			yCoordinate := yDirection * (seat.Weight * float64(seatOffsetY))

			vectors = append(vectors, [2]float64{xCoordinate, yCoordinate})
		}
	}
	return calc.Vector(vectors)
}

func (p *Plane) Balanced() bool {
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

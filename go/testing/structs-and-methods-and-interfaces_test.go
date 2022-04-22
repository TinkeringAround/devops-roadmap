package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{10.0, 10.0}, 40.0},
		{"Circle", Circle{10.0}, 2 * 10.0 * math.Pi},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name+"-Perimeter", func(t *testing.T) {
			got := tt.shape.Perimeter()

			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, math.Pi * 10.0 * 10.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name+"-Area", func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}

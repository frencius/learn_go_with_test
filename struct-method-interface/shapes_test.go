package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// versi 1
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}

	// versi 2
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f, but want %.2f", got, want)
	// 	}
	// }
	// t.Run("rectangle area", func(t *testing.T) {
	// 	rectangle := Rectangle{6.0, 12.0}

	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{10}

	// 	checkArea(t, circle, 314.1592653589793)
	// })

	// versi 1
	// t.Run("rectangle area", func(t *testing.T) {
	// 	rectangle := Rectangle{6.0, 12.0}
	// 	got := Area(rectangle)
	// 	want := 72.0
	// 	if got != want {
	// 		t.Errorf("got %.2f, but want %.2f", got, want)
	// 	}
	// })

	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{6.0, 12.0}
	// 	got := Area(circle)
	// 	want := 72.0
	// 	if got != want {
	// 		t.Errorf("got %.2f, but want %.2f", got, want)
	// 	}
	// })
}

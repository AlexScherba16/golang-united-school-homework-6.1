package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// isOutOfRange check index range
// returns true if index in range
func (b *box) isOutOfRange(i int) bool {
	if i < 0 || i > (len(b.shapes)-1) {
		return true
	}
	return false
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("limit reached")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.isOutOfRange(i) == true {
		return nil, errors.New("index out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.isOutOfRange(i) == true {
		return nil, errors.New("index out of range")
	}
	result := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return result, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.isOutOfRange(i) == true {
		return nil, errors.New("index out of range")
	}
	result := b.shapes[i]
	b.shapes[i] = shape
	return result, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var perimeter float64
	for _, shape := range b.shapes {
		perimeter += shape.CalcPerimeter()
	}
	return perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64
	for _, shape := range b.shapes {
		area += shape.CalcArea()
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	error := errors.New("no circles")

	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(*Circle)
		if ok {
			error = nil
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		}
	}
	return error
}

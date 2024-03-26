package main

import "fmt"

type Color [3]byte

func (c *Color) Print() {
	fmt.Printf("Red: %v, Green: %v, Blue: %v\n", c[0], c[1], c[2])
}

func (c *Color) GetR() byte {
	return c[0]
}

func (c *Color) GetG() byte {
	return c[1]
}

func (c *Color) GetB() byte {
	return c[2]
}

func (c *Color) SetR(r byte) {
	c[0] = r
}

func (c *Color) SetG(g byte) {
	c[1] = g
}

func (c *Color) SetB(b byte) {
	c[2] = b
}

func (c *Color) GetBrightness() float64 {
	return 0.2126*float64(c[0]) + 0.7152*float64(c[1]) + 0.0722*float64(c[2])
}

func maxBrightness(colors []Color) *Color {
	if colors == nil {
		return nil
	}

	maxBrightness := colors[0]

	for _, c := range colors[1:] {
		if c.GetBrightness() > maxBrightness.GetBrightness() {
			maxBrightness = c
		}
	}
	return &maxBrightness
}

func main() {
	//По итогу решил использовать указатели везде, где
	//есть соответствующие методы.
	//Не понимаю логики, для легковесных структур или типов вообще,
	//но после прочтения многих статей и изучения
	//репозиториев крупных продуктов, понял, что большинство (90%+)
	//пишут без смешивания, доверюсь этим людям что-ли

	c := make([]Color, 3)
	c[0] = Color{105, 20, 15}
	c[0].Print()
	fmt.Printf("Reg brightness: %v\n", c[0].GetR())
	c[0].SetR(90)
	fmt.Printf("The brightness of %v: %v\n", c[0], c[0].GetBrightness())
	c[0].Print()
	c[1] = Color{17, 10, 82}
	fmt.Printf("Max brightness now: %v\n", maxBrightness(c))
	c[2] = Color{c[1].GetR(), 128, 128}
	fmt.Printf("Max brightness then: %v\n", maxBrightness(c))
}

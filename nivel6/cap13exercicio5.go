/*- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"*/

package main

import "fmt"

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Area do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	resultado := 2 * c.raio * 3.14
	fmt.Println("Area do circulo:", resultado)
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()

}

func main() {
	x := quadrado{
		lado: 20.0,
	}
	//x.area()
	info(x)
	y := circulo{
		raio: 20.0,
	}
	//y.area()
	info(y)
}

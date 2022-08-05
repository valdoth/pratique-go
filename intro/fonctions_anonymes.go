package main

func main() {
	println("Les fonctions anonymes")

	w := func() {
		println("Je suis une fonction anonyme sans return.")
	}
	w()

	z := func() string {
		println("Je suis une fonction anonyme.")
		return "Alex"
	}()
	println(z)

	x, y := func() (int, int) {
		println("Auncu paramétre. Deux returns")
		return 5, 7
	}()
	println(x, y)

	func(a, b int) {
		println("A * A + B * B =", a*a+b*b)
	}(x, y)
}



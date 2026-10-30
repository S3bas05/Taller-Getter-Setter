package main

import "fmt"

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) GetID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) GetPrecio() float64 {
	return p.precio
}

func (p Producto) GetCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	if nuevoPrecio >= 0 {
		p.precio = nuevoPrecio
	}
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	if nuevaCantidad >= 0 {
		p.cantidad = nuevaCantidad
	}
}

func (p Producto) MostrarInfo() {
	fmt.Printf(
		"ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n",
		p.id, p.nombre, p.precio, p.cantidad,
	)
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: make(map[int]Producto),
	}
}

func (i *Inventario) AgregarProducto(p Producto) {
	id := p.GetID()

	if _, existe := i.productos[id]; existe {
		fmt.Println("Error: ya existe un producto con ese ID.")
		return
	}

	i.productos[id] = p
	fmt.Println("Producto agregado correctamente.")
}

func (i *Inventario) EliminarProducto(id int) {
	if _, existe := i.productos[id]; !existe {
		fmt.Println("Error: producto no encontrado.")
		return
	}

	delete(i.productos, id)
	fmt.Println("Producto eliminado correctamente.")
}

func (i *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := i.productos[id]
	return p, existe
}

func (i *Inventario) ListarProductos() {
	if len(i.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	fmt.Println("\n--- LISTA DE PRODUCTOS ---")

	for _, p := range i.productos {
		p.MostrarInfo()
	}
}

func main() {
	inventario := NuevoInventario()

	p1 := NuevoProducto(1, "Arroz", 1.25, 50)
	p2 := NuevoProducto(2, "Leche", 0.95, 30)
	p3 := NuevoProducto(3, "Pan", 0.50, 20)

	inventario.AgregarProducto(p1)
	inventario.AgregarProducto(p2)
	inventario.AgregarProducto(p3)

	fmt.Println("\n--- CONSULTA DE UN PRODUCTO ---")
	fmt.Println("Nombre:", p1.GetNombre())
	fmt.Println("Precio:", p1.GetPrecio())
	fmt.Println("Cantidad:", p1.GetCantidad())

	p1.SetPrecio(1.50)
	p1.SetCantidad(45)

	inventario.AgregarProducto(p1)

	fmt.Println("\n--- BÚSQUEDA ---")
	producto, existe := inventario.BuscarProducto(2)

	if existe {
		producto.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado.")
	}

	inventario.ListarProductos()

	fmt.Println("\n--- ELIMINACIÓN ---")
	inventario.EliminarProducto(3)

	inventario.ListarProductos()
}

package shipping_box

type Product struct {
	Name string
	Parameter
}

type Box struct {
	Parameter
}

func Volume(parameter Parameter) int {
	return parameter.Width * parameter.Length * parameter.Height
}

func ProductFitsInBox(product Product, box Box) bool {
	if product.Width > box.Width || product.Length > box.Length || product.Height > box.Height {
		return false
	}
	return true
}

type Parameter struct {
	Length int
	Width  int
	Height int
}

func GetBoxes() []Box {
	boxes := make([]Box, 0)
	p := []Parameter{
		{
			Length: 10,
			Width:  10,
			Height: 20,
		},
		{
			Length: 10,
			Width:  20,
			Height: 20,
		},
		{
			Length: 15,
			Width:  20,
			Height: 25,
		},
		{
			Length: 15,
			Width:  30,
			Height: 50,
		},
		{
			Length: 30,
			Width:  30,
			Height: 60,
		},
		{
			Length: 40,
			Width:  40,
			Height: 40,
		},
		{
			Length: 50,
			Width:  40,
			Height: 45,
		},
		{
			Length: 60,
			Width:  60,
			Height: 50,
		},
	}
	for i := range p {
		boxes = append(boxes, Box{p[i]})
	}
	return boxes
}

func GetProducts() []Product {
	products := make([]Product, 0)
	products = append(products, Product{
		Name: "Test",
		Parameter: Parameter{
			Length: 5,
			Width:  5,
			Height: 5,
		},
	})
	products = append(products, Product{
		Name: "Test",
		Parameter: Parameter{
			Length: 10,
			Width:  5,
			Height: 10,
		},
	})
	products = append(products, Product{
		Name: "Test",
		Parameter: Parameter{
			Length: 10,
			Width:  10,
			Height: 10,
		},
	})
	products = append(products, Product{
		Name: "Test",
		Parameter: Parameter{
			Length: 10,
			Width:  15,
			Height: 20,
		},
	})
	products = append(products, Product{
		Name: "Test",
		Parameter: Parameter{
			Length: 20,
			Width:  15,
			Height: 30,
		},
	})
	return products
}

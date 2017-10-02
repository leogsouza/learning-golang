// Example of Interface with Type Embedding and Method Overriding In GO

package main

import "fmt"

/*
The interface type Information is a contract for creating various Product Types
in a catalog. The Information interface provides three behaviors in its
contract: General, Attributes and Inventory
*/
type Information interface {
	General()
	Attributes()
	Inventory()
}

/*
A struct product is declared with fields for holding its state and methods
implemented based on the behaviors defined in the Information interface */
type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (prd Product) General() {
	fmt.Printf("\n%s\n", prd.Name)
	fmt.Printf("\n%s\b", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

// A struct Mobile is declared in which the type Product is embedded.
type Mobile struct {
	Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

func (mob Mobile) Attributes() {
	mob.Product.Attributes()

	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}

	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}

// A struct Shirts is declared in which the type Product is embedded.
type Shirts struct {
	Product
	Size, Pattern, Sleeve, Fabric string
}

// Overrides for the Attributes Methos for the Shirts struct
func (shrt Shirts) Attributes() {
	fmt.Println("\nSpecification: ")
	fmt.Println(shrt.Size)
	fmt.Println(shrt.Pattern)
	fmt.Println(shrt.Sleeve)
	fmt.Println(shrt.Fabric)
}

func (prd Product) Inventory() {
	fmt.Println(prd.Stock)
}

// A struct named Catalog is declared to represent catalog of various products.
// The type of Details field uses a slice of the Information interface
type Catalog struct {
	Name    string
	Details []Information
}

func (c Catalog) CatalogDetails() {

	fmt.Println("***************************************************\n")
	fmt.Printf("\nStore: %s \n", c.Name)

	for _, key := range c.Details {
		fmt.Println("======================Products======================")
		key.General()
		key.Attributes()
		key.Inventory()
		fmt.Println("====================================================")
	}
}

func main() {
	// Create an Instance of Mobile Type and Call methods
	mobilePrd := Mobile{
		Product{
			"Apple Iphone X (Rose Gold, 128GB) (4GB RAM)",
			"Handset, Apple EarPods with Remote and Mic, Lightning to USB Cable",
			40999, 143, 10,
		},
		[]string{"4GB RAM", "5.2 inch Retina HD Display", "No Home Button"},
		[]string{"128GB ROM", "Wireless charger", "Speaker louder"},
	}

	// Create an Instance of Shirts type and Call methods
	shirtPrd := Shirts{
		Product{
			"Reebok Striped Men's Round Neck Grey T-Shirt",
			"Machine Wash as per Tag, Do not Dry Clean, Do not Bleach",
			1124, 400, 25,
		},
		"XL", "Striped", "Half Sleeve", "Cotton",
	}

	// Create an instance of Catalog Type
	category := Catalog{
		"Amazon",
		[]Information{mobilePrd, shirtPrd},
	}

	category.CatalogDetails()
}

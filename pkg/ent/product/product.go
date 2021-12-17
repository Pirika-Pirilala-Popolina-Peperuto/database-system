// Code generated by entc, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPictureURL holds the string denoting the picture_url field in the database.
	FieldPictureURL = "picture_url"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgeShoppingCartOwners holds the string denoting the shopping_cart_owners edge name in mutations.
	EdgeShoppingCartOwners = "shopping_cart_owners"
	// Table holds the table name of the product in the database.
	Table = "products"
	// OrdersTable is the table that holds the orders relation/edge. The primary key declared below.
	OrdersTable = "order_products"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// CategoriesTable is the table that holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "product_categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// ShoppingCartOwnersTable is the table that holds the shopping_cart_owners relation/edge. The primary key declared below.
	ShoppingCartOwnersTable = "shopping_cart"
	// ShoppingCartOwnersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ShoppingCartOwnersInverseTable = "users"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPrice,
	FieldQuantity,
	FieldPictureURL,
}

var (
	// OrdersPrimaryKey and OrdersColumn2 are the table columns denoting the
	// primary key for the orders relation (M2M).
	OrdersPrimaryKey = []string{"order_id", "product_id"}
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"product_id", "category_id"}
	// ShoppingCartOwnersPrimaryKey and ShoppingCartOwnersColumn2 are the table columns denoting the
	// primary key for the shopping_cart_owners relation (M2M).
	ShoppingCartOwnersPrimaryKey = []string{"user_id", "product_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
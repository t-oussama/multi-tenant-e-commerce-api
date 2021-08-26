package constants

type CollectionNamesRegistry struct {
	PRODUCTS              string
	CATEGORIES            string
	STORES                string
	SUB_CATEGORY_PRODUCTS string
	META_DATA             string
	PRODUCT_LISTS         string
}

var CollectionNames CollectionNamesRegistry

func init() {
	CollectionNames = CollectionNamesRegistry{
		PRODUCTS:              "products",
		CATEGORIES:            "categories",
		STORES:                "stores",
		SUB_CATEGORY_PRODUCTS: "subCategoryProducts",
		META_DATA:             "metaData",
		PRODUCT_LISTS:         "productLists",
	}
}

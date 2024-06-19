package inmemory

import (
	"ramenGO/internal/entities"
	errormessage "ramenGO/pkg/error-message"
)

type dataBaseInMemoryRepository struct {
	broths        []*entities.Broth
	proteins      []*entities.Protein
	orderImage    map[string]string
	KeyGetOrderID *string
}

func NewDataBaseInMemoryRepository(key *string) *dataBaseInMemoryRepository {
	return &dataBaseInMemoryRepository{
		broths:        make([]*entities.Broth, 0),
		proteins:      make([]*entities.Protein, 0),
		orderImage:    make(map[string]string),
		KeyGetOrderID: key,
	}
}

func (r *dataBaseInMemoryRepository) GetBroths() ([]*entities.Broth, *errormessage.ErrorResponse) {
	return r.broths, nil
}

func (r *dataBaseInMemoryRepository) GetProteins() ([]*entities.Protein, *errormessage.ErrorResponse) {
	return r.proteins, nil
}

func (r *dataBaseInMemoryRepository) GetBrothsID() ([]string, error) {
	var brothsIDs []string

	for _, broth := range r.broths {
		brothsIDs = append(brothsIDs, broth.ID)
	}

	return brothsIDs, nil
}

func (r *dataBaseInMemoryRepository) GetProductsList() (*entities.ProductList, *errormessage.ErrorResponse) {
	var productList entities.ProductList

	productList.BrothsIDs = make(map[string]entities.Product)
	productList.ProteinsIDs = make(map[string]entities.Product)

	for _, broth := range r.broths {
		productList.BrothsIDs[broth.ID] = entities.Product{Status: true, Name: broth.Name}
	}

	for _, protein := range r.proteins {
		productList.ProteinsIDs[protein.ID] = entities.Product{Status: true, Name: protein.Name}
	}

	return &productList, nil
}

func (r *dataBaseInMemoryRepository) GetAccessKey() string {
	return *r.KeyGetOrderID
}

func (r *dataBaseInMemoryRepository) GetOrderImage(orderProtein string) (string, *errormessage.ErrorResponse) {
	return r.orderImage[orderProtein], nil
}

func (r *dataBaseInMemoryRepository) PopulateInMemoryRepository() {

	broth_1 := entities.Broth{
		ID:               "1",
		ImageInactiveURL: "https://i.imgur.com/9RqsFbh.png",
		ImageActiveURL:   "https://i.imgur.com/rj3CKDl.png",
		Name:             "Shio",
		Description:      "A rich, savory broth, delicately seasoned to bring out a robust and satisfying flavor.",
		Price:            13,
	}

	broth_2 := entities.Broth{
		ID:               "2",
		ImageInactiveURL: "https://i.imgur.com/3Htjgtf.png",
		ImageActiveURL:   "https://i.imgur.com/QP9mS58.png",
		Name:             "Shoyu",
		Description:      "A deep and flavorful soy-based broth, blending salty and slightly sweet notes for a classic Japanese taste.",
		Price:            18,
	}

	broth_3 := entities.Broth{
		ID:               "3",
		ImageInactiveURL: "https://i.imgur.com/Vleln0q.png",
		ImageActiveURL:   "https://i.imgur.com/lO1ss1J.png",
		Name:             "Spicy",
		Description:      "A bold, spicy broth, infused with vibrant chili flavors that add a thrilling heat to every.",
		Price:            15,
	}

	broth_4 := entities.Broth{
		ID:               "4",
		ImageInactiveURL: "https://i.imgur.com/GNEf9SS.png",
		ImageActiveURL:   "https://i.imgur.com/CCmprc8.png",
		Name:             "Charmiss",
		Description:      "A unique broth that masterfully balances a hint of bitterness with a subtle sweetness, creating a complex and delightful flavor profile.",
		Price:            23,
	}

	protein_1 := entities.Protein{
		ID:            "1",
		ImageInactive: "https://i.imgur.com/YGf3LhS.png",
		ImageActive:   "https://i.imgur.com/29LeLqj.png",
		Name:          "Chasu",
		Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
		Price:         10,
	}

	protein_2 := entities.Protein{
		ID:            "2",
		ImageInactive: "https://i.imgur.com/OdHnV52.png",
		ImageActive:   "https://i.imgur.com/aDLB7av.png",
		Name:          "Chicken",
		Description:   "A tender sliced chicken breast, marinated and infused with rich flavors.",
		Price:         13,
	}

	protein_3 := entities.Protein{
		ID:            "3",
		ImageInactive: "https://i.imgur.com/kRIZ1Fv.png",
		ImageActive:   "https://i.imgur.com/7Xh8lKQ.png",
		Name:          "Beef",
		Description:   "Juicy sliced beef, seasoned to perfection and complemented by a vibrant assortment of vegetables.",
		Price:         17,
	}
	protein_4 := entities.Protein{
		ID:            "4",
		ImageInactive: "https://i.imgur.com/VZG8ZLS.png",
		ImageActive:   "https://i.imgur.com/OBT1iZl.png",
		Name:          "Tofu",
		Description:   "Delicately sliced tofu, imbued with savory spices and paired with a colorful array of seasonal vegetables.",
		Price:         8,
	}
	protein_5 := entities.Protein{
		ID:            "5",
		ImageInactive: "https://i.imgur.com/RSZLzcb.png",
		ImageActive:   "https://i.imgur.com/WvGo64u.png",
		Name:          "Tako",
		Description:   "Succulent sliced tako, lightly seasoned and harmoniously matched with a selection of fresh spices.",
		Price:         24,
	}
	protein_6 := entities.Protein{
		ID:            "6",
		ImageInactive: "https://i.imgur.com/b8mIyQ6.png",
		ImageActive:   "https://i.imgur.com/I6ve8au.png",
		Name:          "Salmon",
		Description:   "Silky sliced salmon, seasoned with aromatic herbs.",
		Price:         23,
	}

	r.broths = append(r.broths, &broth_1, &broth_2, &broth_3, &broth_4)
	r.proteins = append(r.proteins, &protein_1, &protein_2, &protein_3, &protein_4, &protein_5, &protein_6)

	r.orderImage["Chasu"] = "https://i.imgur.com/LVTKR3F.png"
	r.orderImage["Chicken"] = "https://i.imgur.com/6oJMsJx.png"
	r.orderImage["Beef"] = "https://i.imgur.com/bWSVeTW.png"
	r.orderImage["Tofu"] = "https://i.imgur.com/sDxUnWv.png"
	r.orderImage["Tako"] = "https://i.imgur.com/dGze1gR.png"
	r.orderImage["Salmon"] = "https://i.imgur.com/G7H27xS.png"
}

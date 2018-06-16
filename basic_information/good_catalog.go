package basic_information

type GoodCatalog interface {
	Add(good Good)
	Remove(good Good)
	Get(id int64) Good
}

package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(maps map[string]interface{}, offset int, size int) (tags []Tag) {
	db.Where(maps).Offset(offset).Limit(size).Find(&tags)

	return
}

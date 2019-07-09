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

func GetTagsTotal(maps map[string]interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

func CreateTag(name string, createdBy string) bool {
	db.Create(&Tag{Name: name, CreatedBy: createdBy})

	return true
}

func UpdateTag(id int, maps map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(maps)

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(Tag{})

	return true
}

func CleanAllTag() bool {
	// unscoped 代表不使用默认的软删除
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})

	return true
}
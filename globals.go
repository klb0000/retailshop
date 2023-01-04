package retailshop

import (
	"os"
	"path"

	"gorm.io/gorm"
)

const MaxIdLen = 32
const MinProductNameLen = 3

var currPath, _ = os.Getwd()
var DefaultDB *gorm.DB = GetDB(path.Join(currPath, "data/data.db"))

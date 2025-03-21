package conf

import "restfulapi/driver"

func InitDbConnX() {
	if preforkFlag && childFlag {
		driver.InitConnX()
	} else if !preforkFlag {
		driver.InitConnX()
	}
}

package packages

type Packages struct {
	ID                  uint   `json:"id"'`
	PackageName         string `json:"package_name"`
	PackagePrice        int    `json:"package_price"`
	PackageDurationDays int    `json:"package_duration_days"`
	Status              bool   `json:"status"`
}

type AssignUserPackage struct {
	PackageID uint `json:"package_id"`
}

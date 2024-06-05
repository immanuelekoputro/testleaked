package users

type UserViewHistories struct {
	VisitorUserId uint   `json:"visitor_user_id"`
	HostUserId    uint   `json:"host_user_id"`
	VisitorAction string `json:"visitor_action"`
	IsSuperlike   bool   `json:"is_superlike"`
}

type UserSubscribes struct {
	UserId    uint   `json:"user_id"`
	PackageId uint   `json:"package_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type VwUserSubscribes struct {
	PackageID          uint   `json:"package_id"`
	PackageName        string `json:"package_name"`
	UserId             uint   `json:"user_id"`
	Name               string `json:"name"`
	SubscribeStatus    bool   `json:"subscribe_status"`
	SubscribeStartDate string `json:"subscribe_start_date"`
	SubscribeEndDate   string `json:"subscribe_end_date"`
	PackageDuration    int    `json:"package_duration_days"`
}

type SubmitActionUser struct {
	HostID        uint   `json:"host_id"`
	VisitorAction string `json:"visitor_action"`
}

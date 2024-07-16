package dto

type QueryManagerAccount struct {
	AvailableUponRegistration string   `form:"available_upon_registration"`
	Roles                     string   `form:"roles"`
	AdditionalRoles           []string `form:"additional_roles"`
	ProjectPermissions        []string `form:"project_permissions"`
}

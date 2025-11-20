package metservices

type MetClient struct {
	BaseURL string
}

// Returns an array of the different departments within the Met
type MetDepartments struct {
	Departments []MetSingleDepartment `json:"departments"`
}

// Container for a single department within the array of Met departments
type MetSingleDepartment struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

type DepartmentsResponse struct {
	Departments []MetSingleDepartment `json:"departments"`
}

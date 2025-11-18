package met

type MetDepartments struct {
	Departments []MetSingleDepartment `json:"departments"`
}

type MetSingleDepartment struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

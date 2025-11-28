package metmuseum

// Init new Client for Met API calls
type MetClient struct {
	BaseURL string
}

// Returns an array of the different departments within the Met
type MetDepartments struct {
	Departments []MetSingleDepartment `json:"departments"`
}

// Struct for returning list of departments at the Met
type DepartmentsResponse struct {
	Departments []MetSingleDepartment `json:"departments"`
}

// Container for a single department within the array of Met departments
type MetSingleDepartment struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

// Returns ID's of all highlights in the Met collection
type GetHighlightIDs struct {
	ObjectIds []int `json:"objectIds"`
}

// Struct for a single piece of art from the Met
type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	IsHighlight       bool   `json:"isHighlight"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
	Department        string `json:"department"`
	ObjectName        string `json:"objectName"`
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
}

export interface MetDepartment {
  departmentId: number;
  displayName: string;
}

export type MetDepartments = MetDepartment[];

export interface MetHighlight {
  objectID: number;
  primaryImage: string;
  primaryImageSmall: string;
  department: string;
  objectName: string;
}

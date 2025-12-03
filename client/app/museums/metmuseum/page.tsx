import { apiGet } from "@/lib/apiGet";
import { MetDepartments } from "@/types/met";

export default async function page() {
  const metDepartments = await apiGet<MetDepartments>("/metmuseum/departments");

  // Returns a list of the departments in the Met
  const departmentList = metDepartments.map((department) => {
    return (
      <div key={department.departmentId}>
        <p>{department.displayName}</p>
      </div>
    );
  });

  return (
    <div>
      <h1>Departments</h1>
      {departmentList}
    </div>
  );
}

import { apiGet } from "@/lib/apiGet";
import { MetDepartments } from "@/types/met";
import Link from "next/link";

export default async function page() {
  const metDepartments = await apiGet<MetDepartments>("/metmuseum/departments");

  // Returns a list of the departments in the Met
  const departmentList = metDepartments.map((department) => {
    return (
      <Link
        href={`/museums/metmuseum/${department.departmentId}`}
        key={department.departmentId}
      >
        <p>{department.displayName}</p>
      </Link>
    );
  });

  return (
    <div>
      <h1>Departments</h1>
      {departmentList}
    </div>
  );
}

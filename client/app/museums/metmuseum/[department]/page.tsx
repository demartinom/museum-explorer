import { apiGet } from "@/lib/apiGet";
import { MetDepartments } from "@/types/met";

interface Props {
  params: { department: string };
}

export default async function DepartmentPage({ params }: Props) {
  const { department } = await params;

  const departments: MetDepartments = await apiGet("/metmuseum/departments");

  const pageDepartment = departments.find(
    (d) => d.departmentId.toString() === department
  );

  if (!pageDepartment) return <p>Department not found</p>;

  return <div>{pageDepartment.displayName}</div>;
}

export async function generateStaticParams() {
  const departments: MetDepartments = await apiGet("/metmuseum/departments");

  return departments.map((d) => ({
    department: d.departmentId.toString(),
  }));
}

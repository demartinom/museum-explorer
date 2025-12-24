import { apiGet } from "@/lib/apiGet";
import { MetDepartments, MetHighlight } from "@/types/met";
import Image from "next/image";
import parse from "html-react-parser";

interface Props {
  params: { department: string };
}

export default async function DepartmentPage({ params }: Props) {
  const { department } = await params;

  const departments: MetDepartments = await apiGet("/metmuseum/departments");
  const pageDepartment = departments.find(
    (d) => d.departmentId.toString() === department
  );
  const highlights: MetHighlight[] = await apiGet(
    `/metmuseum/departments/${pageDepartment?.departmentId}`
  );
  console.log(highlights);

  const departmentHighlights = highlights.map((item, index) => (
    <div key={index}>
      <p>{parse(item.objectName)}</p>
      <Image
        src={item.primaryImageSmall}
        width={300}
        height={300}
        alt="highlight"
      ></Image>
    </div>
  ));
  if (!pageDepartment) return <p>Department not found</p>;

  return (
    <div>
      {pageDepartment.displayName}
      {departmentHighlights}
    </div>
  );
}

export async function generateStaticParams() {
  const departments: MetDepartments = await apiGet("/metmuseum/departments");

  return departments.map((d) => ({
    department: d.departmentId.toString(),
  }));
}

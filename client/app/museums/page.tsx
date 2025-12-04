import { apiGet } from "@/lib/apiGet";
import { Museumlist } from "@/types/museum-list";
import Link from "next/link";

export default async function Museums() {
  const museumArray = await apiGet<Museumlist>("/museums");

  // Returns a list of museums the site has access to
  const museumlist = museumArray.map((museum) => {
    return (
      <Link href={`/museums/${museum.slug}`} key={museum.slug}>
        <p>{museum.name}</p>
      </Link>
    );
  });

  return (
    <div>
      <h1>Explore Museum Collections</h1>
      {museumlist}
    </div>
  );
}

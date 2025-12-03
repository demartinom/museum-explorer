import { apiGet } from "@/lib/apiGet";
import { Museumlist } from "@/types/museum-list";

export default async function Museums() {
  const museumArray = await apiGet<Museumlist>("/museums");
  const museumlist = museumArray.map((museum) => {
    return (
      <a href={`/museums/${museum.slug}`} key={museum.slug}>
        <p>{museum.name}</p>
      </a>
    );
  });
  return (
    <div>
      <h1>Explore Museum Collections</h1>
      {museumlist}
    </div>
  );
}

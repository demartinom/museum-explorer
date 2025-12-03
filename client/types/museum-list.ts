// Interface for taking array of museums from backend and putting them in objects
interface Museum {
  name: string;
  slug: string;
}

export type Museumlist = Museum[];

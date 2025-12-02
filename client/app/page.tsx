import { apiGet } from "@/lib/apiGet";
import { DailyArtwork } from "@/types/daily-artwork";
import Image from "next/image";

export const metadata = {
  title: "Museum Explorer",
};
export default async function Home() {
  // Retrieve artwork of the day
  const daily = await apiGet<DailyArtwork>("/daily/dailyartwork");

  return (
    <div>
      <div>
        <p>{daily.Title}</p>
        <p>{daily.Artist == "" ? "Artist Unknown" : daily.Artist}</p>
        <Image
          src={daily.ImageURL}
          alt="Image of daily artwork"
          width={500}
          height={500}
          preload={true}
        />
        <p>From the {daily.Museum}</p>
      </div>
    </div>
  );
}

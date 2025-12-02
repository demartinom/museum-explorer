//Reusable function to make api calls to Go backend
// Takes in T interface and path to add to base url and returns promise of type T
export async function apiGet<T>(path: string): Promise<T> {
  //Takes base URL from env
  const base = process.env.API_URL;
  if (!base) throw new Error("API_URL not set");

  const res = await fetch(`${base}${path}`, {
    method: "GET",
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(`API GET error: ${res.status} - ${text}`);
  }
  // Returns response in form of provided type
  return res.json() as Promise<T>;
}

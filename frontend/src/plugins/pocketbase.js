import PocketBase from "pocketbase";

const pb = new PocketBase(import.meta.env.VITE_API_ADDRESS || "/");

export default pb;

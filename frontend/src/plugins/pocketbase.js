import PocketBase from "pocketbase";
import { ApiAddress } from "../config/api";

const pb = new PocketBase(ApiAddress);

export default pb;

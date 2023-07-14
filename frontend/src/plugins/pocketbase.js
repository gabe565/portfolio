import PocketBase from "pocketbase";
import { ApiPath } from "../config/api";

const pb = new PocketBase(ApiPath());

export default pb;

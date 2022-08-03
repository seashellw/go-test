import { $ } from "zx";
import { clean, cloneBuildAssets } from "./kit.mjs";

await clean();

await cloneBuildAssets();

await $`wails dev`;

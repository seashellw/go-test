import { $ } from "zx";
import { clean, cloneBuildAssets } from "./kit.mjs";

await clean();

await cloneBuildAssets();

setTimeout(() => {

}, 1000);

await $`wails dev`;

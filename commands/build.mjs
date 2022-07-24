import { rename } from "fs/promises";
import { $ } from "zx";
import { clean, cloneBuildAssets } from "./kit.mjs";

await clean();
await cloneBuildAssets();
await $`wails build`;

await rename("./build/bin/toolkit.exe", "./toolkit.exe");
await clean();
await rename("./toolkit.exe", "./build/toolkit.exe");

import { copyFile } from "fs/promises";
import { $ } from "zx";
import { clean, cloneBuildAssets } from "./kit.mjs";

await clean();
await cloneBuildAssets();
await $`wails build`;

await copyFile("./build/bin/toolkit.exe", "./toolkit.exe");

import { rename } from "fs/promises";
import { $ } from "zx";
import { clean, cloneBuildAssets } from "./kit.mjs";

await clean();
await cloneBuildAssets();
await $`cd ./frontend && pnpm build`;
await $`wails build -s -clean`;

await rename("./build/bin/go.exe", "./go.exe");
await clean();
await rename("./go.exe", "./build/go.exe");

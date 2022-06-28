import { $ } from "zx";
import { rm, rename, mkdir } from "fs/promises";

await $`pnpm vue-tsc --noEmit`;
await $`pnpm vite build`;
await $`go build`;

await rm("./dist", { recursive: true });
await mkdir("./dist");
await rename("./go-test.exe", "./dist/go.exe");

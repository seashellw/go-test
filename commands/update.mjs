import { $, cd } from "zx";

await $`wails update -pre`;
await $`go get -u`;
await $`go mod tidy`;
await $`pnpm up`;
cd("./frontend")
await $`pnpm up`;

import { existsSync } from "fs";
import { copyFile, mkdir, rm, writeFile } from "fs/promises";

export const clean = async () => {
  if (existsSync("./build")) {
    await rm("./build", { recursive: true });
  }
  if (!existsSync("./frontend/dist")) await mkdir("./frontend/dist");
  if (!existsSync("./frontend/dist/index.html"))
    await writeFile("./frontend/dist/index.html", "再次运行 pnpm build");
  await mkdir("./build");
};

export const cloneBuildAssets = async () => {
  if (!existsSync("./build")) await mkdir("./build");
  await copyFile(
    "./frontend/public/appicon.png",
    "./build/appicon.png"
  );
};

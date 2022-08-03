import { existsSync } from "fs";
import { copyFile, mkdir, rm } from "fs/promises";

export const clean = async () => {
  if (existsSync("./build")) await rm("./build", { recursive: true });
  await mkdir("./build");
  if (!existsSync("./frontend/dist")) await mkdir("./frontend/dist");
};

export const cloneBuildAssets = async () => {
  if (!existsSync("./build")) await mkdir("./build");
  await copyFile(
    "./frontend/public/appicon.png",
    "./build/appicon.png"
  );
};

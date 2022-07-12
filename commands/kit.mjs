import { copyFile, mkdir, rm } from "fs/promises";

export const clean = async () => {
  await rm("./build", { recursive: true });
  await mkdir("./build");
};

export const cloneBuildAssets = async () => {
  await copyFile(
    "./frontend/src/assets/appicon.png",
    "./build/appicon.png"
  );
};

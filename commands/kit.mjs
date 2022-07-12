import { mkdir, rename, rm } from "fs/promises";

export const clean = async () => {
  await rm("./build", { recursive: true });
  await mkdir("./build");
};

export const cloneBuildAssets = async () => {
  await rename(
    "./frontend/src/assets/appicon.png",
    "./build/appicon.png"
  );
};

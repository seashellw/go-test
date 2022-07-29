export const reloadApp = async () => {
  await fetch("http://localhost:9001/api/app/reload", {
    method: "POST"
  });
};

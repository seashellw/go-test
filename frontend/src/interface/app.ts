export const reloadApp = () => {
  fetch("http://localhost:9001/api/app/reload", {
    method: "POST",
  });
};

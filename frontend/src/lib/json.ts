export const isStrJson = (str: any) => {
  if (typeof str !== "string") return false;
  str = str.trim();
  if (!str.startsWith("{")) return false;
  try {
    const obj = JSON.parse(str);
    return typeof obj == "object" && obj;
  } catch {
    return false;
  }
};

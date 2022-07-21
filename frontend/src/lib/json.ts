export const isStrJson = (str: string) => {
  if (typeof str !== "string") return false;
  str = str.trim();
  if (!str.startsWith("{")) return false;
  try {
    var obj = JSON.parse(str);
    return typeof obj == "object" && obj;
  } catch {
    return false;
  }
};

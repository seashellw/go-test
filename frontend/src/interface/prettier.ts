declare const prettier: any;
declare const prettierPlugins: any;

export const format = (text: string, type: string) => {
  try {
    return prettier.format(text, {
      parser: type,
      plugins: prettierPlugins,
    }) as string;
  } catch (e) {
    return `${e}`;
  }
};

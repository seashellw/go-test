declare const prettier: any;
declare const prettierPlugins: any;
declare const Prism: any;

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

export const highlightAll = () => {
  requestIdleCallback(() => {
    Prism.highlightAll();
  });
};

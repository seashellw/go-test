import { Routes } from "@/router";
import { ref, watch } from "vue";
import { useRoute } from "vue-router";

export const useRouterTransition = () => {
  const route = useRoute();

  const routeMap = new Map<string, number>();

  for (const item of Routes.filter((item) => item.meta?.show).map(
    (item, index) => ({
      path: item.path,
      index: index + 1,
    })
  )) {
    routeMap.set(item.path, item.index);
  }

  const transitionName = ref<"slide-down" | "slide-up" | "fade">(
    "fade"
  );

  watch(
    () => route.matched[0]?.path,
    (to, from) => {
      let newIndex = routeMap.get(to);
      let oldIndex = routeMap.get(from);
      if (!newIndex || !oldIndex) {
        transitionName.value = "fade";
        return;
      }
      if (newIndex > oldIndex) {
        transitionName.value = "slide-up";
        return;
      }
      if (newIndex < oldIndex) {
        transitionName.value = "slide-down";
        return;
      }
      transitionName.value = "fade";
    }
  );

  return transitionName;
};

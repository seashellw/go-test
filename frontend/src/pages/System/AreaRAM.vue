<script setup lang="ts">
import { Mutable } from "@/lib/type";
import { Gauge, GaugeOptions } from "@antv/g2plot";
import { useIntervalFn } from "@vueuse/core";
import { onMounted, ref } from "vue";
import { SysGetMemPercent } from "wails/go/main/App";

const container = ref<HTMLElement | undefined>(undefined);

let area: Gauge | undefined = undefined;

const currentValue = ref(0);

const options = ref<Mutable<GaugeOptions>>({
  percent: 0,
  axis: {
    label: {
      formatter(v) {
        return Number(v) * 100;
      },
    },
    subTickLine: {
      count: 3,
    },
  },
  statistic: {
    content: {
      formatter: (datum) => {
        return `${datum?.percent * 100}%`;
      },
      style: {
        fontSize: "1rem",
        color: "#fff",
      },
    },
  },
});

useIntervalFn(async () => {
  let percent = await SysGetMemPercent();
  percent = Math.round(percent);
  currentValue.value = percent;
  area?.changeData(percent / 100);
}, 1500);

onMounted(() => {
  if (!container.value) return;
  area = new Gauge(container.value, options.value);
  area.render();
});
</script>

<template>
  <div class="box">
    <p class="text-center pt-2">内存使用率 {{ currentValue }}%</p>
    <div ref="container" class="container"></div>
  </div>
</template>

<style scoped>
.box {
  width: 15rem;
}
.container {
  width: 10rem;
  height: 10rem;
  margin: 0 auto;
}
</style>

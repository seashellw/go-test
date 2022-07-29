<script setup lang="ts">
import { Mutable } from "@/lib/type";
import { Line, LineOptions } from "@antv/g2plot";
import { useElementBounding, useIntervalFn } from "@vueuse/core";
import { computed, onMounted, ref } from "vue";
import {
  SysGetCpuPercent,
  SysGetMemPercent,
} from "wails/go/main/App";

const REFRESH_INTERVAL = 1000;

const container = ref<HTMLElement | undefined>(undefined);

const data = ref<
  {
    value: number;
    time: number;
    name: string;
  }[]
>([]);

const options = ref<Mutable<LineOptions>>({
  data: data.value,
  xField: "time",
  yField: "value",
  seriesField: "name",
  smooth: true,
  yAxis: {
    min: 0,
    max: 100,
  },
  xAxis: {
    label: null,
  },
  area: {
    style: {
      fillOpacity: 0.15,
    },
  },
  tooltip: {
    showTitle: false,
    formatter: (datum) => {
      return { name: datum.name, value: datum.value + "%" };
    },
  },
  legend: {
    position: "bottom",
    itemName: {
      style: {
        fill: "#fff",
      },
    },
  },
});

let area: Line | undefined = undefined;

const { width } = useElementBounding(container);
const xLength = computed(() => {
  return Math.round(width.value / 7);
});

const pushData = (
  percent: { cpu: number; mem: number },
  now?: number
) => {
  now = now || Date.now();
  data.value.push(
    {
      value: percent.cpu,
      time: now,
      name: "CPU",
    },
    {
      value: percent.mem,
      time: now,
      name: "内存",
    }
  );
};

useIntervalFn(async () => {
  pushData({
    cpu: Math.round(await SysGetCpuPercent()),
    mem: Math.round(await SysGetMemPercent()),
  });
  if (data.value.length > xLength.value) {
    data.value = data.value.slice(0 - xLength.value);
  }
  area?.changeData(data.value);
}, REFRESH_INTERVAL);

onMounted(async () => {
  if (!container.value) return;
  let now = Date.now();
  data.value = [];
  const percent = {
    cpu: Math.round(await SysGetCpuPercent()),
    mem: Math.round(await SysGetMemPercent()),
  };
  while (data.value.length < xLength.value) {
    now = now - REFRESH_INTERVAL;
    pushData(percent, now);
  }
  options.value.data = data.value;
  area = new Line(container.value, options.value);
  area.render();
});
</script>

<template>
  <div class="pl-2 pr-4">
    <p class="p-1 text-center text-lg">资源占用率</p>
    <div ref="container" class="area-container"></div>
  </div>
</template>

<style scoped>
.area-container {
  height: 15rem;
}
</style>

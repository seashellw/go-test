<script setup lang="ts">
import { Mutable } from "@/lib/type";
import { Line, LineOptions } from "@antv/g2plot";
import { useElementBounding, useIntervalFn } from "@vueuse/core";
import { onMounted, ref } from "vue";
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

useIntervalFn(async () => {
  let percent = await SysGetCpuPercent();
  percent = Math.round(percent * 10) / 10;
  let now = Date.now();
  data.value.push({
    value: percent,
    time: now,
    name: "CPU",
  });
  percent = await SysGetMemPercent();
  percent = Math.round(percent * 10) / 10;
  data.value.push({
    value: percent,
    time: now,
    name: "内存",
  });
  let length = Math.round(width.value / 2);
  if (data.value.length > length) {
    data.value = data.value.slice(0 - length);
  }
  options.value.data = data.value;
  area?.changeData(options.value.data);
}, REFRESH_INTERVAL);

onMounted(async () => {
  if (!container.value) return;
  let length = Math.round(width.value / 3);
  let now = Date.now();
  let cpu = await SysGetCpuPercent();
  cpu = Math.round(cpu * 10) / 10;
  let mem = await SysGetMemPercent();
  mem = Math.round(mem * 10) / 10;
  let li: typeof data.value = [];
  while (li.length < length) {
    li.push(
      {
        value: cpu,
        time: now,
        name: "CPU",
      },
      {
        value: mem,
        time: now,
        name: "内存",
      }
    );
    now = now - REFRESH_INTERVAL;
  }
  data.value = li;
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

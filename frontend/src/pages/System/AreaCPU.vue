<script setup lang="ts">
import { Mutable } from "@/lib/type";
import { Line, LineOptions } from "@antv/g2plot";
import { useIntervalFn } from "@vueuse/core";
import { onMounted, ref } from "vue";
import { SysGetCpuPercent } from "wails/go/main/App";

const container = ref<HTMLElement | undefined>(undefined);

const data = ref<
  {
    value: number;
    time: number;
  }[]
>([]);

const options = ref<Mutable<LineOptions>>({
  data: data.value,
  xField: "time",
  yField: "value",
  smooth: true,
  yAxis: {
    min: 0,
    max: 100,
  },
  xAxis: {
    label: null,
  },
});

let area: Line | undefined = undefined;

const currentValue = ref(0);

useIntervalFn(async () => {
  let percent = await SysGetCpuPercent();
  percent = Math.round(percent * 10) / 10;
  currentValue.value = percent;
  if (data.value.length > 20) {
    data.value.shift();
  }
  data.value.push({
    value: percent,
    time: Date.now(),
  });
  area?.changeData(options.value.data);
}, 500);

onMounted(() => {
  if (!container.value) return;
  area = new Line(container.value, options.value);
  area.render();
});
</script>

<template>
  <div class="box">
    <p class="text-center pt-2">CPU使用率 {{ currentValue }}%</p>
    <div ref="container" class="container"></div>
  </div>
</template>

<style scoped>
.box,
.container {
  width: 15rem;
}
.container {
  height: 8rem;
}
</style>

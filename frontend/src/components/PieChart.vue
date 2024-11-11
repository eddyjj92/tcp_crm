<template>
  <canvas ref="chartCanvas"></canvas>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { Chart } from 'chart.js/auto';

const props = defineProps({
  chartData: Object,
});

const chartCanvas = ref(null);
let chartInstance;

const createChart = () => {
  if (chartInstance) chartInstance.destroy();
  chartInstance = new Chart(chartCanvas.value, {
    type: 'pie',
    data: props.chartData,
    options: {
      responsive: true,
    },
  });
};

watch(props.chartData, createChart, { immediate: true });
onMounted(createChart);
</script>

<style scoped>
canvas {
  max-width: 300px;
  max-height: 300px;
}
</style>

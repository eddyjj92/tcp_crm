<template>
  <q-page padding class="full-page">
    <div class="q-pa-md full-height">
      <span class="text-h6 q-mb-md">Dashboard</span>
      <div class="row full-height">
        <div class="col-12 full-width">
          <canvas ref="barChart" class="full-height"></canvas>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { Chart } from 'chart.js/auto';

const barChart = ref(null);

let barChartInstance;

// Datos de ejemplo, reemplázalos con datos reales de tu API
const fetchData = () => {
  const purchaseSalesData = {
    labels: ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio', 'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'],
    datasets: [
      {
        label: 'Compras',
        backgroundColor: 'rgba(54, 162, 235, 0.5)',
        borderColor: 'rgba(54, 162, 235, 1)',
        data: [50, 40, 70, 55, 65, 75, 100, 50, 80, 19, 56, 78],
      },
      {
        label: 'Ventas',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
        borderColor: 'rgba(255, 99, 132, 1)',
        data: [60, 45, 80, 60, 70, 85,50, 40, 70, 55, 65, 75],
      },
    ],
  };

  // Inicializar gráficos
  createBarChart(purchaseSalesData);
};

// Funciones para crear los gráficos
const createBarChart = (data) => {
  if (barChartInstance) barChartInstance.destroy();
  barChartInstance = new Chart(barChart.value, {
    type: 'bar',
    data,
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'top',
        },
        title: {
          display: true,
          text: 'Compras vs Ventas',
        },
      },
    },
  });
};

onMounted(fetchData);
</script>

<style scoped>

</style>

<template>
  <div
    style="height: 200px; width: 280px; display: flex; flex-direction: column"
  >
    <vue3-chart-js
      :id="doughnutChart.id"
      :type="doughnutChart.type"
      :data="doughnutChart.data"
      :options="doughnutChart.options"
      @before-render="beforeRenderLogic"
    ></vue3-chart-js>
  </div>
</template>
<script>
import Vue3ChartJs from "@j-t-mcc/vue3-chartjs";

export default {
  name: "App",
  components: {
    Vue3ChartJs,
  },
  props: {
    countEmail: Number,
  },
  setup(props) {
    console.log(props.countEmail);
    const doughnutChart = {
      id: "doughnut",
      type: "doughnut",
      data: {
        labels: ["E-Mails", "SMS"],
        datasets: [
          {
            backgroundColor: ["#9DCBE5", "#578098"],
            data: [props.countEmail, 2],
          },
        ],
      },
      options: {
        cutout: 67,
        plugins: {
          legend: {
            position: "right",
          },
        },
      },
    };

    return {
      doughnutChart,
    };
  },
};
</script>

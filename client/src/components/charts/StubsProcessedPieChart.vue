<template>
    <canvas ref="pieChart"></canvas>
</template>

<script>
	import { Chart, registerables } from 'chart.js';

    export default {
    	name: "StubsProcessedPieChart",
		props: {
			dataset: Array
		},
        mounted() {
        	const ctx = this.$refs.pieChart;
        	Chart.register(...registerables);

			let data = {
				labels:[],
				datasets:[{
					label: "apps by status",
					data: [],
					hoverOffset: 4
				}]
			};
			for (const elem of this.dataset) {
				data.labels.push(elem.Status);
				data.datasets[0].data.push(elem.Count);
			}
			const config = {
  				type: 'doughnut',
  				data: data
			};

            new Chart(ctx, config);
        }
    }
</script>
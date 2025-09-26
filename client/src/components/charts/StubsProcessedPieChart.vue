<template>
    <canvas ref="pieChart"></canvas>
</template>

<script>
	import { Chart, registerables } from 'chart.js';

    export default {
    	name: "StubsProcessedPieChart",
		props: {
			unprocessed: Number,
            total: Number
		},
        mounted() {
        	const ctx = this.$refs.pieChart;
        	Chart.register(...registerables);

			let data = {
				labels:["unprocessed", "processed"],
				datasets:[{
					label: "apps by status",
					data: [this.unprocessed, this.total - this.unprocessed],
					hoverOffset: 4
				}]
			};
			const config = {
  				type: 'doughnut',
  				data: data
			};

            new Chart(ctx, config);
        }
    }
</script>
<template>
    <canvas ref="pieChart"></canvas>
</template>

<script>
	import { Chart, registerables } from 'chart.js';

    export default {
    	name: "GamesByReleaseYearBarChart",
		props: {
			unreleased: Number,
            datasetByYear: Array
		},
        mounted() {
        	const ctx = this.$refs.pieChart;
        	Chart.register(...registerables);

            this.datasetByYear.sort((a, b) => {
                if (a.ReleaseYear > b.ReleaseYear) return 1;
                if (a.ReleaseYear < b.ReleaseYear) return -1;
                return 0;
            });

			let data = {
				labels:["unreleased (yet)"],
				datasets:[{
					label: "games",
					data: [this.unreleased],
					borderWidth: 1
				}]
			};

            for (const elem of this.datasetByYear) {
				data.labels.push(elem.ReleaseYear);
				data.datasets[0].data.push(elem.Count);
			}

			const config = {
  				type: 'bar',
  				data: data,
                options: {
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    }
                }
			};

            new Chart(ctx, config);
        }
    }
</script>
<template>
    <canvas ref="pieChart"></canvas>
</template>

<script>
	import { Chart, registerables } from 'chart.js';

    export default {
    	name: "GamesByReleaseYearBarChart",
		props: {
            dataset: Array
		},
        mounted() {
        	const ctx = this.$refs.pieChart;
        	Chart.register(...registerables);

            this.dataset.TotalGamesReleasedByYears.sort((a, b) => {
                if (a.ReleaseYear > b.ReleaseYear) return 1;
                if (a.ReleaseYear < b.ReleaseYear) return -1;
                return 0;
            });

			let data = {
				labels:["soon"],
				datasets:[{
					label: "games",
					data: [this.dataset.TotalUnreleasedYetGames],
					borderWidth: 1
				}]
			};

            for (const elem of this.dataset.TotalGamesReleasedByYears) {
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
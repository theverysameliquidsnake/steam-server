<template>
    <a class="nav-link" href="#" @click="showCharts">
        <i class="bi bi-bar-chart"></i>
        Charts
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "ChartsNavLink",
        methods: {
            async showCharts() {
                try {
                    this.$parent.switchPage();
                    this.$parent.displaySpinner();
                    const response = await axios.get("http://localhost:8080/chart/dataset");
                    this.$parent.chartsData = response.data.data;
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                } catch(error) {
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    }
</script>
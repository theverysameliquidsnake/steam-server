<template>
    <a class="nav-link" href="#" @click="pullStubs">
        <i class="bi bi-pen"></i>
        All Apps
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "EditExceptionsNavLink",
        methods: {
            async pullStubs() {
                try {
                    this.$parent.switchPage();
                    this.$parent.displaySpinner();
                    const response = await axios.get(`http://localhost:8080/stub/all/${this.$parent.pulledStubs.length}`);
                    this.$parent.pulledStubs = response.data.data;
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                } catch(error) {
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    };
</script>
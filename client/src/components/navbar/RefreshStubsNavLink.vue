<template>
    <a class="dropdown-item" href="#"  @click="refreshStubs">
        <i class="bi bi-arrow-clockwise"></i>
        Refresh Apps
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "RefreshStubsNavLink",
        methods: {
            async refreshStubs() {
                try {
                    this.$parent.displaySpinner();
                    const response = await axios.put("http://localhost:8080/stub/refresh");
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
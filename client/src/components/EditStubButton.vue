<template>
    <button type="button" class="btn btn-link" @click="pullStubs">
        <i class="bi bi-pen"></i>
        <span class="ps-2">Edit "Stubs"</span>
    </button>
</template>

<script>
    import axios from "axios"
    import utils from "../utils/utils"

    export default {
        name: "EditStubsButton",
        methods: {
            async pullStubs() {
                try {
                    this.$parent.displaySpinner();
                    this.$parent.pulledStubs = [];
                    const response = await axios.get("http://localhost:8080/stub/all");
                    this.$parent.pulledStubs = response.data.data;
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                } catch(error) {
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    };
</script>
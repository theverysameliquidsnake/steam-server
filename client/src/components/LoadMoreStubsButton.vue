<template>
    <button type="button" class="btn btn-primary mb-3" @click="pullStubs">Load More</button>
</template>

<script>
    import axios from "axios"
    import utils from "../utils/utils"

    export default {
        name: "LoadMoreStubsButton",
        methods: {
            async pullStubs() {
                try {
                    this.$parent.displaySpinner();
                    const response = await axios.get(`http://localhost:8080/stub/all/${this.$parent.pulledStubs.length}`);
                    if (response.data.data.length) {
                        this.$parent.pulledStubs = this.$parent.pulledStubs.concat(response.data.data);
                    }
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
<template>
    <button type="button" class="btn btn-link" @click="refreshStubs">
        <i class="bi bi-arrow-clockwise"></i>
        <span class="ps-2">Refresh "Stubs"</span>
    </button>
</template>

<script>
    import axios from "axios"
    import utils from "../utils/utils"

    export default {
        name: "RefreshStubsButton",
        methods: {
            async refreshStubs() {
                try {
                    const response = await axios.put("http://localhost:8080/stub/refresh");
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                } catch(error) {
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    };
</script>
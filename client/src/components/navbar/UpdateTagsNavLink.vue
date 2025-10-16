<template>
    <a class="dropdown-item" href="#" @click="updateTags()">
        <i class="bi bi-tag"></i>
        Refresh Tags
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "UpdateTagsNavLink",
        methods: {
            async updateTags() {
                try {
                    this.$parent.displaySpinner();
                    let response = await axios.put("http://localhost:8080/tag/refresh");
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
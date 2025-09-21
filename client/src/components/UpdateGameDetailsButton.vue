<template>
    <button type="button" class="btn btn-link" @click="updateGameDetails">
        <i class="bi bi-controller"></i>
        <span class="ps-2">Update Game Details</span>
    </button>
</template>

<script>
    import axios from "axios"
    import utils from "../utils/utils"

    export default {
        name: "UpdateGameDetailsButton",
        methods: {
            async updateGameDetails() {
                try {
                    this.$parent.displaySpinner();
                    let response = await axios.get("http://localhost:8080/stub/request");
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                    response = await axios.put(`http://localhost:8080/game/insert/${response.data.data.appid}`);
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                    this.$parent.hideSpinner();
                } catch(error) {
                    this.$parent.hideSpinner();
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    };
</script>
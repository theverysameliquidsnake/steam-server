<template>
    <a class="dropdown-item" href="#" @click="updateGameDetails">
        <i class="bi bi-controller"></i>
        Update Game Details (Debug)
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "UpdateGameDetailsNavLink",
        methods: {
            async updateGameDetails() {
                try {
                    this.$parent.displaySpinner();
                    let response = await axios.get("http://localhost:8080/stub/request");
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                    response = await axios.put(`http://localhost:8080/game/insert/${response.data.data.appid}`);
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
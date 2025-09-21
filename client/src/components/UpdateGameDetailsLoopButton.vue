<template>
    <button v-if="updateIsRunning" type="button" class="btn btn-link" @click="stopUpdating">
        <i class="bi bi-clock-history"></i>
        <span class="ps-2">Stop Update Game Details (Loop)</span>
    </button>
    <button v-else type="button" class="btn btn-link" @click="updateGameDetails">
        <i class="bi bi-clock-history"></i>
        <span class="ps-2">Update Game Details (Loop)</span>
    </button>
</template>

<script>
    import axios from "axios"
    import utils from "../utils/utils"

    export default {
        name: "UpdateGameDetailsButton",
        data() {
            return {
                updateIsRunning: false,
                intervalRef: null
            }
        },
        methods: {
            updateGameDetails() {
                this.intervalRef = setInterval(async () => {
                    try {
                        this.$parent.displaySpinner();
                        this.updateIsRunning = true;
                        let response = await axios.get("http://localhost:8080/stub/request");
                        this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                        response = await axios.put(`http://localhost:8080/game/insert/${response.data.data.appid}`);
                        this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                        this.$parent.hideSpinner();
                    } catch(error) {
                        this.$parent.hideSpinner();
                        this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                    }
                }, 2000);
            },

            stopUpdating() {
                this.updateIsRunning = false;
                clearInterval(this.intervalRef);
            }
        }
    };
</script>
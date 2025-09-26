<template>
    <a v-if="updateIsRunning" class="dropdown-item" href="#" @click="stopUpdating">
        <i class="bi bi-clock-history"></i>
        Stop Update Game Details (Loop)
    </a>
    <a v-else class="dropdown-item" href="#" @click="updateGameDetails">
        <i class="bi bi-clock-history"></i>
        Update Game Details (Loop)
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "UpdateGameDetailsLoopNavLink",
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
                        this.$parent.hideSpinner();
                        this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                    } catch(error) {
                        this.$parent.hideSpinner();
                        this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                    }
                }, 2500);
            },

            stopUpdating() {
                this.updateIsRunning = false;
                clearInterval(this.intervalRef);
            }
        }
    };
</script>
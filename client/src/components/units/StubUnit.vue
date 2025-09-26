<template>
    <tr>
        <td><a :href="steamAppLink">{{ appId }}</a></td>
        <td>{{ appName }}</td>
        <td>{{ appType }}</td>
        <td>
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="needsUpdate" :checked="needsUpdate" disabled>
                <label class="form-check-label" for="needsUpdate">Needs Update</label>
            </div>
        </td>
        <td>
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="skip" :checked="skip" disabled>
                <label class="form-check-label" for="skip">Skip</label>
            </div>
        </td>
        <td>
            <div class="form-check">
                <input class="form-check-input" type="checkbox" value="" id="skip" :checked="ignore" @change="handleCheck($event)">
                <label class="form-check-label" for="skip">Ignore</label>
            </div>
        </td>
    </tr>
</template>

<script>
    import { computed } from 'vue';
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "StubUnit",
        props: {
            appName: String,
            appId: Number,
            appType: String,
            needsUpdate: Boolean,
            skip: Boolean,
            ignore: Boolean
        },
        computed: {
            steamAppLink() {
                return `https://store.steampowered.com/app/${this.appId}`;
            }
        },
        methods: {
            async handleCheck(event) {
                try {
                    this.$parent.displaySpinner();
                    let formData = new FormData();
                    formData.append("appid", this.appId);
                    formData.append("ignore", event.target.checked);
                    let response = await axios.patch("http://localhost:8080/stub/ignore", formData);
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
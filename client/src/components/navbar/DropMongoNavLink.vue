<template>
    <a class="dropdown-item" href="#" @click="dropDatabase">
        <i class="bi bi-folder-minus"></i>
        Drop MongoDB
    </a>
</template>

<script>
    import axios from "axios";
    import utils from "../../utils/utils";

    export default {
        name: "DropMongoNavLink",
        methods: {
            async dropDatabase() {
                try {
                    const response = await axios.delete("http://localhost:8080/mongo/drop");
                    this.$parent.addToast("Success", utils.getCurrentTime(), response.data.message);
                } catch(error) {
                    this.$parent.addToast("Error", utils.getCurrentTime(), error.response.data.error);
                }
            }
        }
    }
</script>
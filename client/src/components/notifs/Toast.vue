<template>
    <div ref="liveToast" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
        <div class="toast-header">
            <strong class="me-auto">{{ headerBig }}</strong>
            <small>{{ headerSmall }}</small>
            <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
        <div class="toast-body">
            {{ description }}
        </div>
    </div>
</template>

<script>
    import { Toast } from 'bootstrap';

    export default {
        name: "Toast",
        props: {
            headerBig: String,
            headerSmall: String,
            description: String
        },
        mounted() {
            if (this.$refs.liveToast) {
                const toastLiveExample = this.$refs.liveToast
                const toastBootstrap = Toast.getOrCreateInstance(toastLiveExample);
                toastBootstrap.show();
            }
        }
    }
</script>
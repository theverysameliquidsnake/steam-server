<template>
    <div class="card">
        <h5 class="card-header">{{ appId }}</h5>
        <div class="card-body">
            <h5 class="card-title">{{ appName }}</h5>
        </div>
    </div>
</template>

<script>
    export default {
        name: "StubUnit",
        props: {
            appName: String,
            appId: String,
            needsUpdate: Boolean,
            Skip: Boolean,
            ManualSkip: Boolean
        }
    }
</script>
<template>
    <!-- Top Navbar -->
    <nav class="px-1 navbar navbar-dark bg-dark">
        <div class="container-fluid justify-content-start">
            <i class="bi bi-bar-chart-steps" style="color: white;"></i>
            <a class="navbar-brand ps-2" href="#">Steam Charts</a>
        </div>
    </nav>
    <!-- Side Toolbar -->
    <div class="row">
        <div class="col-3" >
            <div class="d-flex">
                <edit-stub-button></edit-stub-button>
            </div>
            <div class="d-flex">
                <button type="button" class="btn btn-link" data-bs-toggle="collapse" data-bs-target="#collapseExample" aria-expanded="false" aria-controls="collapseExample">
                    <i class="bi bi-tools"></i>
                    <span class="ps-2">Tools</span>
                </button>
            </div>
            <div class="collapse ps-1" id="collapseExample">
                <ul>
                    <li>
                        <refresh-stubs-button></refresh-stubs-button>
                    </li>
                    <li>
                        <update-game-details-button></update-game-details-button>
                    </li>
                    <li>
                        <update-game-details-loop-button></update-game-details-loop-button>
                    </li>
                    <li>
                        <drop-mongo-button></drop-mongo-button>
                    </li>
                </ul>
            </div>
        </div>
        <div class="col-9">
            <div class="row">
                <div v-for="pulledStub in pulledStubs" class="col-4">
                    <stub-unit :app-id="pulledStub.AppId" :app-name="pulledStub.Name"></stub-unit>
                </div>
            </div>
        </div>
    </div>
    <!-- Toast & Spinner -->
    <div class="toast-container position-fixed bottom-0 end-0 p-3">
        <template v-for="toastData in toastsData">
            <toast :header-big="toastData.headerBig" :header-small="toastData.headerSmall" :description="toastData.description"></toast>
        </template>
    </div>
    <div v-if="showSpinner" class="toast-container position-fixed bottom-0 start-0 p-3">
        <div class="spinner-border" role="status">
            <span class="visually-hidden">Loading...</span>
        </div>
    </div>
</template>

<script>
    import StubUnit from './StubUnit.vue';
    import EditStubButton from './EditStubButton.vue';
    import RefreshStubsButton from './RefreshStubsButton.vue';
    import DropMongoButton from './DropMongoButton.vue';
    import UpdateGameDetailsButton from './UpdateGameDetailsButton.vue';
    import UpdateGameDetailsLoopButton from './UpdateGameDetailsLoopButton.vue';
    import Toast from './notifs/Toast.vue';

    export default {
        name: "App",
        data() {
            return {
                showSpinner: false,
                pulledStubs: [],
                toastsData: []
            }
        },
        components: {
            StubUnit,
            EditStubButton,
            RefreshStubsButton,
            DropMongoButton,
            UpdateGameDetailsButton,
            UpdateGameDetailsLoopButton,
            Toast
        },
        methods: {
            addToast(headerBig, headerSmall, description) {
                this.toastsData.push({headerBig: headerBig, headerSmall: headerSmall, description: description})
            },

            displaySpinner() {
                this.showSpinner = true;
            },

            hideSpinner() {
                this.showSpinner = false;
            }
        }
    }
</script>
<style>
    html, body {
        height: 100%;
        max-width: 100%;
        overflow-x: hidden;
    }
</style>

<template>
    <!-- Top Navbar -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
            <brand></brand>
            <button class="navbar-toggler" 
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarNavDropdown"
                aria-controls="navbarNavDropdown"
                aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNavDropdown">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <edit-exceptions-nav-link></edit-exceptions-nav-link>
                    </li>
                    <li class="nav-item">
                        <charts-nav-link></charts-nav-link>
                    </li>
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle"
                            href="#"
                            id="navbarDropdownMenuLink"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false">
                            <i class="bi bi-tools"></i>
                            Tools
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
                            <li>
                                <refresh-stubs-nav-link></refresh-stubs-nav-link>
                            </li>
                            <li>
                                <update-tags-nav-link></update-tags-nav-link>
                            </li>
                            <li>
                                <update-game-details-nav-link></update-game-details-nav-link>
                            </li>
                            <li>
                                <update-game-details-loop-nav-link></update-game-details-loop-nav-link>
                            </li>
                            <li>
                                <drop-mongo-nav-link></drop-mongo-nav-link>
                            </li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </nav> 
    <!-- Stubs Page -->
    <div v-if="isInvalidStubsRendered">
        <table class="table table-hover">
            <tbody>
                <template v-for="pulledStub in pulledStubs">
                    <stub-unit :app-id="pulledStub.AppId"
                        :app-name="pulledStub.Name"
                        :app-type="pulledStub.Type"
                        :new-stub="pulledStub.New"
                        :first-update="pulledStub.FirstUpdate"
                        :second-update="pulledStub.SecondUpdate"
                        :error="pulledStub.Error"
                        :ignore="pulledStub.Ignore">
                    </stub-unit>
                </template>
            </tbody>
        </table>
        <div class="d-flex justify-content-center mt-2">
            <stubs-pagination-button></stubs-pagination-button>
        </div>
    </div>
    <!-- Charts Page -->
    <div v-if="isChartsRendered">
        <div class="row">
            <div class="col-3">
                <stubs-processed-pie-chart :dataset="chartsData.TotalStubsByStatus"></stubs-processed-pie-chart>
            </div>
            <div class="col-3">
                <stubs-type-pie-chart :dataset="chartsData.TotalStubsByType"></stubs-type-pie-chart>
            </div>
            <div class="col-6">
                <games-by-release-year-bar-chart :dataset="chartsData.GamesByYearDataset"></games-by-release-year-bar-chart>
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
    import Brand from './navbar/Brand.vue';
    import EditExceptionsNavLink from './navbar/EditExceptionsNavLink.vue';
    import RefreshStubsNavLink from './navbar/RefreshStubsNavLink.vue';
    import UpdateTagsNavLink from './navbar/UpdateTagsNavLink.vue';
    import UpdateGameDetailsNavLink from './navbar/UpdateGameDetailsNavLink.vue';
    import UpdateGameDetailsLoopNavLink from './navbar/UpdateGameDetailsLoopNavLink.vue';
    import DropMongoNavLink from './navbar/DropMongoNavLink.vue';
    import ChartsNavLink from './navbar/ChartsNavLink.vue';

    import StubUnit from './units/StubUnit.vue';

    import StubsProcessedPieChart from './charts/StubsProcessedPieChart.vue';
    import StubsTypePieChart from './charts/StubsTypePieChart.vue';
    import GamesByReleaseYearBarChart from './charts/GamesByReleaseYearBarChart.vue';

    import StubsPaginationButton from './buttons/StubsPaginationButton.vue';

    import Toast from './notifs/Toast.vue';

    export default {
        name: "App",
        data() {
            return {
                showSpinner: false,
                pulledStubs: [],
                chartsData: {},
                toastsData: []
            }
        },
        components: {
            Brand,
            EditExceptionsNavLink,
            RefreshStubsNavLink,
            UpdateTagsNavLink,
            UpdateGameDetailsNavLink,
            UpdateGameDetailsLoopNavLink,
            DropMongoNavLink,
            ChartsNavLink,

            StubUnit,

            StubsProcessedPieChart,
            StubsTypePieChart,
            GamesByReleaseYearBarChart,
            
            StubsPaginationButton,

            Toast
        },
        computed: {
            isInvalidStubsRendered() {
                return this.pulledStubs.length;
            },

            isChartsRendered() {
                return Object.keys(this.chartsData).length;
            }
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
            },

            switchPage() {
                this.pulledStubs = [];
                this.chartsData = {};
            }
        }
    }
</script>
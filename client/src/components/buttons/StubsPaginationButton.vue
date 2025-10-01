<template>
    <nav aria-label="Page navigation example">
        <ul class="pagination">
            <li class="page-item">
                <a class="page-link" href="#" aria-label="Previous" @click="prevPage">
                    <span aria-hidden="true">&laquo;</span>
                </a>
            </li>
            <template v-for="page in pages">
                <li class="page-item">
                    <a :class="{'page-link': true, 'active': (page === currentPage)}" href="#" @click="toPage(page)">{{ page }}</a>
                </li>
            </template>
            <li class="page-item">
                <a class="page-link" href="#" aria-label="Next" @click="nextPage">
                    <span aria-hidden="true">&raquo;</span>
                </a>
            </li>
        </ul>
    </nav>
</template>

<script>
    import axios from "axios"
    import utils from "../../utils/utils"

    export default {
        name: "StubsPaginationButton",
        data() {
            return {
                currentPage: 1,
                pagesCount: 0,
                pages: []
            }
        },
        async mounted() {
                try {
                    const response = await axios.get("http://localhost:8080/stub/count");
                    this.pagesCount = Math.ceil(response.data.data.count / 50);
                    this.pages = utils.calculatePagination(this.currentPage, this.pagesCount);
                } catch(error) {
                    console.error(error.response.data.error);
                }
        },
        methods: {
            async prevPage() {
                this.currentPage = this.currentPage - 1;
                this.pages = utils.calculatePagination(this.currentPage, this.pagesCount);
                await this.pullStubs(this.currentPage - 1);
            },

            async nextPage() {
                this.currentPage = this.currentPage + 1;
                this.pages = utils.calculatePagination(this.currentPage, this.pagesCount);
                await this.pullStubs(this.currentPage - 1);
            },

            async toPage(page) {
                this.currentPage = page;
                this.pages = utils.calculatePagination(this.currentPage, this.pagesCount);
                await this.pullStubs(this.currentPage - 1);
            },

            async pullStubs(pageOffcet) {
                try {
                    this.$parent.displaySpinner();
                    const response = await axios.get(`http://localhost:8080/stub/all/${pageOffcet * 50}`);
                    if (response.data.data.length) {
                        this.$parent.pulledStubs = response.data.data;
                    }
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
<template>
    <h1> Your Wine List</h1>
    <body>
        <div class="wine_table">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Wine Type</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(wine, index) in wines" :key="index">
                        <td v-text="wine.id"></td>
                        <td v-text="wine.name"></td>
                        <td v-text="wine.wine_type"></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </body>
</template> 

<script>

import axios from 'axios';
export default {
    name: "WineList",
    data() {
        return {            
            wines:[]
        }
    },
    methods: {
        getWines() {
            axios.get("http://localhost:8081/wine")
            .then(res => {
                this.wines =res.data;
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        }
  },
    mounted() {
        this.getWines();
    }
}
</script>


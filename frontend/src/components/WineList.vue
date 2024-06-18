<template>
    <h1> Your Wine List</h1>
    <DataTable :value="wines" dataKey="id" tableStyle="min-width: 60rem">
        <Column field="name" header="Name"></Column>
        <Column field="wine_type" header="Type"></Column>
    </DataTable>
</template>

<script>
import axios from 'axios';
export default {
    name: "WineList",
    data() {
        return {            
            wines:[],
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


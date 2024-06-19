<template>
    <h1> Your Wine List</h1>
    <div class="card flex justify-center">
        <Button label="+ Add Wine" @click="visible = true" />
        <Dialog v-model:visible="visible" modal header="What wine would you like to add?" :style="{ width: '25rem' }">
            <AddWine></AddWine>
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                <Button type="button" label="Save" @click="visible = false"></Button>
            </div>
        </Dialog>
    </div> 
    <DataTable :value="wines" dataKey="id" tableStyle="min-width: 60rem">
        <Column field="name" header="Name"></Column>
        <Column field="wine_type" header="Type"></Column>
    </DataTable>
</template>

<script>
import axios from 'axios';
import AddWine from './AddWine.vue'
export default {
    name: "WineList",
    components: {
        AddWine
    },
    data() {
        return {            
            wines:[],
            visible: false
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
        },
    },
    mounted() {
        this.getWines();
    }
}
</script>


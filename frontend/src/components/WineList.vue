
<script>
import axios from 'axios'
import AddWine from './AddWine.vue'
import DeleteWine from './DeleteWine.vue'
import UpdateWine from './UpdateWine.vue'

export default {
    name: "WineList",
    components: {
        AddWine,
        DeleteWine,
        UpdateWine
    },
    data() {
        return {
            wines: [],
            selected: {},
            wine_types: [],
            countries: [],
            delete_dialog: false,
            wine_dialog: false,
        }
    },
    computed: {
        wineTypeMap: function() {
            return this.wine_types.reduce((acc,cur)=>{acc[cur.id]=cur; return acc},{})
        },
        countryMap: function() {
            return this.countries.reduce((acc,cur)=>{acc[cur.id]=cur; return acc},{})
        },
    },
    methods: {
        getWines() {
            axios.get("http://localhost:8081/wine")
            .then(res => {
                this.wines =res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
        getWineTypes() {
            axios.get("http://localhost:8081/wine_type")
            .then(res => {
                this.wine_types = res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        getCountries() {
            axios.get("http://localhost:8081/country")
            .then(res => {
                this.countries = res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        selectWine(wn) {
            this.selected = wn
            this.wine_dialog = true
        },
        getWineTypeName(wine){
            return this.wineTypeMap[wine.data.wine_type].name
        },
        getCountryName(wine){
            return this.countryMap[wine.data.country].name
        }
    },
    mounted() {
        this.getCountries()
        this.getWineTypes()
        this.getWines()
    }
}
</script>

<template>
    <h1> Your Wine List</h1>
    <AddWine :wineTypes="wine_types" :countries="countries"/>
    <DataTable v-model:selection="selected" :value="wines" dataKey="id" tableStyle="min-width: 60rem">
        <Column field="name" header="Name"></Column>
        <Column field="wine_type" header="Type">
            <template #body="wine">
                {{ getWineTypeName(wine)}}
            </template>
        </Column>
        <Column field="country" header="Country">
            <template #body="wine">
                {{ getCountryName(wine)}}
            </template>
        </Column>
        <Column headerStyle="width:4rem">
            <template #body="item">
                <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Filter" @click="selectWine(item.data)" />
                <Button icon="pi pi-pencil" severity="secondary" rounded text aria-label="Filter" @click="selectWine(item.data)" />
            </template>
        </Column>
    </DataTable>
    <DeleteWine v-model:visible="delete_dialog" :selectedWine="selected" />
    <UpdateWine v-model:visible="wine_dialog" :selectedWine="selected" :wineTypes="wine_types" :countries="countries"/>
</template>
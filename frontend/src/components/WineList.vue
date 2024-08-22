<script>
import axios from 'axios'
import AddWine from './AddWine.vue'
import DeleteWine from './DeleteWine.vue'
import UpdateWine from './UpdateWine.vue'
import { store } from './store'

export default {
    name: "WineList",
    components: {
        AddWine,
        DeleteWine,
        UpdateWine
    },
    data() {
        return {
            expanded_row: {},
            store,
            wines: [],
            selected: {},
            delete_dialog: false,
            wine_dialog: false,
        }
    },
    computed: {
        wineTypeMap: function() {
            return this.store.wine_types.reduce((acc,cur)=>{acc[cur.id]=cur; return acc},{})
        },
        countryMap: function() {
            return this.store.countries.reduce((acc,cur)=>{acc[cur.id]=cur; return acc},{})
        },
    },
    methods: {
        ExpandRow(data) {
            if (this.expanded_row[data.id]) {
                delete this.expanded_row[data.id];
            } else {
                this.expanded_row[data.id] = true;
            }
        },
        getWines() {
            axios.get("/wine")
            .then(res => {
                this.wines =res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
        getWineTypes() {
            axios.get("/wine_type")
            .then(res => {
                this.store.wine_types = res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        getCountries() {
            axios.get("/country")
            .then(res => {
                this.store.countries = res.data
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        selectWine(wn) {
            this.selected = wn
            this.wine_dialog = true
        },
        deleteWine(wn) {
            this.selected = wn
            this.delete_dialog = true
        },
        getWineTypeName(wine){
            return this.wineTypeMap[wine.data.wine_type]?.name || 'Unknown'
        },
        getCountryName(wine){
            return this.countryMap[wine.data.country]?.name || 'Unknown'
        },
        removeWine(wn) {
            this.wines = this.wines.filter(wine => wine.id !== wn)
        },
        addNewWine(wn) {
            this.wines.push(wn)
        },
    },
    async mounted() {
        this.getCountries()
        this.getWineTypes()
        this.getWines()
    }
}
</script>

<template>
    <h1> Your Wine List</h1>
    <AddWine :wines="wines" @wine-added="addNewWine"/> 
    <DataTable v-model:expandedRows="expanded_row" :value="wines" dataKey="id"   tableStyle="min-width: 60rem" @row-click="ExpandRow($event.data)">
        <Column headerStyle="width:4rem">
            <template #body="wine">
                <i class="pi pi-chevron-right" style="color: #708090" v-if="!expanded_row[wine.data.id]" />
                <i class="pi pi-chevron-down" style="color: #708090" v-else />
            </template>
        </Column>
        <Column field="name" header="Name" />
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
        <Column field="score" header="Score">
            <template #body="wine">
                <Rating v-model="wine.data.score" readonly/>
            </template>
        </Column>
        <Column headerStyle="width:4rem">
            <template #body="item">
                <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Filter" @click="deleteWine(item.data)" />
                <Button icon="pi pi-pencil" severity="secondary" rounded text aria-label="Filter" @click="selectWine(item.data)" />
            </template>
        </Column>
        <template #expansion="slotProps">
            <div class="p-4">
                <p>Producer: {{ slotProps.data.producer }}</p>
                <p>Year: {{ slotProps.data.year }}</p>
            </div>
        </template>
    </DataTable>
    <DeleteWine v-model:visible="delete_dialog" :selected="selected" @wine_deleted="removeWine"/>
    <UpdateWine v-model:visible="wine_dialog" :selected="selected" />
</template> 
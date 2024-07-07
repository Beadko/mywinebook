<template>
    <h1> Your Wine List</h1>
    <AddWine></AddWine>
    <DataTable v-model:selection="selected_wine" :value="wines" dataKey="id" tableStyle="min-width: 60rem">
        <Column field="name" header="Name"></Column>
        <Column field="wine_type" header="Type"></Column>
        <Column headerStyle="width:4rem">
            <template #body="item">
                <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Filter" @click="confirmDelete(item.data)"/>
                <Button icon="pi pi-pencil" severity="secondary" rounded text aria-label="Filter" @click="selectWine(item.data)"/>
            </template>
        </Column>
    </DataTable>
    <Dialog v-model:visible="deleteDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="selected_wine">
                    Are you sure you want to delete <b>{{ selected_wine.name }}</b>?
                </span>
            </div>
            <template #footer>
                <Button label="No" icon="pi pi-times" text @click="deleteDialog = false" />
                <Button label="Yes" icon="pi pi-check" @click="deleteWine" />
            </template>
        </Dialog>
        <Dialog v-model:visible="wineDialog" :style="{ width: '450px' }" header="Wine Details" :modal="true">
            <div class="flex flex-col gap-6 p-fluid">
                <div>
                    <label for="name" class="font-semibold w-24">Name</label>
                    <InputText id="name" v-model="selected_wine.name" class="flex-auto" autocomplete="off"/>
                </div>
                <div>
                    <label for="wine_type" class="block font-bold mb-3">Type</label>
                    <Select v-model="selected_wine.wine_type" :options="wine_types" optionLabel="name" optionValue="id" />
                </div>
            </div>
            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="wineDialog = false" />
                <Button label="Save" icon="pi pi-check" @click="updateWine" />
            </template>
        </Dialog>
</template>

<script>
import axios from 'axios'
import AddWine from './AddWine.vue'

export default {
    name: "WineList",
    components: {
        AddWine
    },
    data() {
        return {
            wines: [],
            selected_wine: {},
            wine_types:[],
            deleteDialog: false,
            wineDialog: false
        }
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
        confirmDelete(wn) {
            this.selected_wine = wn
            this.deleteDialog = true
        },
        deleteWine() {
            axios.delete("http://localhost:8081/wine/"+ this.selected_wine.id)
            .then(res => {
                this.deleteDialog = false;
                location.reload()
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
        selectWine(wn) {
            this.selected_wine = wn
            this.wineDialog = true
            axios.get("http://localhost:8081/wine_type")
            .then(res => {
                this.wine_types = res.data;
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        updateWine() {
            axios.put("http://localhost:8081/wine/"+ this.selected_wine.id)
            .then(res => {
                this.wineDialog = false;
                location.reload()
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        }
    },
    mounted() {
        this.getWines()
    }
}
</script>


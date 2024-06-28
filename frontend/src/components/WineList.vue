<template>
    <h1> Your Wine List</h1>
    <AddWine></AddWine>
    <DataTable v-model:selection="selected_wine" :value="wines" dataKey="id" tableStyle="min-width: 60rem">
        <Column field="name" header="Name"></Column>
        <Column field="wine_type" header="Type"></Column>
        <Column headerStyle="width:4rem">
            <template #body="item">
                <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Filter" @click="confirmDelete(item.data)"/>
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
</template>

<script>
import axios from 'axios'
import AddWine from './AddWine.vue'
import { ref } from "vue"

export default {
    name: "WineList",
    components: {
        AddWine
    },
    data() {
        return {
            wines: ref([]),
            selected_wine: ref({}),
            deleteDialog: ref(false)
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
        confirmDelete(wn) {
            this.selected_wine = wn;
            this.deleteDialog = true;
        },
        deleteWine() {
            axios.delete("http://localhost:8081/wine/"+ this.selected_wine.id)
            .then(res => {
                this.deleteDialog = false;
                location.reload()
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


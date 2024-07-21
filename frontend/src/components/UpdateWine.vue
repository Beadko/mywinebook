<script>
    import axios from 'axios'

export default {
    name: "UpdateWine",
    props: {
        selectedWine: Object,
        wineTypes: Array
    },
    methods: {
    updateWine() {
            axios.put("http://localhost:8081/wine/"+ this.selectedWine.id, this.selectedWine)
            .then(res => {
                this.update_dialog = false;
                location.reload()
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
    }
}
</script>

<template>
    <Dialog :style="{ width: '450px' }" header="Wine Details" :modal="true">
        <div class="flex flex-col gap-6 p-fluid">
            <div>
                <label for="name" class="font-semibold w-24">Name</label>
                <InputText id="name" v-model="selectedWine.name" class="flex-auto" autocomplete="off"/>
            </div>
            <div>
                <label for="wineType" class="block font-bold mb-3">Type</label>
                <Select v-model="selectedWine.wine_type" :options="wineTypes" optionLabel="name" optionValue="id" />
            </div>
        </div>
        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="this.$parent.wine_dialog = false" />
            <Button label="Save" icon="pi pi-check" @click="updateWine" />
        </template>
    </Dialog>
</template>
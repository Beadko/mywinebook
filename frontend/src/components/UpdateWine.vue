<script>
    import axios from 'axios'

export default {
    name: "UpdateWine",
    props: {
        selectedWine: Object,
        wineTypes: Array,
        countries: Array
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
    <Dialog :style="{ width: '25rem' }" header="Wine Details" modal>
        <div class="flex items-center gap-4 mb-4">
            <label for="name" class="font-semibold w-24">Name</label>
            <InputText id="name" v-model="selectedWine.name" class="w-full md:w-[14rem]" autocomplete="off"/>
        </div>
        <div class="flex items-center gap-4 mb-4">
            <label for="wineType" class="font-semibold w-24">Type</label>
            <Select v-model="selectedWine.wine_type" :options="wineTypes" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
        </div>
        <div class="flex items-center gap-4 mb-8">
            <label for="country" class="font-semibold w-24">Country</label>
            <Select v-model="selectedWine.country" :options="countries" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
        </div>
        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="this.$parent.wine_dialog = false" />
            <Button label="Save" icon="pi pi-check" @click="updateWine" />
        </template>
    </Dialog>
</template>
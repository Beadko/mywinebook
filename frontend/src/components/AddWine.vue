<template>
    <div class="card flex justify-center">
        <Button label="+ Add Wine" @click="visible = true" />
        <Dialog v-model:visible="visible" modal header="What wine would you like to add?" :style="{ width: '25rem' }">
            <div class="flex items-center gap-4 mb-4">
            <label for="name" class="font-semibold w-24">Name</label>
            <InputText id="name" class="flex-auto" autocomplete="off" v-model="wines" />
        </div>
        <div class="flex items-center gap-4 mb-8">
            <label for="wine_types" class="font-semibold w-24">Type</label>
            <Select v-model="selected" :options="wine_types" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
        </div>
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                <Button type="button" label="Save" @click="visible = false; addWine"></Button>
            </div>
        </Dialog>
    </div> 
</template>

<script>
import axios from 'axios'
import Select from 'primevue/select'
import { ref } from "vue"

export default {
    name: "AddWine",
    data() {
        return {            
            wines:[],
            visible: false,
            selected: ref(),
            wine_types: ref([])
        }
    },
    methods: {
        addWine() {
            const formData = new FormData()
            formData.append("name", "wine_type")
            const { data }= axios.post("http://localhost:8081/wine", formData)
            .then(res => {
                this.wines =res.data;
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        getWineTypes() {
            axios.get("http://localhost:8081/wine_type")
            .then(res => {
                this.wine_types = res.data;
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
    },
    beforeMount() {
        this.addWine()
    },
    mounted() {
        this.getWineTypes();
    }
}
</script>

<template>
    <div class="card flex justify-center">
        <Button label="+ Add Wine" @click="visible = true" />
        <Dialog v-model:visible="visible" modal header="What wine would you like to add?" :style="{ width: '25rem' }">
            <div class="flex items-center gap-4 mb-4">
            <label for="name" class="font-semibold w-24">Name</label>
            <InputText id="name" class="flex-auto" autocomplete="off" v-model="selected.name" />
        </div>
        <div class="flex items-center gap-4 mb-8">
            <label for="wine_types" class="font-semibold w-24">Type</label>
            <Select v-model="selected.wine_type" :options="wine_types" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
        </div>
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false" ></Button>
                <Button type="button" label="Save" @click="addWine"></Button>
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
            visible: false,
            selected: ref({
            name:'',
            wine_type: ''
            }),
            wine_types: ref([])
        }
    },
    methods: {
        addWine() {
            this.visible = false
            axios.post("http://localhost:8081/wine", this.selected)
            .then(res => {
                console.log(res.data)  
                window.alert(`Yay! Another wine added to your library`);          
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
        }
    },
    mounted() {
        this.getWineTypes();
    }
}
</script>

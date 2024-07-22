<script>
import axios from 'axios'

export default {
    name: "AddWine",
    props: {
        wineTypes: Array,
        countries: Array
    },
    data() {
        return {            
            visible: false,
            selected: {
            name:'',
            wine_type: '',
            country: ''
            },
        }
    },
    methods: {
        addWine() {
            this.visible = false
            axios.post("http://localhost:8081/wine", this.selected)
            .then(res => {
                console.log(res.data)  
             })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
    }
}
</script>

<template>
    <div class="card flex justify-center">
        <Button label="+ Add Wine" @click="visible = true" />
        <Dialog v-model:visible="visible" modal header="What wine would you like to add?" :style="{ width: '25rem' }">
            <div class="flex items-center gap-4 mb-4">
                <label for="name" class="font-semibold w-24">Name</label>
                <InputText v-model="selected.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
            </div>
            <div class="flex items-center gap-4 mb-4">
                <label for="wineTypes" class="font-semibold w-24">Type</label>
                <Select v-model="selected.wine_type" :options="wineTypes" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
            </div>
            <div class="flex items-center gap-4 mb-8">
                <label for="countries" class="font-semibold w-24">Country</label>
                <Select v-model="selected.country" :options="countries" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" />
            </div>
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"/>
                <Button type="button" label="Save" @click="addWine" />
            </div>
        </Dialog>
    </div> 
</template>
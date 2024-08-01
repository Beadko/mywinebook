<script>
import axios from 'axios'
import { store } from './store';

export default {
    name: "Countries",
    inheritAttrs:false,
    props: {
        selectedCountry: Object,
    },
    data() {
        return {
            visible: false,
            countries: [],
            new_country: {
                name:''
            },
            store
        }    
    },
    methods: {
        addCountry() {
            axios.post("http://localhost:8081/country", this.new_country)
            .then(res => {
                console.log(res.data)  
                this.visible = false
                this.$emit('country-added');
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        handleCountryChange(value) {
            if (value === 'add_country') {
                this.visible = true;
                this.selected = '';
            }
        },
    },
}
</script>

<template>
    <div class="flex items-center gap-4 mb-8">
        <label for="country" class="font-semibold w-24">Country</label>
        <Select v-model="selectedCountry.country" :options="[...store.countries, { id: 'add_country', name: '+ Add New' }]" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" @change="handleCountryChange($event.value)" />
    </div>    
    <Dialog v-model:visible="visible" modal header="Add a new country" :style="{ width: '25rem' }">
        <div class="flex items-center gap-4 mb-4">
            <label for="name" class="font-semibold w-24">Country Name</label>
            <InputText v-model="new_country.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
        </div>
        <div class="flex justify-end gap-2">
            <Button type="button" label="Cancel" severity="secondary" text @click="visible = false" />
            <Button type="button" label="Save" @click="addCountry" />
        </div>
    </Dialog>
</template>
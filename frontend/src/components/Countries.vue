<script>
import axios from 'axios'
import { store } from './store';

export default {
    name: "Countries",
    inheritAttrs: false,
    props: {
        selected: Object,
    },
    data() {
        return {
            dialog_visible: false,
            countries: [],
            new_country: {
                name:''
            },
            store
        }    
    },
    methods: {
        addCountry() {
            axios.post("/country", this.new_country)
            .then( res => {
                console.log(res),
                this.store.countries.push(res.data),
                this.selected.country = res.data.id,
                this.dialog_visible = false
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        handleCountryChange(value) {
            if (value === 'add_country') {
                this.dialog_visible = true,
                this.selected.country = ''
            }
        },
        clearSelection() {
            this.selected.country = '',
            this.dialog_visible = false
        }
    },
}
</script>

<template>
    <div class="flex items-center gap-2 mb-3">
        <label for="country" class="font-semibold w-20">Country</label>
        <Select v-model="selected.country" :options="[...store.countries, { id: 'add_country', name: '+ Add New' }]" optionLabel="name" optionValue="id" class="w-full md:w-[14rem]" @change="handleCountryChange($event.value)" />
    </div>    
    <Dialog v-model:visible="dialog_visible" modal header="Add a new country" :style="{ width: '25rem' }">
        <div class="flex items-center gap-2 mb-4">
            <label for="name" class="font-semibold w-20">Country Name</label>
            <InputText v-model="new_country.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
        </div>
        <div class="flex justify-end gap-2">
            <Button type="button" label="Cancel" severity="secondary" text @click="clearSelection" />
            <Button type="button" label="Save" @click="addCountry" />
        </div>
    </Dialog>
</template>
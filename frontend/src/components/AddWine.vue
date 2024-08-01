<script>
import axios from 'axios'
import Countries from './Countries.vue';
import WineTypes from './WineTypes.vue';

export default {
    name: "AddWine",
    inheritAttrs:false,
    components: {
        Countries,
        WineTypes
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
        onCountryAdded() {
            this.$emit('country-added');
        },
        onTypeAdded() {
            this.$emit('type-added');
        }
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
            <WineTypes :selected="selected" @type-added="onTypeAdded" />
            <Countries :selected="selected" @country-added="onCountryAdded"/>
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"/>
                <Button type="button" label="Save" @click="addWine" />
            </div>
        </Dialog>
    </div> 
</template>
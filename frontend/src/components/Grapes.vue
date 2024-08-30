
<script>
import axios from 'axios'
import { store } from './store'

export default {
    name: "Grapes",
    inheritAttrs: false,
    props: {
        selected: Object,
    },
    data() {
        return {
            dialog_visible: false,
            grapes: [],
            new_grapes: {
                name:''
            },
            store
        }    
    },
    methods: {
        addGrapes() {
            axios.post("/grapes", this.new_grapes)
            .then( res => {
                this.store.grapes.push(res.data),
                this.selected.grapes = res.data.id,
                this.dialog_visible = false
                this.new_grapes.name = ''
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        handleGrapesChange(value) {
            if (value === 'add_grapes') {
                this.dialog_visible = true,
                this.selected.grapes = []
            }
        },
        clearSelection() {
            this.selected.grapes = [],
            this.dialog_visible = false
        }
    },
}
</script>

<template>
    <div class="flex items-center gap-2 mb-4">
        <label for="grapes" class="font-semibold w-20">Grapes</label>
        <MultiSelect 
            v-model="selected.grapes"
            display="chip"
            :options="store.grapes" 
            optionLabel="name" 
            optionValue="id" 
            filter 
            class="w-full md:w-[14rem]" 
        />
        <Button type="button" icon="pi pi-plus" size="small" @click="dialog_visible = true" />
    </div>    
    <Dialog 
        v-model:visible="dialog_visible" 
        modal 
        header="Add a new grape variety" 
        :style="{ width: '25rem' }"
    >
        <div class="flex items-center gap-2 mb-4">
            <label for="name" class="font-semibold w-20">Variety Name</label>
            <InputText 
                v-model="new_grapes.name" 
                id="name" 
                class="w-full md:w-[14rem]" 
                autocomplete="off" 
            />
        </div>
        <div class="flex justify-end gap-2">
            <Button type="button" label="Cancel" severity="secondary" text @click="clearSelection" />
            <Button type="button" label="Save" @click="addGrapes" />
        </div>
    </Dialog>
</template>
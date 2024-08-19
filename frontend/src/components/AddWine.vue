<script>
import axios from 'axios'
import Countries from './Countries.vue';
import WineTypes from './WineTypes.vue';

export default {
    name: "AddWine",
    inheritAttrs: false,
    components: {
        Countries,
        WineTypes,
    },
    props: {
        wines: Array
    },
    data() {
        return {
            visible: false,
            selected: {
                name: '',
                wine_type: '',
                country: '',
                score: '',
                producer: '',
                year: '',
                alcohol: '',
            },
        }
    },
    methods: {
        addWine() {
            this.visible = false
            axios.post("/wine", this.selected)
            .then(
                this.wines.push(this.selected)
            )
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
            })
        },
        onCountryAdded() {
            this.$emit('country-added')
        },
        onTypeAdded() {
            this.$emit('type-added')
        },
        cancelAdd() {
            this.selected = {},
            this.visible = false
        }
    },
}
</script>

<template>
    <div class="card flex justify-center">
        <Button label="+ Add Wine" @click="visible = true" />
        <Dialog v-model:visible="visible" modal header="What wine would you like to add?" :style="{ width: '25rem' }">
            <div class="flex items-center gap-2 mb-3">
                <label for="name" class="font-semibold w-24">Name</label>
                <InputText v-model="selected.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
            </div>
            <WineTypes :selected="selected" @type-added="onTypeAdded" />
            <Countries :selected="selected" @country-added="onCountryAdded"/>
            <div class="flex items-center gap-2 mb-3">
                <label for="score" class="font-semibold w-24">Score</label>
                <Rating v-model="selected.score" id="score" />
            </div>
            <div class="items-center gap-2 mb-3">
                <Panel header="Add more details" toggleable :collapsed="true">
                    <div class="flex items-center gap-4 mb-4">
                        <label for="producer" class="font-semibold w-24">Producer</label>
                        <InputText v-model="selected.producer" id="producer" class="w-full md:w-[14rem]" autocomplete="off" />
                    </div>
                    <div class="flex items-center gap-2 mb-3">
                        <label for="year" class="font-semibold w-24">Year</label>
                        <InputNumber v-model="selected.year" id="year" inputId="withoutgrouping" :useGrouping="false" fluid class="w-full md:w-[8rem]" />
                    </div>
                    <div class="flex items-center gap-2 mb-3">
                        <label for="alcohol" class="font-semibold w-24">Alcohol</label>
                        <InputNumber v-model="selected.alcohol" id="alcohol" inputId="decimal" :minFractionDigits="1" suffix="%" class="w-full md:w-[8rem]" />
                    </div>
                </Panel>
            </div>
            <div class="flex justify-end gap-4">
                <Button type="button" label="Cancel" severity="secondary" @click="cancelAdd"/>
                <Button type="button" label="Save" @click="addWine" />
            </div>
        </Dialog>
    </div> 
</template>
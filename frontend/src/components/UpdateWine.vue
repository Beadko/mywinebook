<script>
import axios from 'axios'
import WineTypes from './WineTypes.vue';
import Countries from './Countries.vue';

export default {
    name: "UpdateWine",
    components: {
        WineTypes,
        Countries
    },
    props: {
        selected: Object,
    },
    data() {
        return {
            temp_selected: {}
        }
    },
    watch: {
        selected: {
            immediate: true,
            handler(v) {
                if (v) {
                    this.temp_selected = { ...v };
                }
            }
        }
    },
    methods: {
        updateWine() {
            Object.assign(this.selected, this.temp_selected);
            axios.put("/wine/"+ this.selected.id, this.selected)
            .then(
                this.$parent.wine_dialog = false,
            )
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
        onCountryAdded() {
            this.$emit('country-added')
        },
        onTypeAdded() {
            this.$emit('type-added')
        },
    },
}
</script>

<template>
    <Dialog :style="{ width: '25rem' }" header="Wine Details" modal>
        <div class="flex items-center gap-2 mb-4">
            <label for="name" class="font-semibold w-24">Name</label>
            <InputText id="name" v-model="temp_selected.name" class="w-full md:w-[14rem]" autocomplete="off"/>
        </div>
        <WineTypes :selected="temp_selected" @type-added="onTypeAdded" />
        <Countries :selected="temp_selected" @country-added="onCountryAdded"/>
        <div class="flex items-center gap-2 mb-4">
            <label for="score" class="font-semibold w-24">Score</label>
            <Rating v-model="temp_selected.score" />
        </div>
        <div class="flex items-center gap-2 mb-3">
            <label for="year" class="font-semibold w-24">Year</label>
                <InputNumber v-model="temp_selected.year" inputId="withoutgrouping" :useGrouping="false" fluid class="w-full md:w-[8rem]" />
        </div>
        <div class="flex items-center gap-2 mb-3">
            <label for="alcohol" class="font-semibold w-24">Alcohol</label>
            <InputNumber v-model="temp_selected.alcohol" inputId="decimal" :minFractionDigits="1" suffix="%" class="w-full md:w-[8rem]" />
        </div>        
        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="this.$parent.wine_dialog = false" />
            <Button label="Save" icon="pi pi-check" @click="updateWine" />
        </template>
    </Dialog>
</template>
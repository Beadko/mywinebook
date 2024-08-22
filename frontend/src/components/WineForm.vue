<script>
import axios from 'axios';
import WineTypes from './WineTypes.vue'
import Countries from './Countries.vue'

export default {
    name: "WineForm",
    components: {
        WineTypes,
        Countries
    },
    props: {
        mode: {
            type: String,
            required: true,
            validator: value => ['add', 'update'].includes(value)
        },
        selected: {
            type: Object,
            default: () => ({
                name: '',
                wine_type: '',
                country: '',
                score: '',
                producer: '',
                year: null,
                alcohol: null,
            })
        }
    },
    computed : {
        displayedAlcohol: {
            get() {
                return this.formData.alcohol === 0 ? null : this.formData.alcohol
            },
            set(value) {
                this.formData.alcohol = value
            }
        },
        displayedYear: {
            get() {
                return this.formData.year === 0 ? null : this.formData.year
            },
            set(value) {
                this.formData.year = value
            }
        }
    },
    data() {
        return {
            formData: { ...this.selected }
        }
    },
    watch: {
        selected: {
            immediate: true,
            handler(v) {
                this.formData = { ...v }
            }
        }
    },
    methods: {
        saveWine() {
            if (this.mode === 'add') {
                axios.post("/wine", this.formData)
                    .then(res => {
                        this.$emit('wine_added', res.data)
                    })
                    .catch(error => {
                        window.alert(`The API returned an error: ${error}`)
                    });
            } else if (this.mode === 'update') {
                Object.assign(this.selected, this.formData)
                axios.put("/wine/" + this.selected.id, this.selected)
                    .then(() => {
                        this.$emit('wine_updated', this.selected)
                    })
                    .catch(error => {
                        window.alert(`The API returned an error: ${error}`)
                    });
            }
        },
        onCountryAdded() {
            this.$emit('country-added')
        },
        onTypeAdded() {
            this.$emit('type-added')
        },
        cancelForm() {
            this.formData = { ...this.selected }
            this.$emit('cancel')
        }
    }
}
</script>

<template>
    <Dialog v-model:visible="dialog_visible" modal :header="mode === 'add' ? 'Add Wine' : 'Update Wine'" :style="{ width: '25rem' }">
        <div class="flex items-center gap-2 mb-3">
            <label for="name" class="font-semibold w-24">Name</label>
            <InputText v-model="formData.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
        </div>
        <WineTypes :selected="formData" @type-added="onTypeAdded" />
        <Countries :selected="formData" @country-added="onCountryAdded"/>
        <div class="flex items-center gap-2 mb-3">
            <label for="score" class="font-semibold w-24">Score</label>
            <Rating v-model="formData.score" id="score" />
        </div>
        <div class="items-center gap-2 mb-3">
            <Panel header="More details" toggleable :collapsed="true">
                <div class="flex items-center gap-4 mb-4">
                    <label for="producer" class="font-semibold w-24">Producer</label>
                    <InputText v-model="formData.producer" id="producer" class="w-full md:w-[14rem]" autocomplete="off" />
                </div>
                <div class="flex items-center gap-2 mb-3">
                    <label for="year" class="font-semibold w-24">Year</label>
                    <InputNumber v-model="displayedYear" id="year" :useGrouping="false" class="w-full md:w-[8rem]" />
                </div>
                <div class="flex items-center gap-2 mb-3">
                    <label for="alcohol" class="font-semibold w-24">Alcohol %</label>
                    <InputNumber v-model="displayedAlcohol" 
                    id="alcohol" inputId="decimal" :minFractionDigits="1" class="w-full md:w-[8rem]" />
                </div>
            </Panel>
        </div>
        <div class="flex justify-end gap-4">
            <Button type="button" label="Cancel" severity="secondary" @click="cancelForm"/>
            <Button type="button" label="Save" @click="saveWine" />
        </div>
    </Dialog>
</template>

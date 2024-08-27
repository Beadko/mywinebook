<script>
import axios from 'axios';
import WineTypes from './WineTypes.vue'
import Countries from './Countries.vue'
import { store } from './store'

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
                flavour:'',
                sweetness: '',
                acidity: '',
                tannin: '',
                body: '',
                finish: '',
                balance: ''
            })
        },
        severityMap: {
            type: Object,
            required: true
        },
        colourMap: {
            type: Object,
            required: true
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
            formData: { ...this.selected },
            flavours: [],
            sweetnesses: [],
            acidities: [],
            tannins: [],
            bodies: [],
            finishes: [],
            balances: [],
            store
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
        },
        getSeverityClass(selectedValue, itemId) {
            const severity = this.severityMap[itemId]
            return {
                'selected-success': selectedValue === itemId && severity === 'success',
                'selected-warn': selectedValue === itemId && severity === 'warn',
                'selected-danger': selectedValue === itemId && severity === 'danger'
            }
        }
    }
}
</script>

<template>
    <Dialog v-model:visible="dialog_visible" modal :header="mode === 'add' ? 'Add Wine' : 'Update Wine'" :style="{ width: '25rem' }">
        <div class="flex items-center gap-2 mb-3">
            <label for="name" class="font-semibold w-20">Name</label>
            <InputText v-model="formData.name" id="name" class="w-full md:w-[14rem]" autocomplete="off" />
        </div>
        <WineTypes :selected="formData" @type-added="onTypeAdded" />
        <Countries :selected="formData" @country-added="onCountryAdded"/>
        <div class="flex items-center gap-2 mb-3">
            <label for="score" class="font-semibold w-20">Score</label>
            <Rating v-model="formData.score" id="score" />
        </div>
        <div class="items-center gap-2 mb-3">
            <Panel header="More details" toggleable :collapsed="true">
                <div class="flex items-center gap-4 mb-4">
                    <label for="producer" class="font-semibold w-18">Producer</label>
                    <InputText v-model="formData.producer" id="producer" class="w-full md:w-[13.5rem]" autocomplete="off" />
                </div>
                <div class="flex items-center gap-2 mb-3">
                    <label for="year" class="font-semibold w-20">Year</label>
                    <InputNumber v-model="displayedYear" id="year" :useGrouping="false" class="w-full md:w-[8rem]" />
                </div>
                <div class="flex items-center gap-2 mb-5">
                    <label for="alcohol" class="font-semibold w-20">Alcohol %</label>
                    <InputNumber v-model="displayedAlcohol" 
                    id="alcohol" inputId="decimal" :minFractionDigits="1" class="w-full md:w-[8rem]" />
                </div>
                <div class="flex flex-wrap items-center gap-2 mb-3">
                    <label for="flavour" class="font-semibold w-20">Flavour</label>
                    <div v-for="flavour in store.flavours" :key="flavour.id" class="flex items-center gap-2">
                        <Button :label="flavour.name"
                                rounded
                                outlined
                                size="small"
                                @click="formData.flavour = flavour.id" 
                                :class="[colourMap[flavour.id], { 'active': formData.flavour === flavour.id }]" />
                    </div>
                </div>
                <div class="flex flex-wrap items-center gap-4 mb-5">
                    <label for="sweetness" class="font-semibold w-20">Sweetness</label>
                    <div v-for="sweetness in store.sweetnesses" :key="sweetness.id" class="flex items-center gap-2">
                        <RadioButton v-model="formData.sweetness" :value="sweetness.id" :inputId="'sweetness-' + sweetness.id" />
                        <span class="text-sm">{{ sweetness.name }}</span>
                    </div>
                </div>
                <div class="flex flex-wrap items-center gap-4 mb-5">
                    <label for="acidity" class="font-semibold w-20">Acidity</label>
                    <div v-for="acidity in store.acidities" :key="acidity.id" class="flex items-center gap-2">
                        <RadioButton v-model="formData.acidity" :value="acidity.id" :inputId="'tannin-' + acidity.id" />
                        <span class="text-sm">{{ acidity.name }}</span>
                    </div>
                </div>
                <div class="flex flex-wrap items-center gap-4 mb-5">
                    <label for="tannin" class="font-semibold w-20">Tannin</label>
                    <div v-for="tannin in store.tannins" :key="tannin.id" class="flex items-center gap-2">
                        <RadioButton v-model="formData.tannin" :value="tannin.id" :inputId="'tannin-' + tannin.id" />
                        <span class="text-sm">{{ tannin.name }}</span>
                    </div>
                </div>
                <div class="flex flex-wrap items-center gap-4 mb-5">
                    <label for="body" class="font-semibold w-20">Body</label>
                    <div v-for="body in store.bodies" :key="body.id" class="flex items-center gap-2">
                        <RadioButton v-model="formData.body" :value="body.id" :inputId="'body-' + body.id" />
                        <span class="text-sm">{{ body.name }}</span>
                    </div>
                </div>
                <div class="flex items-center gap-2 mb-3">
                    <label for="finish" class="font-semibold w-20">Finish</label>
                    <div v-for="finish in store.finishes" :key="finish.id" class="flex items-center gap-2">
                        <Button :label="finish.name"
                                :severity="severityMap[finish.id]"
                                outlined size="small"
                                @click="formData.finish = finish.id" 
                                :class="getSeverityClass(formData.finish, finish.id)"/>
                    </div>
                </div>
                <div class="flex items-center gap-2 mb-3">
                    <label for="balance" class="font-semibold w-20">Balance</label>
                    <div v-for="balance in store.balances" :key="balance.id" class="flex items-center gap-2">
                        <Button :label="balance.name"
                                :severity="severityMap[balance.id]"
                                outlined 
                                size="small"
                                @click="formData.balance = balance.id" 
                                :class="getSeverityClass(formData.balance, balance.id)" />
                    </div>
                </div>
            </Panel>
        </div>
        <div class="flex justify-end gap-4">
            <Button type="button" label="Cancel" severity="secondary" @click="cancelForm"/>
            <Button type="button" label="Save" @click="saveWine" />
        </div>
    </Dialog>
</template>

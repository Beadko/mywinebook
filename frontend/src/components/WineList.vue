<script>
import axios from 'axios'
import WineForm from './WineForm.vue'
import DeleteWine from './DeleteWine.vue'
import { store } from './store'

export default {
    name: "WineList",
    components: {
        WineForm,
        DeleteWine
    },
    data() {
        return {
            expanded_row: {},
            store,
            wines: [],
            selected: {},
            delete_dialog: false,
            wine_dialog: false,
            form_mode: 'add',
        }
    },
    computed: {
        wineTypeMap() {
            return this.store.wine_types.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        countryMap() {
            return this.store.countries.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        grapesMap() {
            return this.store.grapes.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        colourMap() {
            return this.store.colours.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        clarityMap() {
            return this.store.clarities.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        aromaMap() {
            return this.store.aromas.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        flavourMap() {
            return this.store.flavours.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        sweetnessMap() {
            return this.store.sweetnesses.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        acidityMap() {
            return this.store.acidities.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        tanninMap() {
            return this.store.bodies.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        bodyMap() {
            return this.store.bodies.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        finishMap() {
            return this.store.finishes.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        balanceMap() {
            return this.store.balances.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        severityMap() {
          return  {
                1: 'success',
                2: 'warn',
                3: 'danger'
            }
        },
        attributeMap() {
            return {
                1: 'fruity',
                2: 'vegetal',
                3: 'floral',
                4: 'earthy',
                5: 'woody',
                6: 'spicy',
                7: 'mineral',
                8: 'herbal',
                9: 'smoky'
            }
        },
        shadeMap() {
            return{
                1: '#faf8e3',
                2: '#f0efc3',
                3: '#e7e29e',
                4: '#f2f2c0',
                5: '#f5f09e',
                6: '#f3e64a',
                7: '#f1ecc6',
                8: '#f2e294',
                9: '#f6d166',
                10: '#ecb649',
                11: '#d49234',
                12: '#844927',
                13: '#fdcd69',
                14: '#fdcd69',
                15: '#e97f27',
                16: '#f3d0b2',
                17: '#f7a978',
                18: '#e37037',
                19: '#f6d5ce',
                20: '#f6d5ce',
                21: '#ef7e60',
                22: '#f7d5d6',
                23: '#ef8f93',
                24: '#ee626b',
                25: '#b51e41',
                26: '#a31c3a',
                27: '#6c1a28',
                28: '#a21d48',
                29: '#651b36',
                30: '#411526',
                31: '#cc3628',
                32: '#a11e23',
                33: '#6b1516',
                34: '#b34b28',
                35: '#8f3721',
                36: '#82241a'
            }
        }
    },
    methods: {
        ExpandRow(data) {
            if (this.expanded_row[data.id]) {
                delete this.expanded_row[data.id]
            } else {
                this.expanded_row[data.id] = true
            }
        },
        getWines() {
            axios.get("/wine")
                .then(res => {
                    this.wines = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getWineTypes() {
            axios.get("/wine_type")
                .then(res => {
                    this.store.wine_types = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getCountries() {
            axios.get("/country")
                .then(res => {
                    this.store.countries = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getGrapes() {
            axios.get("/grapes")
                .then(res => {
                    this.store.grapes = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getColours() {
            axios.get("/colour")
                .then(res => {
                    this.store.colours = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getClarities() {
            axios.get("/clarity")
                .then(res => {
                    this.store.clarities = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getAromas() {
            axios.get("/aroma")
                .then(res => {
                    this.store.aromas = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getFlavours() {
            axios.get("/flavour")
                .then(res => {
                    this.store.flavours = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getSweetnesses() {
            axios.get("/sweetness")
                .then(res => {
                    this.store.sweetnesses = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getAcidities() {
            axios.get("/acidity")
                .then(res => {
                    this.store.acidities = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getTannins() {
            axios.get("/tannin")
                .then(res => {
                    this.store.tannins = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getBodies() {
            axios.get("/body")
                .then(res => {
                    this.store.bodies = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                });
        },
        getFinishes() {
            axios.get("/finish")
                .then(res => {
                    this.store.finishes = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        getBalances() {
            axios.get("/balance")
                .then(res => {
                    this.store.balances = res.data
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`)
                })
        },
        selectWine(wn) {
            this.selected = wn
            this.form_mode = 'update'
            this.wine_dialog = true
        },
        deleteWine(wn) {
            this.selected = wn
            this.delete_dialog = true
        },
        getWineTypeName(wine) {
            return this.wineTypeMap[wine.data.wine_type]?.name || 'Unknown'
        },
        getCountryName(wine) {
            return this.countryMap[wine.data.country]?.name || 'Unknown'
        },
        getGrapesName(grapeId) {
            return this.grapesMap[grapeId]?.name
        },
        getColourName(wine) {
            return this.colourMap[wine.data.colour]?.name
        },
        getClarityName(wine) {
            return this.clarityMap[wine.data.clarity]?.name
        },
        getAromaName(wine) {
            return this.aromaMap[wine.data.aroma]?.name
        },
        getFlavourName(wine) {
            return this.flavourMap[wine.data.flavour]?.name
        },
        getSweetnessName(wine) {
            return this.sweetnessMap[wine.data.sweetness]?.name
        },
        getAcidityName(wine) {
            return this.acidityMap[wine.data.acidity]?.name
        },
        getTanninName(wine) {
            return this.tanninMap[wine.data.tannin]?.name
        },
        getBodyName(wine) {
            return this.bodyMap[wine.data.body]?.name
        },
        getFinishName(wine) {
            return this.finishMap[wine.data.finish]?.name
        },
        getBalanceName(wine) {
            return this.balanceMap[wine.data.balance]?.name
        },
        removeWine(wn) {
            this.wines = this.wines.filter(wine => wine.id !== wn)
        },
        addNewWine(wn) {
            this.wines.push(wn)
        },
        openAddWineForm() {
            this.selected = {
                name: '',
                wine_type: '',
                country: '',
                grapes:[],
                score: '',
                producer: '',
                year: null,
                alcohol: null,
                colour: '',
                clarity: '',
                aroma: '',
                flavour: '',
                sweetness: '',
                acidity: '',
                tannin: '',
                body: '',
                finish: '',
                balance: ''
            };
            this.form_mode = 'add'
            this.wine_dialog = true
        },
        handleWineAdded(wn) {
            this.addNewWine(wn)
            this.wine_dialog = false
        },
        handleWineUpdated() {
            this.wine_dialog = false
        }
    },
    async mounted() {
        this.getCountries()
        this.getWineTypes()
        this.getGrapes()
        this.getColours()
        this.getClarities()
        this.getAromas()
        this.getFlavours()
        this.getSweetnesses()
        this.getAcidities()
        this.getTannins()
        this.getBodies()
        this.getFinishes()
        this.getBalances()
        this.getWines()
    }
}
</script>

<template>
    <h1> Your Wine List</h1>
    <Button label="+ Add Wine" @click="openAddWineForm" class="m-8"/>
    <DataTable v-model:expandedRows="expanded_row" :value="wines" dataKey="id" tableStyle="min-width: 60rem" @row-click="ExpandRow($event.data)">
        <Column headerStyle="width:4rem">
            <template #body="wine">
                <i class="pi pi-chevron-right" style="color: #708090" v-if="!expanded_row[wine.data.id]" />
                <i class="pi pi-chevron-down" style="color: #708090" v-else />
            </template>
        </Column>
        <Column field="name" header="Name" />
        <Column field="wine_type" header="Type">
            <template #body="wine">
                {{ getWineTypeName(wine) }}
            </template>
        </Column>
        <Column field="country" header="Country">
            <template #body="wine">
                {{ getCountryName(wine) }}
            </template>
        </Column>
        <Column field="score" header="Score">
            <template #body="wine">
                <Rating v-model="wine.data.score" readonly />
            </template>
        </Column>
        <Column headerStyle="width:4rem">
            <template #body="item">
                <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Filter" @click="deleteWine(item.data)" />
                <Button icon="pi pi-pencil" severity="secondary" rounded text aria-label="Filter" @click="selectWine(item.data)" />
            </template>
        </Column>
        <template #expansion="wine">
            <div class="flex flex-wrap items-center p-4 gap-8">
                <div class="flex items-center" v-if="wine.data.grapes">
                    <div class="font-medium mr-2">Grapes:</div>                    
                    <div class="flex flex-wrap gap-2">
                        <Tag v-for="id in wine.data.grapes" :key="id" :value="getGrapesName(id)" rounded />
                     </div>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.producer">
                    <div class="font-medium mr-2">Producer:</div> {{ wine.data.producer }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.year">
                    <div class="font-medium mr-2">Year:</div> {{ wine.data.year }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.alcohol">
                    <div class="font-medium mr-2">Alcohol:</div> {{ wine.data.alcohol }}%
                </div>
                <div class="flex items-center p-2" v-if="wine.data.colour">
                    <div class="font-medium mr-2">Colour:</div>
                    <div class="flex items-center gap-2">
                        <i class="pi pi-circle-fill" :style="{ color: shadeMap[wine.data.colour], fontSize: '1.5rem' }" />
                        {{ getColourName(wine) }}
                    </div>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.clarity">
                    <div class="font-medium mr-2">Clarity:</div> {{ getClarityName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.aroma">
                    <div class="font-medium mr-2">Nose:</div>
                    <Tag :value="getAromaName(wine)" rounded :class="[attributeMap[wine.data.aroma], 'active']"/>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.flavour">
                    <div class="font-medium mr-2">Flavour:</div>
                    <Tag :value="getFlavourName(wine)" rounded :class="[attributeMap[wine.data.flavour], 'active']"/>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.sweetness">
                    <div class="font-medium mr-2">Sweetness:</div> {{ getSweetnessName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.acidity">
                    <div class="font-medium mr-2">Acidity:</div> {{ getAcidityName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.tannin">
                    <div class="font-medium mr-2">Tannin:</div> {{ getTanninName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.finish">
                    <div class="font-medium mr-2">Body:</div> {{ getBodyName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.finish">
                    <div class="font-medium mr-2">Finish:</div>
                    <Tag :value="getFinishName(wine)" :severity="severityMap[wine.data.finish]"/>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.balance">
                    <div class="font-medium mr-2">Balance:</div>
                    <Tag :value="getBalanceName(wine)" :severity="severityMap[wine.data.balance]"/>
                </div>
            </div>
        </template>
    </DataTable>
    <DeleteWine v-model:visible="delete_dialog" :selected="selected" @wine_deleted="removeWine" />
    <WineForm v-model:visible="wine_dialog" :selected="selected" :severity-map="severityMap" :attribute-map="attributeMap" :shade-map="shadeMap" :mode="form_mode" @wine_added="handleWineAdded" @wine_updated="handleWineUpdated" />
</template>

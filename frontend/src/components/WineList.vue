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
        attributeMap() {
            return (array) => array.reduce((acc, cur) => { acc[cur.id] = cur; return acc }, {})
        },
        wineTypeMap() {
            return this.attributeMap(this.store.wine_types)
        },
        countryMap() {
            return this.attributeMap(this.store.countries)
        },
        grapesMap() {
            return this.attributeMap(this.store.grapes)
        },
        colourMap() {
            return this.attributeMap(this.store.colours)
        },
        clarityMap() {
            return this.attributeMap(this.store.clarities)
        },
        aromaMap() {
            return this.attributeMap(this.store.aromas)
        },
        flavourMap() {
            return this.attributeMap(this.store.flavours)
        },
        sweetnessMap() {
            return this.attributeMap(this.store.sweetnesses)
        },
        acidityMap() {
            return this.attributeMap(this.store.acidities)
        },
        tanninMap() {
            return this.attributeMap(this.store.tannins)
        },
        bodyMap() {
            return this.attributeMap(this.store.bodies)
        },
        finishMap() {
            return this.attributeMap(this.store.finishes)
        },
        balanceMap() {
            return this.attributeMap(this.store.balances)
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
        async fetchData() {
            try {
                const [wines, wineTypes, countries, grapes, colours, clarities, aromas, flavours, sweetnesses, acidities, tannins, bodies, finishes, balances] = await Promise.all([
                    axios.get("/wine"),
                    axios.get("/wine_type"),
                    axios.get("/country"),
                    axios.get("/grapes"),
                    axios.get("/colour"),
                    axios.get("/clarity"),
                    axios.get("/aroma"),
                    axios.get("/flavour"),
                    axios.get("/sweetness"),
                    axios.get("/acidity"),
                    axios.get("/tannin"),
                    axios.get("/body"),
                    axios.get("/finish"),
                    axios.get("/balance")
                ])
                this.wines = wines.data;
                this.store.wine_types = wineTypes.data;
                this.store.countries = countries.data;
                this.store.grapes = grapes.data;
                this.store.colours = colours.data;
                this.store.clarities = clarities.data;
                this.store.aromas = aromas.data;
                this.store.flavours = flavours.data;
                this.store.sweetnesses = sweetnesses.data;
                this.store.acidities = acidities.data;
                this.store.tannins = tannins.data;
                this.store.bodies = bodies.data;
                this.store.finishes = finishes.data;
                this.store.balances = balances.data;
            } catch (error) {
                window.alert(`The API returned an error: ${error}`)
            }
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
        getName(map, id) {
            return map[id]?.name
        },
        getWineTypeName(wine) {
            return this.getName(this.wineTypeMap, wine.data.wine_type)
        },
        getCountryName(wine) {
            return this.getName(this.countryMap, wine.data.country)
        },
        getGrapesName(grapeId) {
            return this.getName(this.grapesMap, grapeId)
        },
        getColourName(wine) {
            return this.getName(this.colourMap, wine.data.colour)
        },
        getClarityName(wine) {
            return this.getName(this.clarityMap, wine.data.clarity)
        },
        getAromaName(wine) {
            return this.getName(this.aromaMap, wine.data.aroma)
        },
        getFlavourName(wine) {
            return this.getName(this.flavourMap, wine.data.flavour)
        },
        getSweetnessName(wine) {
            return this.getName(this.sweetnessMap, wine.data.sweetness)
        },
        getAcidityName(wine) {
            return this.getName(this.acidityMap, wine.data.acidity)
        },
        getTanninName(wine) {
            return this.getName(this.tanninMap, wine.data.tannin)
        },
        getBodyName(wine) {
            return this.getName(this.bodyMap, wine.data.body)
        },
        getFinishName(wine) {
            return this.getName(this.finishMap, wine.data.finish)
        },
        getBalanceName(wine) {
            return this.getName(this.balanceMap, wine.data.balance)
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
        this.fetchData()
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
                        <i class="pi pi-circle-fill" :style="{ color: store.shade[wine.data.colour], fontSize: '1.5rem' }" />
                        {{ getColourName(wine) }}
                    </div>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.clarity">
                    <div class="font-medium mr-2">Clarity:</div> {{ getClarityName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.aroma">
                    <div class="font-medium mr-2">Nose:</div>
                    <Tag :value="getAromaName(wine)" rounded :class="[store.characteristic[wine.data.aroma], 'active']"/>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.flavour">
                    <div class="font-medium mr-2">Flavour:</div>
                    <Tag :value="getFlavourName(wine)" rounded :class="[store.characteristic[wine.data.flavour], 'active']"/>
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
                <div class="flex items-center p-2" v-if="wine.data.body">
                    <div class="font-medium mr-2">Body:</div> {{ getBodyName(wine) }}
                </div>
                <div class="flex items-center p-2" v-if="wine.data.finish">
                    <div class="font-medium mr-2">Finish:</div>
                    <Tag :value="getFinishName(wine)" :severity="store.severity[wine.data.finish]"/>
                </div>
                <div class="flex items-center p-2" v-if="wine.data.balance">
                    <div class="font-medium mr-2">Balance:</div>
                    <Tag :value="getBalanceName(wine)" :severity="store.severity[wine.data.balance]"/>
                </div>
            </div>
        </template>
    </DataTable>
    <DeleteWine v-model:visible="delete_dialog" :selected="selected" @wine_deleted="removeWine" />
    <WineForm v-model:visible="wine_dialog" :selected="selected" :mode="form_mode" @wine_added="handleWineAdded" @wine_updated="handleWineUpdated" />
</template>

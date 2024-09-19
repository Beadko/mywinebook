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
            store,
            wines: [],
            selected: {},
            delete_dialog: false,
            wine_dialog: false,
            wine_card: {},
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
        showCard(wn) {
            if (this.wine_card[wn.id]) {
                delete this.wine_card[wn.id];
            } else {
                this.wine_card = { ...this.wine_card, [wn.id]: true };
            }
        },
        getName(map, id) {
            return map[id]?.name
        },
        getWineTypeName(wine) {
            return this.getName(this.wineTypeMap, wine.wine_type)
        },
        getCountryName(wine) {
            return this.getName(this.countryMap, wine.country)
        },
        getGrapesName(grapeId) {
            return this.getName(this.grapesMap, grapeId)
        },
        getColourName(wine) {
            return this.getName(this.colourMap, wine.colour)
        },
        getClarityName(wine) {
            return this.getName(this.clarityMap, wine.clarity)
        },
        getAromaName(wine) {
            return this.getName(this.aromaMap, wine.aroma)
        },
        getFlavourName(wine) {
            return this.getName(this.flavourMap, wine.flavour)
        },
        getSweetnessName(wine) {
            return this.getName(this.sweetnessMap, wine.sweetness)
        },
        getAcidityName(wine) {
            return this.getName(this.acidityMap, wine.acidity)
        },
        getTanninName(wine) {
            return this.getName(this.tanninMap, wine.tannin)
        },
        getBodyName(wine) {
            return this.getName(this.bodyMap, wine.body)
        },
        getFinishName(wine) {
            return this.getName(this.finishMap, wine.finish)
        },
        getBalanceName(wine) {
            return this.getName(this.balanceMap, wine.balance)
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
    <div class="card w-full">
        <div class="flex flex-row justify-between gap-4 my-6">
            <div class="flex items-center justify-start">
                <span class="material-symbols-outlined text-4xl">wine_bar</span>
                <div class="text-xl font-large ml-2">My Wine Book</div>
            </div>
            <Button icon="pi pi-plus" size="small" @click="openAddWineForm" />
        </div>
        <DataView :value="wines">
            <template #list="slotProps">
                <div class="flex flex-col">
                    <div v-for="item in slotProps.items" :key="item.id" @click="showCard(item)">
                        <div class="flex flex-col shadow-md rounded hover:border p-4">                       
                            <div class="flex flex-row justify-between items-center">
                                <div class="flex justify-around">
                                    <div class="flex flex-col px-6">
                                        <span class="text-xl font-bold">{{ item.name }}</span>
                                        <div class="text-lg font-semibold mt-2">{{ getWineTypeName(item) }}</div>
                                        <span class="text-md font-medium">{{ getCountryName(item) }}</span>
                                    </div>
                                </div>
                                <div class="flex flex-row justify-end items-center">
                                    <div class="px-6">
                                        <div class="bg-surface-100 p-1" style="border-radius: 30px" >
                                        <div class="bg-surface-0 flex items-center gap-2 justify-center py-1 px-2" style="border-radius: 20px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                                            <span class="text-surface-900 font-medium text-sm">{{ item.score }}</span>
                                            <i class="pi pi-star-fill text-yellow-500"></i>
                                        </div>
                                    </div>
                                    </div>
                                    <div class="flex flex-col">
                                        <Button icon="pi pi-pencil" severity="secondary" rounded text aria-label="Edit" @click="selectWine(item)" />
                                        <Button icon="pi pi-trash" severity="secondary" rounded text aria-label="Delete" @click="deleteWine(item)" />
                                    </div>
                                </div>
                            </div>
                            <div v-if="wine_card[item.id]" class="mt-4">
                                <div class="flex flex-wrap items-center p-4 gap-8">
                                    <div class="flex items-center" v-if="item.grapes">
                                        <div class="font-medium mr-2">Grapes:</div>                    
                                        <div class="flex flex-wrap gap-2">
                                            <Tag v-for="id in item.grapes" :key="id" :value="getGrapesName(id)" rounded />
                                        </div>
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.producer">
                                        <div class="font-medium mr-2">Producer:</div> {{ item.producer }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.year">
                                        <div class="font-medium mr-2">Year:</div> {{ item.year }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.alcohol">
                                        <div class="font-medium mr-2">Alcohol:</div> {{ item.alcohol }}%
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.colour">
                                        <div class="font-medium mr-2">Colour:</div>
                                        <div class="flex items-center gap-2">
                                            <i class="pi pi-circle-fill" :style="{ color: store.shade[item.colour], fontSize: '1.5rem' }" />
                                            {{ getColourName(item) }}
                                        </div>
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.clarity">
                                        <div class="font-medium mr-2">Clarity:</div> {{ getClarityName(item) }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.aroma">
                                        <div class="font-medium mr-2">Nose:</div>
                                        <Tag :value="getAromaName(item)" rounded :class="[store.characteristic[item.aroma], 'active']"/>
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.flavour">
                                        <div class="font-medium mr-2">Flavour:</div>
                                        <Tag :value="getFlavourName(item)" rounded :class="[store.characteristic[item.flavour], 'active']"/>
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.sweetness">
                                        <div class="font-medium mr-2">Sweetness:</div> {{ getSweetnessName(item) }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.acidity">
                                        <div class="font-medium mr-2">Acidity:</div> {{ getAcidityName(item) }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.tannin">
                                        <div class="font-medium mr-2">Tannin:</div> {{ getTanninName(item) }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.body">
                                        <div class="font-medium mr-2">Body:</div> {{ getBodyName(item) }}
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.finish">
                                        <div class="font-medium mr-2">Finish:</div>
                                        <Tag :value="getFinishName(item)" :severity="store.severity[item.finish]"/>
                                    </div>
                                    <div class="flex items-center p-2" v-if="item.balance">
                                        <div class="font-medium mr-2">Balance:</div>
                                        <Tag :value="getBalanceName(item)" :severity="store.severity[item.balance]"/>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </DataView>
    </div>
    <DeleteWine v-model:visible="delete_dialog" :selected="selected" @wine_deleted="removeWine" />
    <WineForm v-model:visible="wine_dialog" :selected="selected" :mode="form_mode" @wine_added="handleWineAdded" @wine_updated="handleWineUpdated" class="w-full md:w-1/2 lg:w-1/3" />
</template>

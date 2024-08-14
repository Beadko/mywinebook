import { createApp } from 'vue'
import './style.css'
import 'primeicons/primeicons.css'
import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'
import Column from 'primevue/column'
import DataTable from 'primevue/datatable'
import Button from 'primevue/button'
import Select from 'primevue/select'
import InputText from 'primevue/inputtext'
import Dialog from 'primevue/dialog'
import Rating from 'primevue/rating'
import App from './App.vue'


const app = createApp(App);
app.use(PrimeVue, {
    theme: {
        preset: Aura,
    }
});
app.component('Column', Column)
app.component('DataTable', DataTable)
app.component('Button', Button)
app.component('Select', Select)
app.component('InputText', InputText)
app.component('Dialog', Dialog)
app.component('Rating', Rating)


app.mount('#app')

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
import InputNumber from 'primevue/inputnumber'
import Dialog from 'primevue/dialog'
import Rating from 'primevue/rating'
import Panel from 'primevue/panel'
import RadioButton from 'primevue/radiobutton'
import Tag from 'primevue/tag'
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
app.component('InputNumber', InputNumber)
app.component('Dialog', Dialog)
app.component('Rating', Rating)
app.component('Panel', Panel)
app.component('RadioButton', RadioButton)
app.component('Tag', Tag)

app.mount('#app')

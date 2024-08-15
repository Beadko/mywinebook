<script>
import axios from 'axios';

export default {
    name: "DeleteWine",
    props: {
        selected: Object,
    },

    methods: {
        confirmDelete() {
            axios.delete("/wine/"+ this.selected.id)
            .then(
                this.$parent.delete_dialog = false,
                location.reload()
            )
            .catch((error) => {
                window.alert(`The API returned an error: ${error}`)
            })
        },
    }
}
</script>

<template>
    <Dialog :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="selected">
                Are you sure you want to delete <b>{{ this.selected.name }}</b>?
            </span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="this.$parent.delete_dialog = false" />
            <Button label="Yes" icon="pi pi-check" @click="confirmDelete" />
        </template>
    </Dialog>
</template>
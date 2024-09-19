<script>
import axios from "axios";
import { store } from "./store";

export default {
  name: "WineTypes",
  inheritAttrs: false,
  props: {
    selected: Object,
  },
  data() {
    return {
      dialog_visible: false,
      wine_types: [],
      new_type: {
        name: "",
      },
      store,
    };
  },
  methods: {
    addWineType() {
      axios
        .post("/wine_type", this.new_type)
        .then((res) => {
          this.store.wine_types.push(res.data),
            (this.selected.wine_type = res.data.id),
            (this.dialog_visible = false);
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
    handleTypeChange(value) {
      if (value === "add_type") {
        (this.dialog_visible = true), (this.selected.wine_type = "");
      }
    },
    clearSelection() {
      (this.selected.wine_type = ""), (this.dialog_visible = false);
    },
  },
};
</script>

<template>
  <div class="flex items-center gap-2 mb-4">
    <label for="wine_type" class="font-semibold w-20">Type</label>
    <Select
      v-model="selected.wine_type"
      :options="[...store.wine_types, { id: 'add_type', name: '+ Add New' }]"
      optionLabel="name"
      optionValue="id"
      class="w-full md:w-[14rem]"
      @change="handleTypeChange($event.value)"
    />
  </div>
  <Dialog
    v-model:visible="dialog_visible"
    modal
    header="Add a new wine type"
    :style="{ width: '25rem' }"
  >
    <div class="flex items-center gap-2 mb-4">
      <label for="name" class="font-semibold w-20">Wine Type</label>
      <InputText
        v-model="new_type.name"
        id="name"
        class="w-full md:w-[14rem]"
        autocomplete="off"
      />
    </div>
    <div class="flex justify-end gap-2">
      <Button
        type="button"
        label="Cancel"
        severity="secondary"
        text
        @click="clearSelection"
      />
      <Button type="button" label="Save" @click="addWineType" />
    </div>
  </Dialog>
</template>

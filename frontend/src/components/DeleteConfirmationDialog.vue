<script setup>

import {onUpdated, ref} from "vue";

const props = defineProps({
  visible: Boolean
})

let dialog = ref(false)

onUpdated(()=>{
  dialog.value = props.visible
})
</script>

<template>
  <q-dialog v-model="dialog" persistent @before-hide="$emit('hide')">
    <q-card>
      <q-card-section class="row items-center">
        <q-icon name="delete_outline" size="sm" color="red" />
        <span class="q-ml-sm text-bold">¿Está seguro que desea eliminar?</span>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn icon="cancel" dense outlined label="Cancel" color="grey" v-close-popup />
        <q-btn icon="check" dense outlined @click="$emit('confirm')" label="Confirm" color="primary" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<style scoped>

</style>

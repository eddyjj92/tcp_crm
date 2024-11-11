<template>
  <q-page class="q-pa-sm">
    <q-card>
      <q-table
        ref="tableRef"
        title="Productos"
        :rows="suppliers"
        :hide-header="mode === 'grid'"
        :columns="columns"
        row-key="name"
        :grid="mode === 'grid'"
        :filter="filter"
        :pagination:sync="pagination"
        class="my-sticky-last-column-table"
      >
        <template v-slot:top-right="props">
          <q-btn @click="dialog = true;" icon="add" outline color="primary" label="Registrar" class="q-mr-xs"/>

          <q-input outlined dense debounce="300" v-model="filter" placeholder="Buscar">
            <template v-slot:append>
              <q-icon name="search"/>
            </template>
          </q-input>

          <q-btn
            flat
            round
            dense
            :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
            @click="props.toggleFullscreen"
            v-if="mode === 'list'"
          >
            <q-tooltip
              :disable="$q.platform.is.mobile"
              v-close-popup
            >{{ props.inFullscreen ? 'Exit Fullscreen' : 'Toggle Fullscreen' }}
            </q-tooltip>
          </q-btn>

          <q-btn
            flat
            round
            dense
            :icon="mode === 'grid' ? 'list' : 'grid_on'"
            @click="mode = mode === 'grid' ? 'list' : 'grid'; separator = mode === 'grid' ? 'none' : 'horizontal'"
            v-if="!props.inFullscreen"
          >
            <q-tooltip
              :disable="$q.platform.is.mobile"
              v-close-popup
            >{{ mode === 'grid' ? 'List' : 'Grid' }}
            </q-tooltip>
          </q-btn>

          <q-btn
            color="primary"
            icon-right="archive"
            label="Exportar a csv"
            no-caps
            @click="exportTable"
          />
        </template>
        <template v-slot:header="props">
          <q-tr :props="props">
            <q-th
              v-for="col in props.cols"
              :key="col.name"
              :props="props"
              auto-width
            >
              <q-checkbox :disable="suppliers.length === 0" class="float-left" v-if="col.label === 'ID'" v-model="checkAll"/>
              <span class="float-left" v-else>{{ col.label }}<br>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.field === 'name'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.address" v-if="col.field === 'address'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.phone" v-if="col.field === 'phone'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.email" v-if="col.field === 'email'" type="text"/>
              </span>
            </q-th>
            <q-th auto-width align="center">Opciones</q-th>
          </q-tr>
        </template>
        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td
              v-for="col in props.cols"
              :key="col.name"
              :props="props"
              auto-width
            >
              <q-checkbox v-model="selected" :val="col.value" class="float-left" v-if="col.label === 'ID'"/>
              <span v-if="col.field !== 'id'" class="float-left">{{ col.value }}</span>
            </q-td>
            <q-td auto-width align="center">
              <q-btn @click="edit(props.cols[0].value)" class="q-ml-sm bg-yellow-6 text-white" style="width: 35px">
                <q-icon size="xs" name="edit" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Editar Proveedor.
                </q-tooltip>
              </q-btn>
              <q-btn  @click="deleteDialogOpen(props.cols[0].value)" class="q-ml-sm bg-negative text-white" style="width: 35px">
                <q-icon size="xs" name="delete_forever" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Eliminar Proveedor.
                </q-tooltip>
              </q-btn>
            </q-td>
          </q-tr>
          <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" group="primary" />
          </q-inner-loading>
        </template>
      </q-table>
    </q-card>
    <q-dialog v-model="dialog" persistent @before-hide="selected = []; product = {}">
      <q-card style="width: 700px; max-width: 85vw;">
        <q-card-section>
          <div class="text-h6">
            <span v-if="selected.length === 0"> Registrar Proveedor</span>
            <span v-else> Editar Proveedor</span>
            <q-btn round flat dense icon="close" class="float-right" group="grey-8" v-close-popup></q-btn>
          </div>
        </q-card-section>
        <q-separator inset></q-separator>
        <q-card-section style="max-height: 65vh" class="scroll">
          <q-form class="q-gutter-md">
            <q-list>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Nombre</q-item-label>
                  <q-input dense outlined v-model="supplier.name" label="Nombre"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Dirección</q-item-label>
                  <q-input dense outlined v-model="supplier.address" label="Dirección"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Teléfono</q-item-label>
                  <q-input dense outlined v-model="supplier.phone" label="Teléfono"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Correo Electrónico</q-item-label>
                  <q-input dense outlined v-model="supplier.email" label="Correo Electrónico"/>
                </q-item-section>
              </q-item>
            </q-list>
          </q-form>
        </q-card-section>
        <q-separator inset></q-separator>
        <q-card-actions align="right" class="text-teal">
          <q-btn v-if="selected.length === 0" @click="register" icon="save" label="Guardar" type="submit" color="primary"/>
          <q-btn v-else @click="update" icon="save" label="Actualizar" type="submit" color="primary"/>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
  <delete-confirmation-dialog :visible="deleteDialog !== null" @hide="deleteDialogClose" @confirm="remove"/>
</template>

<script setup>
import {exportFile, useQuasar} from "quasar";
import {onMounted, reactive, ref, watch} from 'vue'
import {storeToRefs} from "pinia";
import DeleteConfirmationDialog from "components/DeleteConfirmationDialog.vue";
import {useProductStore} from "stores/product-store.js";
import {useSupplierStore} from "stores/supplier-store.js";

const $q = useQuasar();
const supplierStore = useSupplierStore();
const {suppliers} = storeToRefs(supplierStore);
let loading = ref(false);

const columns = [
  { name: "id", align: "left", label: "ID",  field: "id", sortable: true },
  { name: "name", align: "left", label: "Nombre",  field: "name", sortable: true },
  { name: "address", align: "left", label: "Dirección",  field: "address", sortable: true },
  { name: "phone", align: "left", label: "Teléfono",  field: "phone", sortable: true },
  { name: "email", align: "left", label: "Correo",  field: "email", sortable: true, format: (val, row) => val ? val : '-', },
];

const tableRef = ref()
const selected = ref([])
let filter = ref("")
let filters = reactive({
  name: "",
  address: "",
  phone: "",
  email: ""
});
let supplier = ref({})
let dialog = ref(false);
let mode = ref("list")
let pagination = ref({
  rowsPerPage: 10
});

const checkAll = ref(false)
watch(checkAll, () => {
  if (checkAll.value && suppliers.value.length !== selected.value.length){
    suppliers.value.forEach(el => {
      selected.value.push(el.id)
    })
  }else if (!checkAll.value){
    selected.value = []
  }
})

watch(selected, () => {
  checkAll.value = suppliers.value.length === selected.value.length && suppliers.value.length > 0;
})

onMounted(async () => {
  await supplierStore.getSuppliers(null)
})

function wrapCsvValue(val, formatFn) {
  let formatted = formatFn !== void 0 ? formatFn(val) : val;

  formatted =
    formatted === void 0 || formatted === null ? "" : String(formatted);

  formatted = formatted.split('"').join('""');

  return `"${formatted}"`;
}
const exportTable = () => {
  // naive encoding to csv format
  const content = [columns.map(col => wrapCsvValue(col.label))]
    .concat(
      suppliers.value.map(row =>
        columns
          .map(col =>
            wrapCsvValue(
              typeof col.field === "function"
                ? col.field(row)
                : row[col.field === void 0 ? col.name : col.field],
              col.format
            )
          )
          .join(",")
      )
    )
    .join("\r\n");

  const status = exportFile("proveedores.csv", content, "text/csv");

  if (status !== true) {
    $q.notify({
      message: "Browser denied file download...",
      group: "negative",
      icon: "warning"
    });
  }
}

watch(filters, async () => {
  loading.value = true;
  const url = `/api/suppliers?name=${filters.name}&description=${filters.description}`
  await supplierStore.getSuppliers(url, false)
    .finally(() => loading.value = false);
})

const register = async () => {
  const registered = await supplierStore.postSupplier(supplier.value)
  if (registered){
    dialog.value = false;
    supplier.value = {};
  }
}

const edit = (id) => {
  const edit = suppliers.value.find(u => u.id === id);
  supplier.value.name = edit.name;
  supplier.value.address = edit.address;
  supplier.value.phone = edit.phone;
  supplier.value.email = edit.email;
  selected.value = [id];
  dialog.value = true;
}

const update = async () => {
  const updated = await supplierStore.putSupplier(selected.value[0], supplier.value)
  if (updated){
    dialog.value = false;
    selected.value = [];
  }
}

let deleteDialog = ref(null)
const deleteDialogOpen = (id) => {
  deleteDialog.value = id;
}

const deleteDialogClose = () => {
  deleteDialog.value = null;
}

const remove = async () => {
  const deleted = await supplierStore.deleteSupplier(deleteDialog.value)
  if (deleted){
    deleteDialog.value = null;
  }
}

</script>


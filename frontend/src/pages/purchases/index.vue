<template>
  <q-page class="q-pa-sm">
    <q-card>
      <q-table
        ref="tableRef"
        title="Compras"
        :rows="purchases"
        :hide-header="mode === 'grid'"
        :columns="columns"
        row-key="name"
        :grid="mode === 'grid'"
        :filter="filter"
        :pagination:sync="pagination"
        class="my-sticky-last-column-table"
      >
        <template v-slot:top-right="props">
          <q-btn to="/purchases/create" icon="add" outline color="primary" label="Registrar" class="q-mr-xs"/>

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
              :style="{ width: col.width + 'px' }"
            >
              <q-checkbox :disable="purchases.length === 0" class="float-left" v-if="col.label === 'ID'" v-model="checkAll"/>
              <span class="float-left text-bold" style="font-size: 11px" v-else>{{ col.label }}<br>
                <q-input v-if="col.field === 'date'" outlined dense v-model="filters.date" mask="##/##/####" debounce="300" style="min-width: 100px">
                  <template v-slot:append>
                    <q-icon size="xs" name="event" class="cursor-pointer q-mr-xs">
                      <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                        <q-date v-model="filters.date" mask="DD/MM/YYYY">
                          <div class="row items-center justify-end">
                            <q-btn v-close-popup label="Close" color="primary" flat />
                          </div>
                        </q-date>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
                <q-input v-if="col.field === 'total_units'" debounce="300" outlined dense v-model="filters.total_units" type="text" style="min-width: 60px"/>
                <q-input v-if="col.field === 'total_price'" debounce="300" outlined dense v-model="filters.total_price" type="text" style="min-width: 60px"/>
                <q-select v-if="col.field === 'supplier'" outlined dense option-label="name" option-value="id"
                            :options="[{name: 'Todos', id: null}, ...suppliers]"
                            v-model="filters.supplier_id" style="min-width: 100px"
                            type="text" emit-value map-options
                >
                  <template v-slot:selected-item="{ opt }">
                    <span style="margin-bottom: 20px;">{{ opt.name }}</span>
                  </template>
                </q-select>
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
              <q-btn class="q-ml-sm bg-yellow-6 text-white" style="width: 35px">
                <q-icon size="xs" name="edit" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Editar Compra.
                </q-tooltip>
              </q-btn>
              <q-btn @click="deleteDialogOpen(props.cols[0].value)" class="q-ml-sm bg-negative text-white" style="width: 35px">
                <q-icon size="xs" name="delete_forever" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Eliminar Compra.
                </q-tooltip>
              </q-btn>
            </q-td>
          </q-tr>
          <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
          </q-inner-loading>
        </template>
      </q-table>
    </q-card>
  </q-page>
  <delete-confirmation-dialog :visible="deleteDialog !== null" @hide="deleteDialogClose" @confirm="remove"/>
</template>

<script setup>
import {date, exportFile, useQuasar} from "quasar";
import {onMounted, reactive, ref, watch} from 'vue'
import {storeToRefs} from "pinia";
import DeleteConfirmationDialog from "components/DeleteConfirmationDialog.vue";
import {usePurchaseStore} from "stores/purchase-store.js";
import {useSupplierStore} from "stores/supplier-store.js";

const $q = useQuasar();
const purchaseStore = usePurchaseStore();
const {purchases} = storeToRefs(purchaseStore);
const supplierStore = useSupplierStore()
const {suppliers} = storeToRefs(supplierStore)

let loading = ref(false);

const columns = [
  { name: "id", align: "left", label: "ID",  field: "id", sortable: true, },
  { name: "date", align: "left", label: "Fecha",  field: "date", sortable: true, format: (val, row) => date.formatDate(new Date(val), "DD/MM/YYYY"), },
  { name: "total_units", align: "left", label: "Unidades",  field: "total_units", sortable: true },
  { name: "total_price", align: "left", label: "Total",  field: "total_price", sortable: true },
  { name: "supplier", align: "left", label: "Proveedor",  field: "supplier", sortable: true, format: (val, row) => val.name },
];

const tableRef = ref()
const selected = ref([])
let filter = ref("")
let filters = reactive({
  date: null,
  total_units: null,
  total_price: null,
  supplier_id: null,
});
let purchase = ref({})
let dialog = ref(false);
let mode = ref("list")
let pagination = ref({
  rowsPerPage: 10
});

const checkAll = ref(false)
watch(checkAll, () => {
  if (checkAll.value && purchases.value.length !== selected.value.length){
    purchases.value.forEach(p => {
      selected.value.push(p.id)
    })
  }else if (!checkAll.value){
    selected.value = []
  }
})

watch(selected, () => {
  checkAll.value = purchases.value.length === selected.value.length && purchases.value.length > 0;
})

onMounted(async () => {
  await purchaseStore.getPurchases(null, !purchaseStore.fetched)
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
      purchases.value.map(row =>
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

  const status = exportFile("compras.csv", content, "text/csv");

  if (status !== true) {
    $q.notify({
      message: "Browser denied file download...",
      color: "negative",
      icon: "warning"
    });
  }
}

watch(filters, async () => {
  loading.value = true;
  const url = `/api/purchases?name=${filters.name}`
  await purchaseStore.getPurchases(url, false)
    .finally(() => loading.value = false);
})

let deleteDialog = ref(null)
const deleteDialogOpen = (id) => {
  deleteDialog.value = id;
}

const deleteDialogClose = () => {
  deleteDialog.value = null;
}

const remove = async () => {
  const deleted = await purchaseStore.deletePurchase(deleteDialog.value)
  if (deleted){
    deleteDialog.value = null;
  }
}

</script>
<style scoped>
.header-cell {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.resizer {
  width: 5px;
  height: 100%;
  cursor: col-resize;
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.1);
}
</style>

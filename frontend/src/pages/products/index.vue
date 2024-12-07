<template>
  <q-page class="q-pa-sm">
    <q-card>
      <q-table
        ref="tableRef"
        title="Productos"
        :rows="products"
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
              <q-checkbox :disable="products.length === 0" class="float-left" v-if="col.label === 'ID'" v-model="checkAll"/>
              <span class="float-left" v-else>{{ col.label }}<br>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.field === 'name'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.description" v-if="col.field === 'description'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.purchase_price" v-if="col.field === 'purchase_price'" type="text"/>
                <q-input debounce="300" outlined dense v-model="filters.sale_price" v-if="col.field === 'sale_price'" type="text"/>
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
              <q-img class="rounded-borders" width="40px" :ratio="1" v-if="col.field === 'image_path'" :src="`http://localhost:3000/${props.cols[1].value}`" />
              <span v-if="col.field !== 'id' && col.field !== 'image_path' && col.field !== 'active'" class="float-left">{{ col.value }}</span>
            </q-td>
            <q-td auto-width align="center">
              <q-btn @click="edit(props.cols[0].value)" class="q-ml-sm bg-yellow-6 text-white" style="width: 35px">
                <q-icon size="xs" name="edit" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Editar Producto.
                </q-tooltip>
              </q-btn>
              <q-btn  @click="deleteDialogOpen(props.cols[0].value)" class="q-ml-sm bg-negative text-white" style="width: 35px">
                <q-icon size="xs" name="delete_forever" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Eliminar Producto.
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
            <span v-if="selected.length === 0"> Registrar Producto</span>
            <span v-else> Editar Producto</span>
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
                  <q-input dense outlined v-model="product.name" label="Nombre"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Descripción</q-item-label>
                  <q-input dense outlined v-model="product.description" label="Descripción"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Precio de Compras</q-item-label>
                  <q-input dense outlined v-model="product.purchase_price" label="Precio de Compras"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Precio de Venta</q-item-label>
                  <q-input dense outlined v-model="product.sale_price" label="Precio de Venta"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Imagen</q-item-label>
                  <q-file
                    dense
                    outlined
                    v-model="uploadImage"
                    @update:model-value="handleFileAdded"
                    label="Imagen"
                  >
                    <template v-slot:prepend>
                      <q-icon name="cloud_upload" @click.stop.prevent />
                    </template>
                    <template v-slot:append>
                      <q-btn
                        :disable="uploadImage === null || uploadImage === undefined"
                        outlined
                        unelevated
                        icon="close"
                        @click.stop.prevent="uploadImage = null"
                        class="cursor-pointer"
                        style="margin-right: -10px; width: 40px"
                      />
                    </template>
                  </q-file>
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

const $q = useQuasar();
const productStore = useProductStore();
const {products} = storeToRefs(productStore);
let loading = ref(false);

const columns = [
  { name: "id", align: "left", label: "ID",  field: "id", sortable: true },
  { name: "image_path", align: "left", label: "Imagen",  field: "image_path", sortable: true },
  { name: "name", align: "left", label: "Producto",  field: "name", sortable: true },
  { name: "description", align: "left", label: "Descripción",  field: "description", sortable: true },
  { name: "purchase_price", align: "left", label: "Precio de Compra",  field: "purchase_price", sortable: true, format: (val, row) => `$${val}`, },
  { name: "sale_price", align: "left", label: "Precio de Venta",  field: "sale_price", sortable: true, format: (val, row) => `$${val}`, },
];

const tableRef = ref()
const selected = ref([])
let filter = ref("")
let filters = reactive({
  name: "",
  description: "",
  purchase_price: "",
  sale_price: ""
});
let product = ref({})
let uploadImage = ref(null);
let dialog = ref(false);
let mode = ref("list")
let pagination = ref({
  rowsPerPage: 10
});

const checkAll = ref(false)
watch(checkAll, () => {
  if (checkAll.value && products.value.length !== selected.value.length){
    products.value.forEach(el => {
      selected.value.push(el.id)
    })
  }else if (!checkAll.value){
    selected.value = []
  }
})

watch(selected, () => {
  checkAll.value = products.value.length === selected.value.length && products.value.length > 0;
})

onMounted(async () => {
  await productStore.getProducts(null)
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
      products.value.map(row =>
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

  const status = exportFile("products.csv", content, "text/csv");

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
  const url = `/api/items?name=${filters.name}&description=${filters.description}`
  await productStore.getProducts(url, false)
    .finally(() => loading.value = false);
})

const register = async () => {
  const registered = await productStore.postProduct(product.value)
  if (registered){
    dialog.value = false;
    product.value = {};
  }
}

const edit = (id) => {
  const edit = products.value.find(u => u.id === id);
  product.value.name = edit.name;
  product.value.description = edit.description;
  product.value.purchase_price = edit.purchase_price;
  product.value.sale_price = edit.sale_price;
  selected.value = [id];
  dialog.value = true;
}

const update = async () => {
  const updated = await productStore.putProduct(selected.value[0], product.value)
  if (updated){
    dialog.value = false;
    loading.value = true;
    selected.value = [];
    await productStore.getProducts(null, false)
      .finally(() => loading.value = false);
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
  const deleted = await productStore.deleteProduct(deleteDialog.value)
  if (deleted){
    deleteDialog.value = null;
    loading.value = true;
    await productStore.getProducts(null, false)
      .finally(() => loading.value = false);
  }
}

const handleFileAdded = async (file) => {
  if (file) {
    product.value.image_path = await toBase64(file);
    alert(product.value.image_path)
  }
};

function toBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = reject;
    reader.readAsDataURL(file); // Convertir el archivo a base64
  });
}

</script>


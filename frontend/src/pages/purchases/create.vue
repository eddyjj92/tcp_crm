<template>
  <q-page class="q-pa-sm full-height row">
    <q-card class="flex-grow col-auto q-px-none" style="max-height: 180px">
      <q-card-section class="full-width  q-pt-sm q-pb-xs">
        <span class="text-h6">Nueva Compra</span>
      </q-card-section>
      <q-card-section class="q-pt-none q-px-none">
        <q-item class="q-py-xs">
          <q-item-section style="min-width: 25%">
            <label for="" class="text-bold" style="font-size: 12px">Proveedor</label>
            <q-select
              :ref="setInputRef('purchase.entity_id')"
              outlined
              v-model="purchase.entity_id"
              :options="suppliers"
              option-label="trade_name"
              option-value="id"
              emit-value
              map-options
              dense
              @keyup.enter="purchase.entity_id ? focusNext('purchase.invoice') : null"
            >
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-italic text-grey">
                    No existen proveedores.
                  </q-item-section>
                </q-item>
              </template>
            </q-select>
          </q-item-section>
          <q-item-section>
            <label for="" class="text-bold" style="font-size: 12px">Factura</label>
            <q-input :ref="setInputRef('purchase.invoice')" outlined v-model="purchase.invoice" dense @keyup.enter="purchase.invoice ? focusNext('purchase.payment_shape') : null"/>
          </q-item-section>
          <q-item-section>
            <label for="" class="text-bold" style="font-size: 12px">Forma de Pago</label>
            <q-select
              :ref="setInputRef('purchase.payment_shape')"
              outlined
              v-model="purchase.payment_shape"
              :options="payment_shapes"
              map-options
              emit-value
              dense
              @keyup.enter="purchase.payment_shape ? focusNext('purchase.term') : null"
              @update:model-value="purchase.payment_shape ? focusNext('purchase.term') : null"
            />
          </q-item-section>
          <q-item-section>
            <label for="" class="text-bold" style="font-size: 12px">Plazo</label>
            <q-input :ref="setInputRef('purchase.term')" outlined v-model="purchase.term" dense :disabled="purchase.payment_shape === 1"/>
          </q-item-section>
          <q-item-section>
            <label for="" class="text-bold" style="font-size: 12px">Fecha</label>
            <q-input :ref="setInputRef('purchase.invoice_date')" outlined v-model="purchase.invoice_date" type="date" dense/>
          </q-item-section>
        </q-item>
        <q-item class="q-py-xs">
          <q-item-section style="min-width: 25%">
            <label for="" class="text-bold" style="font-size: 12px">Producto</label>
            <q-select
              :ref="setInputRef('detailInProcess.item_id')"
              outlined
              v-model="detailInProcess.item_id"
              :options="products"
              option-label="name"
              option-value="id"
              emit-value
              map-options
              dense
              use-input
              @click="detailInProcess.item_id = null;detailInProcess.color_id = null; detailInProcess.iva = null; detailInProcess.final_purchase_price = 0"
              @filter="productFilter"
            >
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-italic text-grey">
                    No existen proveedores.
                  </q-item-section>
                </q-item>
              </template>
              <template v-slot:before-options>
                <q-markup-table separator="cell" flat bordered dense>
                  <thead>
                  <tr class="bg-grey-4">
                    <th class="text-left" style="width: 120px">Producto</th>
                    <th class="text-left" style="width: 200px">Descripción</th>
                    <th class="text-left" style="width: 50px">% IVA</th>
                    <th class="text-left" style="width: 50px">P. Compra</th>
                    <th class="text-left" style="width: 50px">P. Venta</th>
                  </tr>
                  </thead>
                  <tbody>
                  <template v-for="option in products" :key="option.id">
                    <tr @click="detailInProcess.item_id = option.id;refs['detailInProcess.item_id'].hidePopup()">
                      <td style="width: 120px">{{ option.name }}</td>
                      <td style="width: 200px">{{ option.description }}</td>
                      <td style="width: 50px">{{ option.iva.percent }}%</td>
                      <td style="width: 50px">{{ option.last_purchase_price.toLocaleString('en', { style: 'currency', currency: 'USD' }) }}</td>
                      <td style="width: 50px">{{ option.last_sale_price.toLocaleString('en', { style: 'currency', currency: 'USD' }) }}</td>
                    </tr>
                  </template>
                  </tbody>
                </q-markup-table>
              </template>
              <template v-slot:option></template>
            </q-select>
          </q-item-section>
          <q-item-section style="min-width: 14%">
            <q-item class="q-pa-none">
              <q-item-section style="min-width: 60%">
                <label for="" class="text-bold" style="font-size: 12px">Color</label>
                <q-select
                  :ref="setInputRef('detailInProcess.color_id')"
                  outlined
                  v-model="detailInProcess.color_id"
                  :options="products.find(i => i.id === detailInProcess.item_id)?.colors "
                  option-label="name"
                  option-value="id"
                  emit-value
                  map-options
                  dense
                  :disable="!detailInProcess.item_id"
                  :class="!detailInProcess.item_id ? 'bg-grey-3' : ''"
                />
              </q-item-section>
              <q-item-section>
                <label for="" class="text-bold" style="font-size: 12px">% IVA</label>
                <q-input outlined dense v-model="detailInProcess.iva_percent" :disable="!detailInProcess.item_id" :class="!detailInProcess.item_id ? 'bg-grey-3' : ''"/>
              </q-item-section>
            </q-item>
          </q-item-section>
          <q-item-section  style="min-width: 20%">
            <q-item class="q-pa-none">
              <q-item-section>
                <label for="" class="text-bold" style="font-size: 12px">Precio Compra</label>
                <q-input outlined dense v-model="detailInProcess.purchase_price" :disable="!detailInProcess.item_id" :class="!detailInProcess.item_id ? 'bg-grey-3' : ''"/>
              </q-item-section>
              <q-item-section style="min-width: 55%">
                <label for="" class="text-bold" style="font-size: 12px">Precio Venta + IVA</label>
                <q-input outlined dense v-model="detailInProcess.sale_price" :disable="!detailInProcess.item_id" :class="!detailInProcess.item_id ? 'bg-grey-3' : ''"/>
              </q-item-section>
            </q-item>
          </q-item-section>
          <q-item-section style="min-width: 28%">
            <q-item class="q-pa-none">
              <q-item-section>
                <label for="" class="text-bold" style="font-size: 12px">% Desc.</label>
                <q-input outlined dense v-model="detailInProcess.discount_percent"/>
              </q-item-section>
              <q-item-section style="min-width: 40%">
                <label for="" class="text-bold" style="font-size: 12px">Precio comp. final*</label>
                <q-input outlined v-model="detailInProcess.final_purchase_price" dense/>
              </q-item-section>
              <q-item-section style="min-width: 38%">
                <label for="" class="text-bold" style="font-size: 12px">Precio comp. + iva*</label>
                <q-input outlined dense/>
              </q-item-section>
            </q-item>
          </q-item-section>
          <q-item-section>
            <q-item class="q-pa-none">
              <q-item-section>
                <q-btn @click="openSizesDialog" color="primary" icon="pan_tool_alt" dense class="q-mt-md">
                  <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                    <strong>Tallas</strong>
                    (<q-icon name="keyboard_arrow_up"/>)
                  </q-tooltip>
                </q-btn>
              </q-item-section>
              <q-item-section>
                <label for="" class="text-bold" style="font-size: 12px">Cantidad</label>
                <q-input outlined dense class="input-box"/>
              </q-item-section>
            </q-item>
          </q-item-section>
        </q-item>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-sm q-pa-none flex-grow">
      <q-card-section class="full-width scroll q-pa-none">
        <q-markup-table separator="cell" flat bordered dense>
          <thead>
            <tr class="bg-grey-4">
              <th class="text-left">Producto</th>
              <th class="text-right">Descripcion</th>
              <th class="text-right">Talla</th>
              <th class="text-right">Color</th>
              <th class="text-right">% Iva</th>
              <th class="text-right">P/Compra</th>
              <th class="text-right">P/Venta</th>
              <th class="text-right">% Desc</th>
              <th class="text-right">P/Final</th>
              <th class="text-right">P/Iva</th>
              <th class="text-right">Cant</th>
              <th class="text-right">Subtotal</th>
              <th class="text-right">Eliminar</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(dt, i) in purchase.details" :key="i">
              <td class="text-left">{{dt.item_id}}</td>
              <td class="text-right">Descripcion</td>
              <td class="text-right">Talla</td>
              <td class="text-right">Color</td>
              <td class="text-right">% Iva</td>
              <td class="text-right">P/Compra</td>
              <td class="text-right">P/Venta</td>
              <td class="text-right">% Desc</td>
              <td class="text-right">P/Final</td>
              <td class="text-right">P/Iva</td>
              <td class="text-right">Cant</td>
              <td class="text-right">Subtotal</td>
              <td class="text-right">Eliminar</td>
            </tr>
          </tbody>
        </q-markup-table>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-sm">
      <q-card-section>
        <q-item>
          <q-item-section>
            <q-input outlined dense label="Subtotal"></q-input>
          </q-item-section>
          <q-item-section>
            <q-input outlined dense label="Descuento"></q-input>
          </q-item-section>
          <q-item-section>
            <q-input outlined dense label="IVA"></q-input>
          </q-item-section>
          <q-item-section>
            <q-input outlined dense label="Total"></q-input>
          </q-item-section>
          <q-item-section>
            <q-input outlined dense label="Valor abono"></q-input>
          </q-item-section>
          <q-item-section>
            <q-btn outline icon="save" color="primary" dense>Guardar</q-btn>
          </q-item-section>
        </q-item>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import {useQuasar} from "quasar";
import {computed, onMounted, reactive, ref, watch} from 'vue'
import {storeToRefs} from "pinia";
import {usePurchaseStore} from "stores/purchase-store.js";
import {useProductStore} from "stores/product-store.js";
import {useRouter} from "vue-router";
import Swal from "sweetalert2";

const $router = useRouter();
const $q = useQuasar();
const purchaseStore = usePurchaseStore();
const {purchases,} = storeToRefs(purchaseStore);
const productStore = useProductStore()
const {products} = storeToRefs(productStore)

const refs = ref([]);
const setInputRef = (index) => (el) => {
  refs.value[index] = el; // Almacena la referencia en el arreglo
};
const focusNext = (index) => {
  setTimeout(() => {
    if (refs.value[index]) {
      refs.value[index].focus();
    } else if (index === refs.value.length) {
      // Vuelve al primer campo si se supera el índice
      refs.value[0].focus();
    }
  }, 100);
};

let filteredProducts = ref(products.value);
let purchase = ref({
  details: []
})

let detailInProcess = reactive({
  item_id: null,
  color_id: null,
  purchase_price: computed(() => products.value.find(i => i.id === detailInProcess.item_id)?.purchase_price),
  sale_price: computed(() => products.value.find(i => i.id === detailInProcess.item_id)?.sale_price),
  final_purchase_price: 0,
  discount_percent: 0,
  size_id: null,
  units: 0,
})

watch(detailInProcess, () => {
  if (detailInProcess.item_id){
    detailInProcess.final_purchase_price = detailInProcess.purchase_price - detailInProcess.purchase_price * detailInProcess.discount_percent/100
  }
})

const addDetail = () => {
  purchase.value.details.push({
    item_id: null,
    color_id: null,
    size_id: null,
    purchase_price: null,
    sale_price: null,
    percent_discount: 0,
    units: 1,
  });
}

const removeDetail = () => {
  purchase.value.details.pop();
}

const register = async () => {
  const response = await purchaseStore.postPurchase(token.value, purchase.value)
  if(await response){
    await $router.push({ path: '/purchases' })
  }
}

const productFilter  = (val, update) => {
  if (val === '') {
    update(() => {
      filteredProducts.value = products.value

      // here you have access to "ref" which
      // is the Vue reference of the QSelect
    })
    return
  }

  update(() => {
    const needle = val.toLowerCase()
    filteredProducts.value = products.value.filter(v => v.name.toLowerCase().indexOf(needle) > -1 || v.description.toLowerCase().indexOf(needle) > -1)
  })
}

const openSizesDialog = async () => {
  let message = null;
  if (!purchase.value.supplier_id){
    message = 'El campo proveedor es requerido';
  }
  if (message){
    await Swal.fire({
      title: 'Error de Validación!',
      text: message,
      icon: 'error',
      confirmButtonText: 'Aceptar'
    })
    await focusNext('purchase.entity_id')
  }
}
</script>
<style scoped>
.full-height {
  height: 100vh; /* Altura total de la pantalla */
  display: flex;
  flex-direction: column;
}

.flex-grow {
  flex-grow: 1; /* Expande el elemento para ocupar el espacio disponible */
  overflow-y: auto; /* Añade un scrollbar si el contenido es muy grande */
}

/*Input*/
:deep(.input .q-field__control),
:deep(.input .q-field__marginal),
:deep(.q-input .q-field__control),
:deep(.q-input .q-field__marginal) {
  height: 32px;
  font-size: 11px;
}

/*Select*/
:deep(.q-field--auto-height .q-field__control, ) {
  min-height: 32px !important;
  height: 32px !important;
  font-size: 12px;
}

:deep(.q-field__label) {
  font-size: 12px;
  margin-top: -2px;
}

:deep(.q-field--auto-height .q-field__native, .q-field--auto-height
.q-field__prefix,.q-field--auto-height .q-field__suffix ) {
  height: 32px !important;
  min-height: 32px !important;
}

:deep(.q-field__marginal) {
  height: 32px !important;
}

:deep(.q-field--standard .q-field__control:before) {
  border: 0px !important;
  outline: none !important;
  padding: 2px !important;
}

:deep(.q-field--auto-height .q-field__control-container) {
  padding-bottom: 2px !important;
  outline: none !important;
}
</style>

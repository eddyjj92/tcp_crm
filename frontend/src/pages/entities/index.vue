<template>
  <q-page class="q-pa-sm">
    <q-card>
      <q-table
        ref="tableRef"
        title="Entidades"
        :rows="entities"
        :hide-header="mode === 'grid'"
        :columns="columns"
        row-key="name"
        :grid="mode === 'grid'"
        :filter="filter"
        :pagination:sync="pagination"
        class="my-sticky-last-column-table"
      >
        <template v-slot:top-right="props">
          <q-btn v-if="user?.roles[0]?.permissions.find(p => p.name === 'crear tiendas')" @click="dialog = true;" icon="add" outline color="primary" label="Registrar" class="q-mr-xs"/>

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
              <q-checkbox :disable="entities.length === 0" class="float-left" v-if="col.label === 'ID'" v-model="checkAll"/>
              <span class="float-left text-bold" style="font-size: 11px" v-else>{{ col.label }}<br>
                <q-select v-if="col.name === 'entity_type'" outlined dense
                  :options="[{label: 'Todos', value: ''}, ...entity_types]"
                  v-model="filters.entity_type" style="min-width: 100px"
                  type="text" emit-value map-options
                >
                  <template v-slot:selected-item="{ opt }">
                    <span style="margin-bottom: 20px;">{{ opt.label }}</span>
                  </template>
                </q-select>
                <q-select v-if="col.name === 'identification_type'" outlined dense
                          :options="[{label: 'Todos', value: ''}, ...id_types]"
                          v-model="filters.identification_type" style="min-width: 100px"
                          type="text" emit-value map-options
                >
                  <template v-slot:selected-item="{ opt }">
                    <span style="margin-bottom: 20px;">{{ opt.label }}</span>
                  </template>
                </q-select>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.name === 'identification'" type="text" style="min-width: 60px"/>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.name === 'name'" type="text" style="min-width: 60px"/>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.name === 'lastnames'" type="text" style="min-width: 60px"/>
                <q-input debounce="300" outlined dense v-model="filters.name" v-if="col.name === 'trade_name'" type="text" style="min-width: 60px"/>
                <q-input v-if="col.name === 'birthdate'" outlined dense v-model="filters.birthdate" mask="##/##/####" debounce="300" style="min-width: 100px">
                  <template v-slot:append>
                    <q-icon size="xs" name="event" class="cursor-pointer q-mr-xs">
                      <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                        <q-date v-model="filters.birthdate" mask="DD/MM/YYYY">
                          <div class="row items-center justify-end">
                            <q-btn v-close-popup label="Close" color="primary" flat />
                          </div>
                        </q-date>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
                <q-input debounce="300" outlined dense v-model="filters.address" v-if="col.name === 'address'" type="text" style="min-width: 60px"/>
                <q-input debounce="300" outlined dense v-model="filters.phone" v-if="col.name === 'phone'" type="text" style="min-width: 60px"/>
                <q-input debounce="300" outlined dense v-model="filters.cellular" v-if="col.name === 'cellular'" type="text" style="min-width: 60px"/>
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
              <span v-if="col.field === 'hexadecimal_code'" class="float-left flex flex-center">
                {{ col.value }}
                <div style="display: inline-block; width: 50px; height: 20px; margin-left: 10px;border-radius: 5px; border: 1px solid" :style="`background-color: ${col.value}`"></div>
              </span>
              <span v-if="col.field !== 'id' && col.field !== 'hexadecimal_code'" class="float-left">{{ col.value }}</span>
            </q-td>
            <q-td auto-width align="center">
              <q-btn v-if="user?.roles[0]?.permissions.find(p => p.name === 'editar entidades')" @click="edit(props.cols[0].value)" class="q-ml-sm bg-yellow-6 text-white" style="width: 35px">
                <q-icon size="xs" name="edit" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Editar Entidad.
                </q-tooltip>
              </q-btn>
              <q-btn v-if="user?.roles[0]?.permissions.find(p => p.name === 'eliminar entidades')" @click="deleteDialogOpen(props.cols[0].value)" class="q-ml-sm bg-negative text-white" style="width: 35px">
                <q-icon size="xs" name="delete_forever" />
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">
                  Eliminar Entidad.
                </q-tooltip>
              </q-btn>
              <template v-if="!(user?.roles[0]?.permissions.find(p => p.name === 'editar colores' && user?.roles[0]?.permissions.find(p => p.name === 'eliminar colores')))">
                <span>-</span>
              </template>
            </q-td>
          </q-tr>
          <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
          </q-inner-loading>
        </template>
      </q-table>
    </q-card>
    <q-dialog v-model="dialog" persistent @before-hide="selected = []; entity = {social_networks: [], active: false}">
      <q-card style="width: 650px; max-width: 80vw;">
        <q-card-section>
          <div class="text-h6">
            <span v-if="selected.length === 0"> Registrar Entidad</span>
            <span v-else> Editar Entidad</span>
            <q-btn round flat dense icon="close" class="float-right" color="grey-8" v-close-popup></q-btn>
          </div>
        </q-card-section>
        <q-separator inset></q-separator>
        <q-card-section style="max-height: 65vh" class="scroll">
          <q-form class="q-gutter-md">
            <q-list>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Tipo de Entidad</q-item-label>
                  <q-select dense outlined emit-value map-options v-model="entity.entity_type" :options="entity_types" label="Tipo de Entidad"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Tipo de Identificación</q-item-label>
                  <q-select dense outlined emit-value map-options v-model="entity.identification_type" :options="id_types" label="Tipo de Identificación"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Identificación</q-item-label>
                  <q-input dense outlined v-model="entity.identification" label="Identificación"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Nombre</q-item-label>
                  <q-input dense outlined v-model="entity.name" label="Nombre"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Apellidos</q-item-label>
                  <q-input dense outlined v-model="entity.lastnames" label="Apellidos"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Nombre Comercial</q-item-label>
                  <q-input dense outlined v-model="entity.trade_name" label="Nombre Comercial"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Correo Electrónico</q-item-label>
                  <q-input dense outlined v-model="entity.email" label="Correo Electrónico"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Fecha de Nacimiento</q-item-label>
                  <q-input outlined dense v-model="entity.birthdate" mask="##/##/####" readonly>
                    <template v-slot:append>
                      <q-icon name="event" class="cursor-pointer">
                        <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                          <q-date v-model="entity.birthdate" mask="DD/MM/YYYY">
                            <div class="row items-center justify-end">
                              <q-btn v-close-popup label="Close" color="primary" flat />
                            </div>
                          </q-date>
                        </q-popup-proxy>
                      </q-icon>
                    </template>
                  </q-input>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Dirección</q-item-label>
                  <q-input dense outlined v-model="entity.address" label="Dirección"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Teléfono</q-item-label>
                  <q-input dense outlined v-model="entity.phone" label="Teléfono"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Celular</q-item-label>
                  <q-input dense outlined v-model="entity.cellular" label="Celular"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Fax</q-item-label>
                  <q-input dense outlined v-model="entity.fax" label="Fax"/>
                </q-item-section>
              </q-item>

              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Tipo de Persona</q-item-label>
                  <q-select
                    dense
                    outlined
                    :options="person_types"
                    v-model="entity.person_type_id"
                    label="Tipo de Persona"
                    option-label="name"
                    option-value="id"
                    map-options
                    emit-value
                  />
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Tipo de Impuesto</q-item-label>
                  <q-select
                    dense
                    outlined
                    :options="tax_levels"
                    v-model="entity.tax_level_id"
                    label="Tipo de Impuesto"
                    option-label="name"
                    option-value="id"
                    map-options
                    emit-value
                  />
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Tipo de Régimen</q-item-label>
                  <q-select
                    dense
                    outlined
                    :options="regime_types"
                    v-model="entity.regime_type_id"
                    label="Tipo de Red Social"
                    option-label="name"
                    option-value="id"
                    map-options
                    emit-value
                  />
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Capacidad de Credito</q-item-label>
                  <q-input dense outlined v-model="entity.credit_cap" label="Capacidad de Credito"/>
                </q-item-section>
                <q-item-section>
                  <q-item-label class="q-pb-xs">DV</q-item-label>
                  <q-input dense outlined v-model="entity.dv" label="DV"/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Redes Sociales</q-item-label>
                  <q-item v-for="(sn, i) in entity.social_networks" :key="i" class="q-px-none">
                    <q-item-section>
                      <q-select dense outlined emit-value map-options :options="sn_types" v-model="sn.social_network" label="Tipo de Red Social"/>
                    </q-item-section>
                    <q-item-section>
                      <q-input dense outlined v-model="sn.url" label="URL Perfil"/>
                    </q-item-section>
                  </q-item>
                  <span class="text-center text-bold text-h6 text-grey-8" v-if="entity.social_networks.length === 0">Sin redes sociales</span>
                  <q-item>
                    <q-item-section>
                      <q-btn @click="addSocialNetwork" color="positive" icon="add" >Agregar</q-btn>
                    </q-item-section>
                    <q-item-section v-if="entity.social_networks.length > 0">
                      <q-btn @click="removeSocialNetwork" color="negative" icon="remove" >Remover</q-btn>
                    </q-item-section>
                  </q-item>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-checkbox v-model="entity.active" label="Activo" />
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label class="q-pb-xs">Observaciones</q-item-label>
                  <q-input outlined type="textarea" v-model="entity.observations" label="Observaciones" />
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
import {useEntityStore} from "stores/entity-store";

const $q = useQuasar();
const entityStore = useEntityStore();
const {entities, entity_types, id_types, sn_types, person_types, regime_types, tax_levels} = storeToRefs(entityStore);
let loading = ref(false);

const columns = [
  { name: "id", align: "left", label: "ID",  field: "id", sortable: true },
  { name: "entity_type", align: "left", label: "Tipo",  field: "entity_type", sortable: true },
  { name: "identification_type", align: "left", label: "Tipo ID",  field: "identification_type", sortable: true },
  { name: "identification", align: "left", label: "ID No.",  field: "identification", sortable: true },
  { name: "name", align: "left", label: "Nombre",  field: "name", sortable: true },
  { name: "lastnames", align: "left", label: "Apellidos",  field: "lastnames", sortable: true },
  { name: "trade_name", align: "left", label: "Nombre Comercial",  field: "trade_name", sortable: true },
  { name: "birthdate", align: "left", label: "Fecha Nacimiento",  field: "birthdate", sortable: true },
  { name: "address", align: "left", label: "Dirección",  field: "address", sortable: true },
  { name: "phone", align: "left", label: "Teléfono",  field: "phone", sortable: true },
  { name: "cellular", align: "left", label: "Celular",  field: "cellular", sortable: true },
];

const tableRef = ref()
const selected = ref([])
let filter = ref("")
let filters = reactive({
  entity_type: "",
  identification_type: "",
  name: "",
  birthdate: null
});
let entity = ref({
  social_networks: [],
  active: false,
})
let dialog = ref(false);
let mode = ref("list")
let pagination = ref({
  rowsPerPage: 10
});

const checkAll = ref(false)
watch(checkAll, () => {
  if (checkAll.value && entities.value.length !== selected.value.length){
    entities.value.forEach(st => {
      selected.value.push(st.id)
    })
  }else if (!checkAll.value){
    selected.value = []
  }
})

watch(selected, () => {
  checkAll.value = entities.value.length === selected.value.length && entities.value.length > 0;
})

onMounted(async () => {
  await entityStore.getEntities(null, !entityStore.fetched)
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
      entities.value.map(row =>
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

  const status = exportFile("entidades.csv", content, "text/csv");

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
  const url = `/api/entities?name=${filters.name}`
  await entityStore.getEntities(url, false)
    .finally(() => loading.value = false);
})

const register = async () => {
  const registered = await entityStore.postEntity(token.value, entity.value)
  if (registered){
    dialog.value = false;
    loading.value = true;
    entity.value = {};
    await entityStore.getEntities(null, false)
      .finally(() => loading.value = false);
  }
}

const edit = (id) => {
  const edit = entities.value.find(u => u.id === id);
  entity.value.entity_type = edit.entity_type;
  entity.value.identification_type = edit.identification_type;
  entity.value.identification = edit.identification;
  entity.value.name = edit.name;
  entity.value.lastnames = edit.lastnames;
  entity.value.trade_name = edit.trade_name;
  entity.value.birthdate = edit.birthdate;
  entity.value.address = edit.address;
  entity.value.phone = edit.phone;
  entity.value.cellular = edit.cellular;
  entity.value.fax = edit.fax;
  entity.value.email = edit.email;
  entity.value.credit_cap = edit.credit_cap;
  entity.value.dv = edit.dv;
  entity.value.social_networks = edit.social_networks;
  entity.value.active = edit.active;
  entity.value.observations = edit.observations;
  entity.value.person_type_id = edit.person_type_id;
  entity.value.tax_level_id = edit.tax_level_id;
  entity.value.regime_type_id = edit.regime_type_id;

  selected.value = [id];
  dialog.value = true;
}

const update = async () => {
  const updated = await entityStore.putEntity(selected.value[0], entity.value)
  if (updated){
    dialog.value = false;
    loading.value = true;
    selected.value = [];
    await entityStore.getEntities(null, false)
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
  const deleted = await entityStore.deleteEntity(deleteDialog.value)
  if (deleted){
    deleteDialog.value = null;
    loading.value = true;
    await entityStore.getEntities(null, false)
      .finally(() => loading.value = false);
  }
}

const addSocialNetwork = () => {
  entity.value.social_networks.push({ social_network: "", url: "" });
}

const removeSocialNetwork = () => {
  entity.value.social_networks.pop()
}

</script>

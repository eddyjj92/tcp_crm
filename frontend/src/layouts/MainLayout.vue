<template>
  <q-layout view="hHh lpR lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />
        <q-btn
          v-if="canGoBackRef"
          flat
          dense
          round
          icon="arrow_back"
          aria-label="Menu"
          @click="router.back()"
        />
        <q-toolbar-title>
          TCP-CRM
        </q-toolbar-title>


      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item-label
          header
        >
          <span class="text-h6">MENÚ PRINCIPAL</span>
        </q-item-label>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Dashboard
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="settings" />
          </q-item-section>
          <q-item-section class="text-bold">
            Configuración
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/products">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="view_in_ar" />
          </q-item-section>
          <q-item-section class="text-bold">
            Productos
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/suppliers">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="diversity_3" />
          </q-item-section>
          <q-item-section class="text-bold">
            Proveedores
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/purchases">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="shopping_cart" />
          </q-item-section>
          <q-item-section class="text-bold">
            Compras
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="point_of_sale" />
          </q-item-section>
          <q-item-section class="text-bold">
            Ventas
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="inventory" />
          </q-item-section>
          <q-item-section class="text-bold">
            Inventario
          </q-item-section>
        </q-item>
        <q-separator inset />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer elevated class="bg-grey-8 text-white" style="height: 40px">
      <q-toolbar class="full-width q-pb-sm">
        <q-toolbar-title>
          <div class="text-center" style="font-size: 12px">Desarrollador: Lic.Eddy Javier Jorge Herrera</div>
          <div class="text-center" style="font-size: 12px">Correo: eddyjj92@gmail.com - Teléfono: +5356840677</div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script setup>
import {onUpdated, ref} from 'vue'
import {useRouter} from "vue-router";

defineOptions({
  name: 'MainLayout'
})

const leftDrawerOpen = ref(false)

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

const router = useRouter();
let canGoBackRef = ref(false)
const canGoBack = () => {
  const previousRoute = router.options.history.state.back;
  canGoBackRef.value = !!previousRoute;
};

onUpdated(() => {
  canGoBack()
})
</script>

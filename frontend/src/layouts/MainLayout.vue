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

        <q-toolbar-title>
          TCP-CRM
        </q-toolbar-title>

        <div>Quasar v{{ $q.version }}</div>
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
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Configuración
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/products">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Productos
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Proveedores
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Compras
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
          </q-item-section>
          <q-item-section class="text-bold">
            Ventas
          </q-item-section>
        </q-item>
        <q-separator inset />
        <q-item to="/">
          <q-item-section avatar class="flex flex-center q-pa-none">
            <q-icon class="text-center" name="house" />
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
import { ref } from 'vue'

defineOptions({
  name: 'MainLayout'
})

const leftDrawerOpen = ref(false)

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>

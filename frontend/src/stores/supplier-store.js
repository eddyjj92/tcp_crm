import { defineStore } from 'pinia';
import {ref} from "vue";
import {Loading} from "quasar";
import {Delete, List, Store, Update} from "app/wailsjs/go/controllers/SupplierController.js";
import {helpers} from "src/helpers/index.js";

const {message} = helpers();

export const useSupplierStore = defineStore('supplier', () => {
  let suppliers = ref([]);
  let fetched = ref(false);
  (function construct(){
    suppliers.value = JSON.parse(localStorage.getItem('suppliers') ?? '[]');
  })();

  const getSuppliers = async (filters = null, loading = true, hideLoading = true) => {
    loading ? await Loading.show() : null;
    return await List().then(response => {
      suppliers.value = response;
      localStorage.setItem('suppliers', JSON.stringify(suppliers.value))
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => hideLoading ? Loading.hide() : null)
  }

  const postSupplier = async (supplier) => {
    await Loading.show();
    return await Store(supplier).then(response => {
      suppliers.value = [...suppliers.value, response];
      localStorage.setItem('suppliers', JSON.stringify(suppliers.value))
      message(`Proveedor ${response.name} registrado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const putSupplier = async (id, product) => {
    await Loading.show();
    return await Update(id, product).then(response => {
      suppliers.value = [...suppliers.value.filter(el => el.id !== id), response];
      localStorage.setItem('suppliers', JSON.stringify(suppliers.value))
      message(`Proveedor ${response.name} actualizado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const deleteSupplier = async (id) => {
    await Loading.show();
    return await Delete(id).then(response => {
      suppliers.value = suppliers.value.filter(el => el.id !== id);
      localStorage.setItem('suppliers', JSON.stringify(suppliers.value))
      message(`Proveedor ${response.name} eliminado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }


  return {
    suppliers,
    fetched,
    getSuppliers,
    postSupplier,
    putSupplier,
    deleteSupplier,
  }
});

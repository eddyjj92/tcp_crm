import { defineStore } from 'pinia';
import {ref} from "vue";
import {Loading} from "quasar";
import {Delete, List, Store, Update} from "app/wailsjs/go/controllers/ProductController.js";
import {helpers} from "src/helpers/index.js";

const {message} = helpers();

export const useProductStore = defineStore('product', () => {
  let products = ref([]);
  let fetched = ref(false);
  (function construct(){
    products.value = JSON.parse(localStorage.getItem('products') ?? '[]');
  })();

  const getProducts = async (filters = null, loading = true, hideLoading = true) => {
    loading ? await Loading.show() : null;
    return await List().then(response => {
      products.value = response;
      localStorage.setItem('products', JSON.stringify(products.value))
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => hideLoading ? Loading.hide() : null)
  }

  const postProduct = async (product) => {
    await Loading.show();
    return await Store(product).then(response => {
      products.value = [...products.value, response];
      localStorage.setItem('products', JSON.stringify(products.value))
      message(`Producto ${response.name} registrado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const putProduct = async (id, product) => {
    await Loading.show();
    return await Update(id, product).then(response => {
      products.value = [...products.value.filter(el => el.id !== id), response];
      localStorage.setItem('products', JSON.stringify(products.value))
      message(`Producto ${response.name} actualizado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const deleteProduct = async (id) => {
    await Loading.show();
    return await Delete(id).then(response => {
      products.value = products.value.filter(el => el.id !== id);
      localStorage.setItem('products', JSON.stringify(products.value))
      message(`Producto ${response.name} eliminado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }


  return {
    products,
    fetched,
    getProducts,
    postProduct,
    putProduct,
    deleteProduct,
  }
});

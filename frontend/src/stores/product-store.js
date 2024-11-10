import { defineStore } from 'pinia';
import {ref} from "vue";
import {api} from "boot/axios";
import {Loading, Notify} from "quasar";
import {List, Store} from "app/wailsjs/go/controllers/ProductController.js";

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
      Notify.create({
        message: error,
        type: "negative"
      });
      return false
    }).finally(() => hideLoading ? Loading.hide() : null)
  }

  const postProduct = async (product) => {
    await Loading.show();
    return await Store(product).then(response => {
      products.value = [...products.value, response];
      localStorage.setItem('products', JSON.stringify(products.value))
      return true;
    }).catch(error => {
      const close = Notify.create({
        message: error,
        type: "negative",
        position: "top-right",
        html: true,
        badgeClass: "custom-notify",
        timeout: 5000, // Tiempo en milisegundos, 5000ms = 5 segundos
        actions: [
          {
            icon: "close", // Icono de cerrar
            color: "white", // Color del icono
            handler: () => {
              close()
            }
          }
        ],
        progress: true, // Barra de progreso visible
      });
      return false
    }).finally(() => Loading.hide())
  }

  const putProduct = async (id, product) => {

  }

  const deleteProduct = async (id) => {

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

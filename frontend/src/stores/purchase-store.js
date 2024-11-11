import { defineStore } from 'pinia';
import {ref} from "vue";
import {Loading} from "quasar";
import {Delete, List, Store, Update} from "app/wailsjs/go/controllers/PurchaseController.js";
import {helpers} from "src/helpers/index.js";

const {message} = helpers();

export const useProductStore = defineStore('product', () => {
  let purchases = ref([]);
  let fetched = ref(false);
  (function construct(){
    purchases.value = JSON.parse(localStorage.getItem('purchases') ?? '[]');
  })();

  const getPurchases = async (filters = null, loading = true, hideLoading = true) => {
    loading ? await Loading.show() : null;
    return await List().then(response => {
      purchases.value = response;
      localStorage.setItem('purchases', JSON.stringify(purchases.value))
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => hideLoading ? Loading.hide() : null)
  }

  const postPurchase = async (purchase) => {
    await Loading.show();
    return await Store(purchase).then(response => {
      purchases.value = [...purchases.value, response];
      localStorage.setItem('purchases', JSON.stringify(purchases.value))
      message(`Compra ${response.name} registrada con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const putPurchase = async (id, purchase) => {
    await Loading.show();
    return await Update(id, purchase).then(response => {
      purchases.value = [...purchases.value.filter(el => el.id !== id), response];
      localStorage.setItem('products', JSON.stringify(products.value))
      message(`Compra ${response.name} actualizado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  const deletePurchase = async (id) => {
    await Loading.show();
    return await Delete(id).then(response => {
      purchases.value = purchases.value.filter(el => el.id !== id);
      localStorage.setItem('purchases', JSON.stringify(purchases.value))
      message(`Compra ${response.name} eliminado con éxito`, "positive")
      return true;
    }).catch(error => {
      message(error, "negative")
      return false
    }).finally(() => Loading.hide())
  }

  return {
    purchases,
    fetched,
    getPurchases,
    postPurchase,
    putPurchase,
    deletePurchase,
  }
});

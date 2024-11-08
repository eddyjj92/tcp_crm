import { defineStore } from 'pinia';
import {ref} from "vue";
import {api} from "boot/axios";
import {Loading, Notify} from "quasar";
import {useIvaStore} from "stores/iva-store.js";
import {useStoreStore} from "stores/store-store.js";

const ivaStore = useIvaStore();
const storeStore = useStoreStore();

export const usePurchaseStore = defineStore('purchase', () => {
  let purchases = ref([]);
  let suppliers = ref([]);
  let payment_shapes = ref([]);
  let purchase_states = ref([]);
  let fetched = ref(false);
  (function construct(){
    purchases.value = JSON.parse(localStorage.getItem('purchases') ?? '[]');
    suppliers.value = JSON.parse(localStorage.getItem('suppliers') ?? '[]');
    payment_shapes.value = JSON.parse(localStorage.getItem('payment_shapes') ?? '[]');
    purchase_states.value = JSON.parse(localStorage.getItem('purchase_states') ?? '[]');
  })();

  const getPurchases = async (token, filters = null, loading = true, hideLoading = true) => {
    loading ? await Loading.show() : null;
    return await api.get(`${filters ?? '/api/purchases'}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        fetched.value = true;
        purchases.value = res.data.purchases;
        suppliers.value = res.data.suppliers;
        payment_shapes .value = res.data.payment_shapes ;
        purchase_states.value = res.data.purchase_states;
        ivaStore.ivas = res.data.ivas;
        storeStore.stores = res.data.stores;
        localStorage.setItem("purchases", JSON.stringify(purchases.value))
        localStorage.setItem("suppliers", JSON.stringify(suppliers.value))
        localStorage.setItem("payment_shapes", JSON.stringify(payment_shapes.value))
        localStorage.setItem("purchase_states", JSON.stringify(purchase_states.value))
        localStorage.setItem("ivas", JSON.stringify(ivaStore.ivas))
        localStorage.setItem("stores", JSON.stringify(storeStore.stores))
        return true
      })
      .catch(error => {
        let message;
        if (error.response?.data.message){
          message = error.response?.data.message
        }else {
          message = error.message
        }
        Notify.create({
          type: 'negative',
          message: message,
          position: 'top',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return false
      })
      .finally(() => hideLoading ? Loading.hide() : null)
  }

  const postPurchase = async (token, purchase) => {
    await Loading.show();
    return await api.post(`/api/purchases`, purchase,{
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        Notify.create({
          type: 'positive',
          message: res.data.message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return true
      })
      .catch(error => {
        let message;
        if (error.response?.data.validations){
          for(let msg of Object.values(error.response.data.validations)){
            Notify.create({
              type: 'negative',
              message: msg,
              position: 'top-right',
              progress: true,
              actions: [{ icon: 'close', color: 'white' }]
            })
          }
          return false;
        }else if (error.response?.data.message){
          message = error.response?.data.message
        }else {
          message = error.message
        }
        Notify.create({
          type: 'negative',
          message: message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return false;
      })
      .finally(() => Loading.hide())
  }

  const putPurchase = async (token, id, purchase) => {
    await Loading.show();
    return await api.put(`/api/purchases/${id}`, purchase,{
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        Notify.create({
          type: 'positive',
          message: res.data.message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return true
      })
      .catch(error => {
        let message;
        if (error.response?.data.validations){
          for(let msg of Object.values(error.response.data.validations)){
            Notify.create({
              type: 'negative',
              message: msg,
              position: 'top-right',
              progress: true,
              actions: [{ icon: 'close', color: 'white' }]
            })
          }
          return false;
        }else if (error.response?.data.message){
          message = error.response?.data.message
        }else {
          message = error.message
        }
        Notify.create({
          type: 'negative',
          message: message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return false;
      })
      .finally(() => Loading.hide())
  }

  const deletePurchase = async (token, id) => {
    await Loading.show();
    return await api.delete(`/api/purchases/${id}`,{
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        Notify.create({
          type: 'positive',
          message: res.data.message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return true
      })
      .catch(error => {
        let message;
        if (error.response?.data.message){
          message = error.response?.data.message
        }else {
          message = error.message
        }
        Notify.create({
          type: 'negative',
          message: message,
          position: 'top-right',
          progress: true,
          actions: [{ icon: 'close', color: 'white' }]
        })
        return false;
      })
      .finally(() => Loading.hide())
  }

  return {
    purchases,
    suppliers,
    payment_shapes,
    purchase_states,
    fetched,
    getPurchases,
    postPurchase,
    putPurchase,
    deletePurchase
  }
});

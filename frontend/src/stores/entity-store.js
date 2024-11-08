import { defineStore } from 'pinia';
import {ref} from "vue";
import {api} from "boot/axios";
import {Loading, Notify} from "quasar";

export const useEntityStore = defineStore('entity', () => {
  let entities = ref([]);
  let entity_types = ref([]);
  let id_types = ref([]);
  let sn_types = ref([]);
  let person_types = ref([]);
  let regime_types = ref([]);
  let tax_levels = ref([]);
  let fetched = ref(false);

  (function construct(){
    entities.value = JSON.parse(localStorage.getItem('entities') ?? '[]');
    entity_types.value = JSON.parse(localStorage.getItem('entity_types') ?? '[]');
    id_types.value = JSON.parse(localStorage.getItem('id_types') ?? '[]');
    sn_types.value = JSON.parse(localStorage.getItem('sn_types') ?? '[]');
    person_types.value = JSON.parse(localStorage.getItem('person_types') ?? '[]');
    regime_types.value = JSON.parse(localStorage.getItem('regime_types') ?? '[]');
    tax_levels.value = JSON.parse(localStorage.getItem('tax_levels') ?? '[]');
  })();

  const getEntities = async (token, filters = null, loading = true, hideLoading = true) => {
    loading ? await Loading.show() : null;
    return await api.get(`${filters ?? '/api/entities'}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        fetched.value = true;
        entities.value = res.data.entities;
        entity_types.value = res.data.entity_types;
        id_types.value = res.data.id_types;
        sn_types.value = res.data.sn_types;
        person_types.value = res.data.person_types;
        regime_types.value = res.data.regime_types;
        tax_levels.value = res.data.tax_levels;
        localStorage.setItem("entities", JSON.stringify(entities.value));
        localStorage.setItem("entity_types", JSON.stringify(entity_types.value));
        localStorage.setItem("id_types", JSON.stringify(id_types.value));
        localStorage.setItem("sn_types", JSON.stringify(sn_types.value));
        localStorage.setItem("person_types", JSON.stringify(person_types.value));
        localStorage.setItem("regime_types", JSON.stringify(regime_types.value));
        localStorage.setItem("tax_levels", JSON.stringify(tax_levels.value));
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

  const postEntity = async (token, entity) => {
    await Loading.show();
    return await api.post(`/api/entities`, entity,{
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

  const putEntity = async (token, id, entity) => {
    await Loading.show();
    return await api.put(`/api/entities/${id}`, entity,{
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

  const deleteEntity = async (token, id) => {
    await Loading.show();
    return await api.delete(`/api/entities/${id}`,{
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
    entities,
    entity_types,
    id_types,
    sn_types,
    person_types,
    regime_types,
    tax_levels,
    fetched,
    getEntities,
    postEntity,
    putEntity,
    deleteEntity
  }
});

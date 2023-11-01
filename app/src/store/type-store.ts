import { defineStore } from "pinia";
import { ref } from "vue";
import lodash from "lodash";

export const useTypeStore = defineStore("type.store", () => {
  const rows = ref<DocumentType[]>([]);

  function create(payload: DocumentType) {
    rows.value.unshift(payload);
  }

  function edit(payload: DocumentType) {
    rows.value = rows.value.map((r) => (r.id === payload.id ? payload : r));
  }

  function destroy(payload: DocumentType) {
    const index = lodash.findIndex(rows.value, { id: payload.id });
    if (index >= 0) rows.value.splice(index, 1);
  }

  function setRows(payload: DocumentType[]) {
    rows.value = payload;
  }

  return { rows, create, edit, destroy, setRows };
});

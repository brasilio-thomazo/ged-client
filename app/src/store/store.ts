import axios from "axios";
import { defineStore } from "pinia";
import { inject, ref } from "vue";

export const useStore = defineStore("main.store", () => {
  const user = ref<User>();

  function setUser(payload: User) {
    user.value = payload;
  }

  function setToken(payload: string) {
    localStorage.setItem("token", payload);
  }

  function http() {
    const http = inject("http", axios);
    const token = localStorage.getItem("token");
    http.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    return http;
  }

  return { setUser, setToken, user, http };
});

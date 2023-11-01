import { createApp } from "vue";
import { createPinia } from "pinia";
import main from "./app.vue";
import router from "./router";
import axios from "axios";
import "./scss/main.scss";

const token = localStorage.getItem("token");
axios.defaults.baseURL = import.meta.env.VITE_API_URL;
console.log(import.meta.env.VITE_API_URL);
axios.defaults.headers.common["X-Requested-With"] = "XMLHttpRequest";
axios.defaults.headers.common["Content-Type"] = "application/json";
axios.defaults.headers.common.Accept = "application/json";
axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;

const app = createApp(main);
app.provide("http", axios);
app.use(createPinia());
app.use(router);
app.mount("#app");

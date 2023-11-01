<template>
  <header v-if="store.user">
    <nav class="navbar">
      <div class="brand">
        <a href="/">
          <span class="material-icons">picture_as_pdf</span>
          <span class="title">EDContent</span>
        </a>
      </div>
      <ul class="menu">
        <li><RouterLink to="/group">Grupos</RouterLink></li>
        <li><RouterLink to="/user">Usuários</RouterLink></li>
        <li><RouterLink to="/search">Pesquisa</RouterLink></li>
      </ul>
      <div class="buttons">
        <button
          @click="router.push({ name: 'profile' })"
          type="button"
          class="material-symbols-outlined"
          title="Perfil"
        >
          account_circle
        </button>
        <button
          @click="logout"
          type="button"
          class="material-symbols-outlined"
          title="Sair"
        >
          logout
        </button>
      </div>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { useStore } from "@/store/store";
import { useRouter } from "vue-router";
const store = useStore();
const router = useRouter();
function logout() {
  localStorage.removeItem("token");
  store.setToken("");
  router.push({ name: "login" });
}
</script>

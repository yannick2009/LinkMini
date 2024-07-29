import { createWebHistory, createRouter } from "vue-router";
import Home from "./pages/Home.vue";
import Expired from "./pages/Expired.vue";
import Stats from "./pages/Stats.vue";

const routes = [
  { name: "home", path: "/", component: Home },
  { name: "expired", path: "/expired", component: Expired },
  { name: "stats", path: "/stats", component: Stats },
  {name: "not-found", path: "/:pathMatch(.*)*", redirect: { name: "home" }}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

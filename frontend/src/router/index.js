import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../views/Home.vue";
import AllServices from "../views/AllServices.vue";
import AllLogs from "../views/AllLogs.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/service/all/",
    name: "Services",
    component: AllServices,
  },
  {
    path: "/service/new/",
    name: "NewService",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/NewService.vue"),
  },
  {
    path: "/logs/all/",
    name: "Logs",
    component: AllLogs,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;

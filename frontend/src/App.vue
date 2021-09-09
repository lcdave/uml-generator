<template>
  <side-navigation />
  <router-view />
</template>

<script>
import SideNavigation from "@/components/SideNavigation.vue";

export default {
  name: "Home",
  components: {
    SideNavigation,
  },
};
</script>

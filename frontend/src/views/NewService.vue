<template>
  <main class="main">
    <top-navigation />
    <div class="view-title">
      <h2>Add new service</h2>
    </div>
    <div class="widgets">
      <widget title="Please fill the form" size="10">
        <template #content>
          <service-form />
          <router-link to="/"> Back to dashboard </router-link>
        </template>
      </widget>
    </div>
  </main>
</template>

<script>
import TopNavigation from "@/components/TopNavigation.vue";
import Widget from "@/components/Widget.vue";
import ServiceForm from "@/components/ServiceForm.vue";

export default {
  name: "New Service",
  components: {
    TopNavigation,
    Widget,
    ServiceForm,
  },
};
</script>

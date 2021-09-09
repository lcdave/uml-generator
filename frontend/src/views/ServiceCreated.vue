<template>
  <main class="main">
    <top-navigation />
    <div class="view-title">
      <h2>Add new service</h2>
    </div>
    <div class="widgets">
      <widget title="Thank you!" size="10">
        <template #content>
          <p>Your service has been created.</p>
          <div class="cta">
            <router-link to="/"> Back to dashboard </router-link>
          </div>
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

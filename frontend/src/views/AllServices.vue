<template>
  <main class="main">
    <top-navigation />
    <div class="view-title">
      <h2>Services</h2>
    </div>
    <div class="widgets">
      <widget title="Services overview" size="10">
        <template #content>
          <services showAll="true" />
        </template>
      </widget>
    </div>
  </main>
</template>

<script>
import TopNavigation from "@/components/TopNavigation.vue";
import Widget from "@/components/Widget.vue";
import Services from "@/components/Services.vue";

export default {
  name: "AllServices",
  components: {
    TopNavigation,
    Widget,
    Services,
  },
};
</script>

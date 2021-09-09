<template>
  <main class="main">
    <top-navigation />
    <div class="view-title">
      <h2>Logs</h2>
    </div>
    <div class="widgets">
      <widget title="Logs overview" size="10">
        <template #content>
          <logs showAll="true" />
        </template>
      </widget>
    </div>
  </main>
</template>

<script>
import TopNavigation from "@/components/TopNavigation.vue";
import Widget from "@/components/Widget.vue";
import Logs from "@/components/Logs.vue";

export default {
  name: "AllServices",
  components: {
    TopNavigation,
    Widget,
    Logs,
  },
};
</script>

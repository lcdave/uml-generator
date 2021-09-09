<template>
  <main class="main">
    <top-navigation />
    <div class="view-title">
      <h2>Dashboard</h2>
      <p>Monitoring overview with stats, logs, and summaries</p>
    </div>
    <div class="widgets">
      <widget title="Services overview" size="7" hasIcon="true">
        <template #content>
          <services />
        </template>
      </widget>
      <widget title="Get Premium" size="3">
        <template #content>
          Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam
          nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
        </template>
      </widget>
    </div>
    <div class="widgets">
      <widget title="Logs" size="6">
        <template #content>
          <logs />
        </template>
      </widget>
      <widget title="Notifications" size="2">
        <template #content>
          <donut v-if="this.countEmail" :countEmail="this.countEmail" />
        </template>
      </widget>
      <div class="widget stacked var_2">
        <widget title="" size="10" quickStat="1">
          <template #content>
            <quick-stat>
              <template #icon>
                <span class="icon-clock var_green"></span>
              </template>
              <template #value> 43d 19h </template>
              <template #description> global uptime </template>
            </quick-stat>
          </template>
        </widget>
        <widget title="" size="10" quickStat="1">
          <template #content>
            <quick-stat>
              <template #icon>
                <span class="icon-alert var_purple"></span>
              </template>
              <template #value>
                {{ this.countLogs }}
              </template>
              <template #description> upsees detected </template>
            </quick-stat>
          </template>
        </widget>
      </div>
    </div>
  </main>
</template>

<script>
import TopNavigation from "@/components/TopNavigation.vue";
import Widget from "@/components/Widget.vue";
import Logs from "@/components/Logs.vue";
import Services from "@/components/Services.vue";
import Donut from "@/components/Donut.vue";
import QuickStat from "@/components/QuickStat.vue";

import axios from "axios";

export default {
  name: "Home",
  components: {
    TopNavigation,
    Widget,
    Logs,
    Services,
    Donut,
    QuickStat,
  },
  data() {
    return {
      countEmailAPIUrl: "http://localhost:8085/emails/count/",
      countEmail: 0,
      countLogsAPIUrl: "http://localhost:8085/logs/count/",
      countLogs: 0,
    };
  },
  mounted() {
    this.setEmailCounter();
    this.setLogsCounter();
  },
  methods: {
    async setEmailCounter() {
      let response = "";
      response = await axios.get(this.countEmailAPIUrl);
      this.countEmail = response.data;
    },
    async setLogsCounter() {
      let response = "";
      response = await axios.get(this.countLogsAPIUrl);
      console.log(response);
      this.countLogs = response.data;
    },
  },
};
</script>

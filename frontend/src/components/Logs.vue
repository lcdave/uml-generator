<template>
  <div class="logs">
    <div class="logs__header">
      <div class="logs__title">Service</div>
      <div class="logs__title--large">Message</div>
      <div class="logs__title--short">Time</div>
    </div>
    <div class="logs__list">
      <div class="logs__item" v-for="logEntry in logs" :key="logEntry">
        <span class="logs__service">{{ logEntry.url }}</span>
        <span class="logs__message">{{ logEntry.description }}</span>
        <span class="logs__time">{{ logEntry.time }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Logs",
  data() {
    return {
      logs: [],
      limit: 3,
      apiUrl: "http://localhost:8085/logs/limited/",
    };
  },
  props: {
    showAll: Boolean,
  },
  created() {
    this.getList();
  },
  methods: {
    async getList() {
      let response = "";

      if (this.showAll) {
        response = await axios.get(`http://localhost:8085/logs/`);
        this.logs = response.data.logsOutput;
      } else {
        response = await axios.get(`${this.apiUrl}${this.limit}`);
        this.logs = response.data.logsOutput;
      }
    },
  },
};
</script>

<style lang="scss" scoped src="../scss/components/_logs.scss"></style>

<template>
  <div class="services">
    <div class="services__header">
      <div class="services__title">Service</div>
      <div class="services__title--large">Graph</div>
      <div class="services__title--short">Action</div>
    </div>
    <div class="services__list">
      <div class="services__item" v-for="service in services" :key="service">
        <div class="services__url">{{ service.url }}</div>
        <div class="services__graph">
          <graph :serviceID="service.id" />
        </div>
        <div class="services__actions">
          <span class="icon-analytics var_grey"></span>
          <span class="icon-pen var_grey"></span>
          <span class="icon-delete var_red"></span>
        </div>
      </div>
    </div>
    <div class="services__link">
      <router-link to="/service/all" v-if="!showAll">
        <button class="button">Show all</button>
      </router-link>
      <router-link to="/" v-if="showAll"> Back to dashboard </router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Graph from "./Graph.vue";

export default {
  components: { Graph },
  name: "Services",
  data() {
    return {
      services: [],
      limit: 3,
      apiUrl: "http://localhost:8085/services/limited/",
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
        response = await axios.get(`http://localhost:8085/services/`);
      } else {
        response = await axios.get(`${this.apiUrl}${this.limit}`);
      }

      this.services = response.data.services;
    },
    async getServiceResponseTimes(service_id) {
      const response = await axios.get(
        `http://localhost:8085/requests/${service_id}`
      );
      response.data.requests.forEach((request) => {
        let responseTimes = [];

        responseTimes.push(request.response_time);
        return this.responseTimes;
      });
    },
  },
};
</script>
<style lang="scss" src="../scss/components/_services.scss"></style>

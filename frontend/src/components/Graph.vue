<template>
  <sparkline v-if="this.responseTimes" v-bind:data="this.responseTimes" />
</template>

<script>
import axios from "axios";
import Sparkline from "./Sparkline.vue";

export default {
  components: { Sparkline },
  name: "Graph",
  data() {
    return {
      responseTimes: [],
    };
  },
  props: {
    serviceID: Number,
  },
  created() {
    this.getServiceResponseTimes(this.serviceID);
  },
  methods: {
    async getServiceResponseTimes(serviceID) {
      const response = await axios.get(
        `http://localhost:8085/requests/${serviceID}`
      );
      response.data.requests.forEach((request) => {
        this.responseTimes.push(request.response_time);
      });
    },
  },
};
</script>

<style lang="scss" scoped></style>

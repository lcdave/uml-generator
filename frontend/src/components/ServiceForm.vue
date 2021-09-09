<template>
  <div class="serviceform">
    <notification
      :variant="notification.variant"
      :message="notification.message"
      ref="notification"
    />
    <form class="form" v-on:submit.prevent="submitForm">
      <input
        type="text"
        id="name"
        name="name"
        placeholder="Service name"
        v-model="form.name"
      />
      <input
        type="text"
        id="url"
        name="url"
        placeholder="https://example.com"
        v-model="form.url"
      />
      <button type="submit">add service</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import Notification from "@/components/Notification.vue";

export default {
  data() {
    return {
      form: {
        name: "",
        url: "",
      },
      created: false,
      notification: {
        variant: "--neutral",
        message: "",
      },
    };
  },
  methods: {
    submitForm() {
      axios
        .post("http://localhost:8085/services/", this.form)
        .then((res) => {
          console.log(res);
          this.notification.variant = "--success";
          this.notification.message =
            "Success! Service has been added to the dashboard.";
          this.$refs.notification.fadeMe();
        })
        .catch((error) => {
          // error.response.status Check status code
          console.log(error);
          this.notification.variant = "--error";
          this.notification.message =
            "Failure! Service has not been created. Please try again or contact support@upsee.io";
          this.$refs.notification.fadeMe();
        });
    },
  },
  components: {
    Notification,
  },
};
</script>

<style lang="scss" src="../scss/components/_form.scss" scoped></style>
